package rental

func (r *Rental) Validate() error {
	return nil
}

func ParseRentalStatus(value string) (*Status, error) {
	switch Status(value) {
	case Active, Rejected, Pending, Approved, Completed, Cancelled:
		rt := Status(value)
		return &rt, nil
	default:
		return nil, ErrInvalidStatus
	}
}
