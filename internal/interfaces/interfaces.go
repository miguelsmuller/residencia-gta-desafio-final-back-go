package interfaces

type Restaurant struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Image       string `json:"image`
}
type Product struct {
	ID            int     `json:"id"`
	ID_Restaurant int     `json:"name"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float32 `json:"price"`
	Image         string  `json:"image"`
}
type Additional struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
