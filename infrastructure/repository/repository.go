package repository

import (
	"github.com/go-kipi/currency/domain/repository"
)

type currencyRepository struct {
}

func NewCurrencyRepository() repository.CurrencyRepository {
	return currencyRepository{}
}
