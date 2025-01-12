package market

import (
	"github.com/funtoy/huobi_golang/pkg/model/base"
	"github.com/shopspring/decimal"
)

type SubscribeCandlestickResponse struct {
	base    base.WebSocketResponseBase
	Channel string `json:"ch"`
	Tick    *Tick
	Data    []Tick
}

func (receiver SubscribeCandlestickResponse) GetChannel() string {
	return receiver.Channel
}

type Tick struct {
	Id     int64           `json:"id"`
	Amount decimal.Decimal `json:"amount"`
	Count  int             `json:"count"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	Low    decimal.Decimal `json:"low"`
	High   decimal.Decimal `json:"high"`
	Vol    decimal.Decimal `json:"vol"`
}
