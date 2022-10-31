package xhttp

type RequestType string

const (
	POST                 = "POST"
	GET                  = "GET"
	TypeJSON RequestType = "json"
	TypeXML  RequestType = "XML"
	TypeForm RequestType = "form"
)

var types = map[RequestType]string{
	TypeJSON: "application/json;charset=UTF-8",
	TypeXML: "application/xml;charset=UTF-8",
	TypeForm: "application/x-www-form-urlencoded;charset=UTF-8",
}
