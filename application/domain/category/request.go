package category

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type FilterListing struct {
	Page string `json:"page"`
	Size string `json:"size"`
}
