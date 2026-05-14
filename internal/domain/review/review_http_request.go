package review

type ReqCreate struct {
	NewReview Review `json:"new_review"`
}

type ReqUpdate struct {
	Id       string `json:"id"`
	ToUpdate Review `json:"to_update"`
}
