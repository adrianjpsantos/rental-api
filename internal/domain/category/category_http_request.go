package category

type RequestCreateCategory struct {
	NewCategory CategoryCreateInput `json:"new_category"`
}

type RequestExistsCategory struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type RequestUpdateCategory struct {
	ToUpdate CategoryUpdate `json:"to_update"`
}
