package util

// constants for all supported currencies
const (
	INR = "INR"
	AED = "AED"
	USD = "USD"
	EUR = "EUR"
)

// isSupportedCurrencies returns true if the currency is supported
func IsSupportedCurrencies(currency string) bool {
	switch currency {
	case INR, AED, USD, EUR:
		return true
	}
	return false
}
