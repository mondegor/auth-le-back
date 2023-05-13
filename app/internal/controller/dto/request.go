package dto

type CreateAccount struct {
   UserEmail string  `json:"userEmail" validate:"required,min=7,max=128,email"`
}

type AuthAccount struct {
    UserLogin string  `json:"userLogin" validate:"required,max=128"`
}

type ChangeLogin struct {
    NewLogin string  `json:"newLogin" validate:"required,min=5,max=32,login"`
}

type ChangeEmail struct {
    NewEmail string  `json:"newEmail" validate:"required,min=7,max=128,email"`
}

type ConfirmEmail struct {
    Code string  `json:"code" validate:"required,len=6,numeric"`
}

type CheckLoginValue struct {
    UserLogin string `json:"userLogin" validate:"required,min=5,max=32,login"`
}

type CheckEmailValue struct {
    UserEmail string `json:"userEmail" validate:"required,min=7,max=128,email"`
}
