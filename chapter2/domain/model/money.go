package model

import "errors"

type money struct {
	amount   int
	currency string
}

// NewMoney コンストラクタ
func NewMoney(amount int, currency string) *money {
	return &money{
		amount:   amount,
		currency: currency,
	}
}

// Amount getter
func (m *money) Amount() int {
	return m.amount
}

// Currency getter
func (m *money) Currency() string {
	return m.currency
}

// Add 金額を加算する
func (m *money) Add(other *money) (*money, error) {
	if m.currency != other.currency {
		return nil, errors.New("通貨が異なります。")
	}
	return NewMoney(m.amount+other.amount, m.currency), nil
}
