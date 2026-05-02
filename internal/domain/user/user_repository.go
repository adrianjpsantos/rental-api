package user

import "github.com/google/uuid"

type InterfaceUserRepository interface {
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByCPF(cpf string) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error // soft delete
	ExistsByEmail(email string) (bool, error)
	ExistsByCPF(cpf string) (bool, error)
	UpdateReputation(userID uuid.UUID, newRating float32) error
}
