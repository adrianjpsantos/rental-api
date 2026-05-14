package user

type ReqByEmail struct {
	Email string `json:"email"`
}

type ReqByCPF struct {
	Cpf string `json:"cpf"`
}

type ReqUpdate struct {
	Id       string     `json:"id"`
	ToUpdate UserUpdate `json:"to_update"`
}
type ReqCreate struct {
	NewUser User `json:"new_user"`
}
