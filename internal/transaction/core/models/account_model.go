package models

import "math/rand"

type Account struct {
	Number string `json:"no_rekening"`
	PIN    string `json:"-"`

	Balance float64 `json:"saldo"`
}

func (a *Account) GenerateNumber() {
	var numbers = []rune("0987654321")
	b := make([]rune, 20)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}

	a.Number = string(b)
}
