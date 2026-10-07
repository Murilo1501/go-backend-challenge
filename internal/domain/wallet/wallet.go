package wallet

import (
	"backend-challenge-go/internal/domain/money"
	"time"
)

type Wallet struct {
	Id       string
	PlayerId string
	Currency string
	Balance  money.Money

	CreatedAt time.Time
	UpdatedAt time.Time
}
