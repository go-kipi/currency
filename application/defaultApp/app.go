package defaultApp

import (
	"github.com/go-kipi/currency/domain/application"
)

type currencyApp struct {
}

func NewCurrencyApp() application.CurrencyApp {
	return &currencyApp{}
}
