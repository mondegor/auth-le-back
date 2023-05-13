package dto

import "auth-le-back/internal/entity"

type WaitingAccessAccountResponse struct {
    OperationToken entity.OperationToken  `json:"operationToken"` // required
    AccountId entity.AccountPrimaryKey `json:"accountId"`
    Message string `json:"message"` // required
}

type AccountResponse struct {
    Id        entity.AccountPrimaryKey `json:"accountId"`
    Login     string `json:"userLogin"`
    Email     string `json:"userEmail"`
    LoggedAt  string `json:"loggedAt"`
    LoggedIP  string `json:"lastIP"`
    Status    entity.AccountStatus `json:"status"`
}

func (ar *AccountResponse) Init(item *entity.AccountUser) {
    ar.Id = item.Account.Id
    ar.Login = item.User.Login
    ar.Email = item.User.Email
    ar.LoggedAt = item.User.LoggedAt
    ar.LoggedIP = item.User.LoggedIP.String()
    ar.Status = item.Account.Status
}
