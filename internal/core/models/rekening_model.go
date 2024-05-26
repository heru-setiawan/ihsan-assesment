package models

import "math/rand"

type Rekening struct {
	No    string  `json:"no_rekening" gorm:"primaryKey;"`
	Saldo float64 `json:"saldo"`

	NasabahNIK string `json:"-"`
}

func (m *Rekening) GenerateNo() {
	var numbers = []rune("0987654321")
	b := make([]rune, 20)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}

	m.No = string(b)
}
