package main

import (
    "auth-le-back/config"
    "auth-le-back/internal/app"
    "auth-le-back/pkg/mrapp"
    "auth-le-back/pkg/mrlib"
    "flag"
)

// go env

// Go MODULS https://www.youtube.com/watch?v=RRy286VuOkE
// go mod init auth-le-back
// go mod tidy
// go clean -i github.com/jackc/pgx/v5
// go mod vendor // перенос зависимостей в проект

// https://pkg.go.dev/testing
// https://www.gorillatoolkit.org
// https://gokit.io/
// https://echo.labstack.com/
// Также в Go есть некоторые встроенные возможности по наблюдению с помощью пакета expvar, позволяющего публиковать внутренние статусы и метрики, и облегчающего их добавление.

var configPath string

func init() {
   flag.StringVar(&configPath,"config-path", "./config/config.yaml", "Path to application config file")
}

func main() {
    flag.Parse()

    cfg := config.New(configPath)
    logger := mrlib.NewLogger(cfg.Log.Level, !cfg.Log.NoColor)
    translator := mrlib.NewTranslator(
        logger,
        mrlib.TranslatorOptions{
            DirPath: cfg.Translation.DirPath,
            FileType: cfg.Translation.FileType,
            LangCodes: mrapp.CastToLangCodes(cfg.Translation.LangCodes...),
        },
    )

    if cfg.Debug {
      logger.Info("DEBUG MODE: ON")
    }

    logger.Info("LOG LEVEL: %s", cfg.Log.Level)
    logger.Info("APP PATH: %s", cfg.AppPath)
    logger.Info("CONFIG PATH: %s", configPath)

    app.Run(cfg, logger, translator)
}
