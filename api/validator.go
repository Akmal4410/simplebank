package api

import (
	"github.com/akmal4410/simple_bank/util"
	"github.com/go-playground/validator/v10"
)

var validateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check currency is supported or not
		return util.IsSupportedCurrencies(currency)
	}
	return false
}
