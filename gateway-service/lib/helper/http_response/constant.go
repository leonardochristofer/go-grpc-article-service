package http_response

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	STANDARD_200_STATUS = "OK"
	STANDARD_201_STATUS = "Created"
	STANDARD_400_STATUS = "Bad Request"
	STANDARD_401_STATUS = "Unauthorized"
	STANDARD_403_STATUS = "Forbidden"
	STANDARD_404_STATUS = "Not Found"
	STANDARD_409_STATUS = "Conflict"
	STANDARD_415_STATUS = "Unsupported Media Type"
	STANDARD_422_STATUS = "Unprocessable Entity"
	STANDARD_500_STATUS = "Internal Server Error"
)

var STATUS_MAP = map[codes.Code]int{
	codes.InvalidArgument: http.StatusBadRequest,
	codes.Aborted:         http.StatusGone,
	codes.Canceled:        http.StatusGone,
	codes.NotFound:        http.StatusNotFound,
	codes.Unauthenticated: http.StatusUnauthorized,
	codes.Internal:        http.StatusInternalServerError,
	codes.OK:              http.StatusOK,
	codes.Unavailable:     http.StatusServiceUnavailable,
}

var STATUS_DESC = map[int]string{
	http.StatusBadRequest:          "Bad Request",
	http.StatusGone:                "Request Aborted/Canceled",
	http.StatusNotFound:            "Not Found",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusInternalServerError: "Internal Server Error",
	http.StatusOK:                  "OK",
	http.StatusServiceUnavailable:  "Service Unavailable",
}
