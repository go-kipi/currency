/* This file is auto-generated and should not be modified */

package http

import (
	"github.com/go-kipi/currency/api"
	"github.com/go-kipi/currency/api/server/monolith"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/transport"
)

func InitHandlers(router transport.IRouter, m monolith.CurrencyInterface, c configuration.Config) {
	router.POST(api.GetQuoteMethod, transport.HandleFunc(m.GetQuote))
}
