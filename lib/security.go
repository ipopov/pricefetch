package pricefetch

type Security interface {
  GetPrice() (float64, error)
  GetName() string
}
