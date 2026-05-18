package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	Number string
	Owner  string
}

type CryptoWallet struct {
	Address  string
	Currency string
}

type PayPal struct {
	Number string
	Owner  string
}

func (cc CreditCard) Process(amount float64) string {
	return fmt.Sprintf("Оплата картой %s на сумму %.2f выполнена", cc.Number, amount)
}

func (cw CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf("Оплата кошельком %s на сумму %.2f выполнена", cw.Currency, amount)
}

func (pp PayPal) Process(amount float64) string {
	return fmt.Sprintf("Оплата кошельком %s на сумму %.2f выполнена", pp.Number, amount)
}

func main() {
	processors := []PaymentProcessor{
		CreditCard{Number: "1234 5678 9012 3456", Owner: "Аскеров Алескер"},
		CryptoWallet{Address: "123QWERTY", Currency: "USTD"},
		PayPal{Number: "paypal@email.com", Owner: "Аскеров Алескер"},
	}
	for _, p := range processors {
		fmt.Println(p.Process(100.00))
	}
}
