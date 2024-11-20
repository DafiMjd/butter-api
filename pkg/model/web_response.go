package model

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type SingleDoc struct {
	Doc interface{} `json:"doc"`
}

type MultiDocs struct {
	Docs interface{} `json:"docs"`
}
