/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/go-kipi/currency/api"
	"github.com/go-kipi/currency/api/server/monolith"
	. "github.com/orchestd/servicereply"
)

type CurrencyMonolithClient struct {
	MonolithInterface monolith.CurrencyInterface
}

func NewCurrencyMonolithClient(monolithInterface monolith.CurrencyInterface) api.CurrencyApi {
	return CurrencyMonolithClient{MonolithInterface: monolithInterface}
}

func (p CurrencyMonolithClient) GetQuote(c context.Context, req api.GetQuoteRequest) (api.GetQuoteResponse, ServiceReply) {
	return p.MonolithInterface.GetQuote(c, req)
}
