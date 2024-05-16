package api

// validator: Creates a custom validator for currencies in require json
import (
	"github.com/ekefan/bank_panda/utils"
	validator "github.com/go-playground/validator/v10"
)


var validCurrency validator.Func  = func(fl validator.FieldLevel) bool {
	//convert the unpack the field from the reflect value and convert it to a string
	if currency, ok := fl.Field().Interface().(string); ok {
		// check if currency is supported
		return utils.IsSupportedCurrency(currency)
	}
	return false
}