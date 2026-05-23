package rental

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DeliveryMethod string

const (
	DeliveryPickup   DeliveryMethod = "pickup"   // Retirada no local
	DeliveryDelivery DeliveryMethod = "delivery" // Entrega pelo locador
	DeliveryShipping DeliveryMethod = "shipping" // Frete (transportadora)
)

type RentalStatus string

const (
	Pending   RentalStatus = "pending"
	Approved  RentalStatus = "approved"
	Active    RentalStatus = "active"
	Completed RentalStatus = "completed"
	Cancelled RentalStatus = "cancelled"
	Rejected  RentalStatus = "rejected"
)

type PaymentStatus string

const (
	PayPending   PaymentStatus = "pending"
	PayCompleted PaymentStatus = "completed"
	PayRefunded  PaymentStatus = "refunded"
	PayFailed    PaymentStatus = "failed"
)

type Rental struct {
	Id                 uuid.UUID
	ItemID             uuid.UUID
	LesseeID           uuid.UUID
	LessorID           uuid.UUID
	StartDate          time.Time
	EndDate            time.Time
	TotalAmount        float64
	Status             RentalStatus
	PaymentStatus      PaymentStatus
	DeliveryMethod     DeliveryMethod
	Notes              string
	CancellationReason string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	StartedAt          *time.Time
	CompletedAt        *time.Time
	CancelledAt        *time.Time
}

type RentalCreateInput struct {
	ItemID         uuid.UUID
	LesseeID       uuid.UUID
	LessorID       uuid.UUID
	StartDate      time.Time
	EndDate        time.Time
	TotalAmount    float64
	DeliveryMethod DeliveryMethod
	Notes          string
}

// NewRental cria uma nova solicitação de aluguel
func NewRental(input RentalCreateInput) (*Rental, error) {
	rental := &Rental{
		Id:             uuid.New(),
		ItemID:         input.ItemID,
		LesseeID:       input.LesseeID,
		LessorID:       input.LessorID,
		StartDate:      input.StartDate,
		EndDate:        input.EndDate,
		TotalAmount:    input.TotalAmount,
		Status:         Pending,
		PaymentStatus:  PayPending,
		DeliveryMethod: input.DeliveryMethod,
		Notes:          input.Notes,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := rental.Validate(); err != nil {
		return nil, err
	}

	return rental, nil
}

// UpdateStatus atualiza o status do aluguel
func (r *Rental) UpdateStatus(newStatus RentalStatus) error {
	if !r.CanChangeTo(newStatus) {
		return ErrInvalidStatusTransition
	}

	r.Status = newStatus
	r.UpdatedAt = time.Now()

	// Atualiza timestamps conforme o status
	switch newStatus {
	case Active:
		now := time.Now()
		r.StartedAt = &now
	case Completed:
		now := time.Now()
		r.CompletedAt = &now
	case Cancelled:
		now := time.Now()
		r.CancelledAt = &now
	}

	return r.Validate()
}

// CanChangeTo verifica se é possível mudar para um novo status
func (r *Rental) CanChangeTo(newStatus RentalStatus) bool {
	switch r.Status {
	case Pending:
		return newStatus == Approved || newStatus == Rejected || newStatus == Cancelled
	case Approved:
		return newStatus == Active || newStatus == Cancelled
	case Active:
		return newStatus == Completed || newStatus == Cancelled
	default:
		return false
	}
}

// IsActive retorna se o aluguel está em andamento
func (r *Rental) IsActive() bool {
	return r.Status == Active
}

// IsPending retorna se ainda está aguardando aprovação
func (r *Rental) IsPending() bool {
	return r.Status == Pending
}

func (r *Rental) Validate() error {
	fmt.Println("Validate() Rental Not Implemented")
	return nil
}

func ParseRentalStatus(value string) (*RentalStatus, error) {
	switch RentalStatus(value) {
	case Active, Rejected, Pending, Approved, Completed, Cancelled:
		rt := RentalStatus(value)
		return &rt, nil
	default:
		return nil, ErrInvalidStatus
	}
}
