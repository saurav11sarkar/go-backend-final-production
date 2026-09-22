package product

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
	CategoryID  string  `json:"categoryId"`
	CreatedAt   string  `json:"createdAt"`
}
