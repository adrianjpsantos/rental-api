package rental

type RequestCreateRentalInput struct {
	NewRental Rental `json:"new_rental"`
}

type RequestUpdateStatusRentalInput struct {
	NewStatus Status `json:"new_status"`
}

type RequestCancelRentalInput struct {
	CancellationReason string `json:"cancellation_reason"`
}
