package user

type ReqByEmail struct {
	Email string `json:"email"`
}

type ReqByCPF struct {
	Cpf string `json:"cpf"`
}

type ReqUpdate struct {
	ToUpdate UserUpdate `json:"to_update"`
}
type ReqCreate struct {
	NewUser User `json:"new_user"`
}
