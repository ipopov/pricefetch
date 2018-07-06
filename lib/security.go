package pricefetch

type Security struct {
	Name string
	// TODO: Convert this to decimal.
	Price float64
}

type SecurityFetcher interface {
	Run() ([]Security, error)
}
