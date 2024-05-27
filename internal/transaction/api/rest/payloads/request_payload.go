package payloads

type RequestRegistration struct {
	PIN string `json:"pin"`
}

type RequestTransaction struct {
	Number string  `json:"no_rekening"`
	Amount float64 `json:"nominal"`
}
