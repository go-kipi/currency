/* This file is auto-generated and should not be modified */

package application

import (
	"context"
	"github.com/go-kipi/currency/domain/entity"
	. "github.com/orchestd/servicereply"
)

type CurrencyApp interface {
	GetQuote(c context.Context, currencyId string) (entity.Currency, ServiceReply)
}
