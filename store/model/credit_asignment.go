package model

type CreditAssignment struct {
	Total                int     `json:"total"`
	SuccessTotal         int     `json:"exito_total"`
	SuccessTotalNot      int     `json:"no_exito_total"`
	InvestmentSuccess    float32 `json:"inversion_exito"`
	InvestmentSuccessNot float32 `json:"inversion_no_exito"`
}
