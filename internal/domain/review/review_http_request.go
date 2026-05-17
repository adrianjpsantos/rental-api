package review

type RequestCreateReviewInput struct {
	NewReview Review `json:"new_review"`
}

type RequestUpdateReviewInput struct {
	ReviewID string `json:"review_id"`
	ToUpdate Review `json:"to_update"`
}
