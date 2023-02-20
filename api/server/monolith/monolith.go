/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/go-kipi/currency/api"
	"github.com/go-kipi/currency/domain/application"
	"github.com/orchestd/dependencybundler/interfaces/validations"
	. "github.com/orchestd/servicereply"
)

type CurrencyInterface struct {
	currencyApp application.CurrencyApp
	validation  validations.Validations
}

func NewCurrencyInterface(currencyApp application.CurrencyApp, validation validations.Validations) CurrencyInterface {
	return CurrencyInterface{currencyApp: currencyApp, validation: validation}
}

func (i CurrencyInterface) GetQuote(c context.Context, req api.GetQuoteRequest) (api.GetQuoteResponse, ServiceReply) {
	if vErr := i.validation.Validate(req); vErr != nil && !vErr.IsSuccess() {
		return api.GetQuoteResponse{}, vErr
	}
	res, err := i.currencyApp.GetQuote(c, req.CurrencyId)
	return api.GetQuoteResponse(res), err
}
