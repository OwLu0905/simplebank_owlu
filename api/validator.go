package api

import (
	"github.com/OwLu0905/simplebank_owlu/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {

		// TODO : check currency is support or not
		return util.IsSupportCurrency(currency)
	}

	return false
}
