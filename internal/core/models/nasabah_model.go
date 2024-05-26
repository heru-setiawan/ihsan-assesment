package models

type Nasabah struct {
	NIK  string `json:"nik" gorm:"primaryKey;"`
	Nama string `json:"nama"`
	NoHp string `json:"no_hp" gorm:"unique;"`

	Rekening Rekening `json:"-" gorm:"foreignKey:NasabahNIK"`
}

func (m *Nasabah) NewRekening() {
	newRekening := new(Rekening)
	newRekening.GenerateNo()
	newRekening.NasabahNIK = m.NIK

	m.Rekening = *newRekening
}
