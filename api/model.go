/* This file is auto-generated and should not be modified */

package api

import (
	"github.com/go-kipi/currency/domain/entity"
)

type Currency entity.Currency

type GetQuoteRequest struct {
	CurrencyId string `json:"currencyId" validate:"required"`
}

type GetQuoteResponse Currency

const (
	GetQuoteMethod = "getQuote"
)
