package interfaces

type CurrencyRepo interface {
	GetIDR() (float64, error)
}
