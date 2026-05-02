package item

import "github.com/google/uuid"

// Validate valida as regras de negócio da entidade Item
func (i *Item) Validate() error {
	if i.OwnerID == uuid.Nil {
		return ErrInvalidOwnerID
	}
	if i.CategoryID == uuid.Nil {
		return ErrInvalidCategoryID
	}
	if len(i.Title) < 5 || len(i.Title) > 100 {
		return ErrInvalidTitle
	}
	if i.PricePerDay <= 0 {
		return ErrPricePerDayTooLow
	}
	if i.Quantity <= 0 {
		return ErrQuantityZeroOrNegative
	}

	// Valida condição
	validConditions := map[string]bool{
		"Novo": true, "Usado": true, "Semi-novo": true, "Reformado": true,
	}
	if !validConditions[i.Condition] {
		return ErrInvalidCondition
	}

	return nil
}
