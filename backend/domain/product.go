package domain

type Products []Product

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Publisher string `json:"publisher"`
	Price     string `json:"price"`
}
