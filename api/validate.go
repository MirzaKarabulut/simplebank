package api

import (
	"github.com/MirzaKarabulut/simplebank/util"
	"github.com/go-playground/validator/v10"
)

// custom validator for the currency we support
var validCurrency validator.Func = func(FieldLevel validator.FieldLevel) bool {
	if currency, ok := FieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
