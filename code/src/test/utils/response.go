package utils

type Base struct {
	RequestSuccessful bool      `json:"requestSuccessful"`
	ResponseMessage   string    `json:"responseMessage"`
	ResponseCode      string    `json:"responseCode"`
}

type RespAuth struct {
	Base
	ResponseBody      AuthLogin `json:"responseBody"`
}

type AuthLogin struct {
	AccessToken string `json:"accessToken"`
	ExpiresLn   int32  `json:"expiresLn"`
}

type RespBalance struct {
	Base
	ResponseBody Balance `json:"responseBody"`
}

type Balance struct {
	AvailableBalance float64 `json:"availableBalance"`
	LedgerBalance float64	`json:"ledgerBalance"`
}