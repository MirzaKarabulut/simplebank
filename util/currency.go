package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// check the currency that we support
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}