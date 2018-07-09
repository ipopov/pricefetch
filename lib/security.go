package pricefetch

type Security struct {
	Name string
	// We're not doing arithmetic on these, and don't expect to need
	// more than four digits of decimal precision. So there's no need to
	// use a decimal (fixed point) type.
	Price float64
}

type SecurityFetcher interface {
	Run() ([]Security, error)
}
