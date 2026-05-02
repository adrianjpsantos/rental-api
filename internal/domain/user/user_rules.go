package user

// IsLessor retorna se o usuário é locador
func (u *User) IsLessor() bool {
	return u.Role == Lessor || u.Role == Admin
}

// IsLessee retorna se o usuário é locatário
func (u *User) IsLessee() bool {
	return u.Role == Lessee || u.Role == Admin
}

// IsAdmin retorna se o usuário é administrador
func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

// CanCreateItem verifica se o usuário pode cadastrar itens para alugar
func (u *User) CanCreateItem() bool {
	return u.IsLessor() || u.IsAdmin()
}

// CanRentItem verifica se o usuário pode alugar itens
func (u *User) CanRentItem() bool {
	return u.IsLessee() || u.IsAdmin()
}
