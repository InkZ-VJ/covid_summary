package dtos

type CovidRecord struct {
	Nation         string `json:"Nation"`
	District       string `json:"District"`
	ConfirmDate    string `json:"ConfirmDate"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceEn     string `json:"ProvinceEn"`
	No             int    `json:"No"`
	Age            int    `json:"Age"`
	ProvinceID     int    `json:"ProvinceId"`
	StatQuarantine int    `json:"StatQuarantine"`
}

type CovidResponse struct {
	Data []CovidRecord `json:"Data"`
}
