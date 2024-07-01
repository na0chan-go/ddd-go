package model

import (
	"testing"
)

func TestNewMoney(t *testing.T) {
	amount := 100
	currency := "USD"
	m := NewMoney(amount, currency)

	if m.Amount() != amount {
		t.Errorf("Expected amount %d, got %d", amount, m.Amount())
	}

	if m.Currency() != currency {
		t.Errorf("Expected currency %s, got %s", currency, m.Currency())
	}
}

func TestMoney_Add(t *testing.T) {
	m1 := NewMoney(100, "USD")
	m2 := NewMoney(200, "USD")
	expectedAmount := 300

	sum, err := m1.Add(m2)
	if err != nil {
		t.Fatalf("Add() error = %v, wantErr %v", err, false)
	}

	if sum.Amount() != expectedAmount {
		t.Errorf("Expected sum amount %d, got %d", expectedAmount, sum.Amount())
	}

	// 異なる通貨を加算しようとした場合のテスト
	m3 := NewMoney(100, "EUR")
	_, err = m1.Add(m3)
	if err == nil {
		t.Errorf("Expected error when adding different currencies, got nil")
	}

	expectedErrMsg := "通貨が異なります。"
	if err != nil && err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}
}
