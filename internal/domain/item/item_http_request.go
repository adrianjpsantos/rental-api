package item

type ReqCreate struct {
	NewItem ItemCreateInput `json:"new_item"`
}

type ReqUpdate struct {
	ToUpdate ItemUpdate `json:"to_update"`
}

type ReqListByFilter struct {
	Filter ItemFilter `json:"filter"`
}
