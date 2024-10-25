package items

type CreateDataRequest struct {
	CategoryId  int     `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

type FilterListing struct {
	Page string `json:"page"`
	Size string `json:"size"`
}
