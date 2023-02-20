/* This file is auto-generated and should not be modified */

package http

import (
	"context"
	"github.com/go-kipi/currency/api"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	. "github.com/orchestd/servicereply"
)

const serviceName = "currency"

type currencyHTTPClient struct {
	client transport.HttpClient
}

func NewCurrencyApiHTTPClient(client transport.HttpClient) api.CurrencyApi {
	return currencyHTTPClient{client: client}
}

func (h currencyHTTPClient) GetQuote(c context.Context, req api.GetQuoteRequest) (api.GetQuoteResponse, ServiceReply) {
	var res api.GetQuoteResponse
	if reply := h.client.Call(c, req, serviceName, api.GetQuoteMethod, &res, nil); !reply.IsSuccess() {
		return res, reply
	}
	return res, nil
}
