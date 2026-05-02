package rental

import (
	"time"

	"github.com/google/uuid"
)

type DeliveryMethod string

const (
	DeliveryPickup   DeliveryMethod = "pickup"   // Retirada no local
	DeliveryDelivery DeliveryMethod = "delivery" // Entrega pelo locador
	DeliveryShipping DeliveryMethod = "shipping" // Frete (transportadora)
)

type Status string

const (
	Pending   Status = "pending"
	Approved  Status = "approved"
	Active    Status = "active"
	Completed Status = "completed"
	Cancelled Status = "cancelled"
	Rejected  Status = "rejected"
)

type PaymentStatus string

const (
	PayPending   PaymentStatus = "pending"
	PayCompleted PaymentStatus = "completed"
	PayRefunded  PaymentStatus = "refunded"
	PayFailed    PaymentStatus = "failed"
)

type Rental struct {
	Id                 uuid.UUID      `gorm:"type:uuid;primaryKey"`
	ItemID             uuid.UUID      `gorm:"type:uuid;not null;index"`
	LesseeID           uuid.UUID      `gorm:"type:uuid;not null;index"`
	LessorID           uuid.UUID      `gorm:"type:uuid;not null;index"`
	StartDate          time.Time      `gorm:"not null"`
	EndDate            time.Time      `gorm:"not null"`
	TotalAmount        float64        `gorm:"not null;type:decimal(12,2)"`
	Status             Status         `gorm:"size:30;not null;index"`
	PaymentStatus      PaymentStatus  `gorm:"size:30;not null"`
	DeliveryMethod     DeliveryMethod `gorm:"size:50"`
	Notes              string         `gorm:"type:text"`
	CancellationReason string         `gorm:"type:text"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	StartedAt          *time.Time
	CompletedAt        *time.Time
	CancelledAt        *time.Time
}

// NewRental cria uma nova solicitação de aluguel
func NewRental(itemID, lesseeID, lessorID uuid.UUID, startDate, endDate time.Time, totalAmount float64, deliveryMethod DeliveryMethod) (*Rental, error) {
	rental := &Rental{
		Id:             uuid.New(),
		ItemID:         itemID,
		LesseeID:       lesseeID,
		LessorID:       lessorID,
		StartDate:      startDate,
		EndDate:        endDate,
		TotalAmount:    totalAmount,
		Status:         Pending,
		PaymentStatus:  PayPending,
		DeliveryMethod: deliveryMethod,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := rental.Validate(); err != nil {
		return nil, err
	}

	return rental, nil
}

// UpdateStatus atualiza o status do aluguel
func (r *Rental) UpdateStatus(newStatus Status) error {
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
