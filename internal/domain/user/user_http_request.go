package user

type ReqByEmail struct {
	Email string `json:"email"`
}

type ReqByCPF struct {
	Cpf string `json:"cpf"`
}

type ReqById struct {
	Id string `json:"id"`
}
