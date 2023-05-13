// Package app configures and runs application
package app

import (
    "auth-le-back/config"
    "auth-le-back/internal/controller/dto"
    "auth-le-back/internal/controller/http_v1"
    "auth-le-back/internal/infrastructure/repository"
    "auth-le-back/internal/usecase"
    "auth-le-back/pkg/client/postgresql"
    "auth-le-back/pkg/mrapp"
    "auth-le-back/pkg/mrhttp"
    "auth-le-back/pkg/mrlib"
    "context"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// Run creates objects via constructors.
func Run(cfg *config.Config, logger mrapp.Logger, translator mrapp.Translator) {
    logger.Info("Create postgre connection")

    postgreOptions := postgresql.Options{
        Host: cfg.Storage.Host,
        Port: cfg.Storage.Port,
        Username: cfg.Storage.Username,
        Password: cfg.Storage.Password,
        Database: cfg.Storage.Database,
        MaxPoolSize: 1,
        ConnAttempts: 1,
        ConnTimeout: time.Duration(cfg.Storage.Timeout),
    }

    postgreSQLClient := postgresql.New(logger)

    if err := postgreSQLClient.Connect(context.TODO(), postgreOptions); err != nil {
        logger.Fatal("%s", err.Error())
    }

    accountStorage := repository.NewAccount(postgreSQLClient)
    userStorage := repository.NewUser(postgreSQLClient)

    authService := usecase.NewAuth(logger, accountStorage, userStorage)

    requestValidator := mrlib.NewValidator(logger)
    requestValidator.Register("login", dto.ValidateLogin)

    authHttp := http_v1.NewAuth(logger, requestValidator, authService)
    authCheckHttp := http_v1.NewAuthCheck(logger, requestValidator, authService)

    logger.Info("Create router")

    corsOptions := mrhttp.CorsOptions{
        AllowedOrigins: cfg.Cors.AllowedOrigins,
        AllowedMethods: cfg.Cors.AllowedMethods,
        AllowedHeaders: cfg.Cors.AllowedHeaders,
        ExposedHeaders: cfg.Cors.ExposedHeaders,
        AllowCredentials: cfg.Cors.AllowCredentials,
        Debug: cfg.Debug,
    }

    router := mrhttp.NewRouter(logger)
    router.RegisterMiddleware(
        mrhttp.NewCors(corsOptions),
        mrhttp.MiddlewareHeaders(logger, translator),
        mrhttp.MiddlewareAuthenticateUser(logger),
    )

    router.Register(authHttp, authCheckHttp)
    router.HandlerFunc(http.MethodGet, "/", MainPage)

    appStart(cfg, logger, router)
}

func appStart(cfg *config.Config, logger mrapp.Logger, router mrapp.Router) {
    logger.Info("Initialize application")

    server := mrhttp.NewServer(logger, mrhttp.ServerOptions{
        Handler: router,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 5 * time.Second,
        ShutdownTimeout: 3 * time.Second,
    })

    logger.Info("Start application")

    server.Start(mrhttp.ListenOptions{
        AppPath: cfg.AppPath,
        Type: cfg.Listen.Type,
        SockName: cfg.Listen.SockName,
        BindIP: cfg.Listen.BindIP,
        Port: cfg.Listen.Port,
    })

    signalAppChan := make(chan os.Signal, 1)
    signal.Notify(
        signalAppChan,
        syscall.SIGABRT,
        syscall.SIGQUIT,
        syscall.SIGHUP,
        os.Interrupt,
        syscall.SIGTERM,
    )

    select {
        case signalApp := <-signalAppChan:
            logger.Info("Application shutdown, signal: " + signalApp.String())

        case err := <-server.Notify():
            logger.Error(fmt.Errorf("http server shutdown: %w", err))
    }

    closeItems := []io.Closer{server}

    for _, closer := range closeItems {
        if err := closer.Close(); err != nil {
            logger.Error(fmt.Errorf("failed to close %v: %w", closer, err))
        }
    }
}

func MainPage(w http.ResponseWriter, r *http.Request) error {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))

    return nil
}
