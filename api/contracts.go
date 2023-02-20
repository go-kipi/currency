/* This file is auto-generated and should not be modified */

package api

import (
	"context"
	. "github.com/orchestd/servicereply"
)

type CurrencyApi interface {
	GetQuote(c context.Context, req GetQuoteRequest) (GetQuoteResponse, ServiceReply)
}
