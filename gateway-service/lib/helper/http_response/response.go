package http_response

// status code 200
type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
}

type Empty struct{}
