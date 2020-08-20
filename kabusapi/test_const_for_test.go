package kabusapi_test

const apiToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

var commonHeader = map[string]string{
	"Content-Type": "application/json",
}

var commonHeaderWithAPIKey = map[string]string{
	"X-API-KEY":    apiToken,
	"Content-Type": "application/json",
}
