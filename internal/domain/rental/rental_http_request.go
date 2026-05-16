package rental

type ReqCreate struct {
	NewRental Rental `json:"new_rental"`
}

type ReqUpdateStatus struct {
	NewStatus Status `json:"new_status"`
}

type ReqCancel struct {
	CancellationReason string `json:"cancellation_reason"`
}
