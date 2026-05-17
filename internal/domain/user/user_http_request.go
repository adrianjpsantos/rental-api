package user

type RequestGetByEmailInput struct {
	Email string `json:"email"`
}

type RequestGetByCpfInput struct {
	Cpf string `json:"cpf"`
}

type RequestUpdateUserInput struct {
	ToUpdate UserUpdate `json:"to_update"`
}
type RequestCreateUserInput struct {
	NewUser UserCreateInput `json:"new_user"`
}
