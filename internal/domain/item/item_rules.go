package item

// IsAvailable verifica se o item tem quantidade disponível
func (i *Item) IsAvailableToRent(quantityRequested int) bool {
	return i.IsActive && i.Quantity >= quantityRequested
}

// CanBeRented verifica regras básicas para aluguel
func (i *Item) CanBeRented() error {
	if !i.IsActive {
		return ErrItemNotActive
	}
	if i.Quantity <= 0 {
		return ErrItemOutOfStock
	}
	return nil
}
