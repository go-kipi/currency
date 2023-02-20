package defaultApp

import (
	"context"
	"github.com/go-kipi/currency/domain/entity"
	. "github.com/orchestd/servicereply"
)

func (app currencyApp) GetQuote(c context.Context, currencyId string) (entity.Currency, ServiceReply) {
	r := entity.Currency{}
	return r, nil
}
