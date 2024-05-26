package payloads

type RequestTransaksiRekening struct {
	No      string  `json:"no_rekening"`
	Nominal float64 `json:"nominal"`
}
