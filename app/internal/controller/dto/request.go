package dto

//type CreateAccount struct {
//    UserEmail string  `json:"userEmail" validate:"required"`
//}

type AuthAccount struct {
    UserLogin string  `json:"userLogin" validate:"required"`
}

type ChangeLogin struct {
    NewLogin string  `json:"newLogin" validate:"required"`
}

type ChangeEmail struct {
    NewEmail string  `json:"newEmail" validate:"required"`
}

type ConfirmEmail struct {
    Code string  `json:"code" validate:"required"`
}


