package rental

// CanChangeTo verifica se é possível mudar para um novo status
func (r *Rental) CanChangeTo(newStatus Status) bool {
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
