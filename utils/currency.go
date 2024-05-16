//used in validator.go where the validator needs to verify if the currency is supported or not
package utils

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

//IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
