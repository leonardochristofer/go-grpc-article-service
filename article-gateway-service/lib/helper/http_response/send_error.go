package http_response

import (
	"context"
	"gateway-service/lib/helper/timestamp"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendError(err error, c *gin.Context) string {
	// set error codes, only if previous error doesnt have any
	httpStatus := http.StatusInternalServerError
	errMsg := err.Error()
	if st, ok := status.FromError(err); ok {
		httpStatus = STATUS_MAP[st.Code()]
		errMsg = st.Message()
	}

	response := Response{
		StatusCode: httpStatus,
		Message:    errMsg,
		Status:     STATUS_DESC[httpStatus],
		Timestamp:  timestamp.GetNow(),
		Data:       Empty{},
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Authorization, X-SKIP-AUTH")
	c.Header("Access-Control-Allow-Methods", "DELETE, POST, HEAD, PATCH, OPTIONS, GET, PUT")

	c.JSON(httpStatus, response)
	return errMsg
}

func SendErrorAborted(c *gin.Context) string {
	switch c.Request.Context().Err() {
	case context.Canceled:
		return SendError(status.Error(codes.Canceled, "Request Canceled"), c)
	case context.DeadlineExceeded:
		return SendError(status.Error(codes.DeadlineExceeded, "Deadline Exceeded"), c)
	default:
		return ""
	}
}

func SendAbortedUnauthenticate(c *gin.Context) string {
	return SendError(status.Error(codes.Unauthenticated, "Request Unauthenticate"), c)
}

func SendAbortedUnauthorized(c *gin.Context) string {
	return SendError(status.Error(codes.Unauthenticated, "Request Unauthorized"), c)
}

func AbortError(err error, c *gin.Context) {
	httpStatus := http.StatusInternalServerError
	errMsg := err.Error()
	if st, ok := status.FromError(err); ok {
		httpStatus = STATUS_MAP[st.Code()]
		errMsg = st.Message()
	}

	response := Response{
		StatusCode: httpStatus,
		Message:    errMsg,
		Status:     STATUS_DESC[httpStatus],
		Timestamp:  timestamp.GetNow(),
		Data:       Empty{},
	}
	c.AbortWithStatusJSON(httpStatus, response)
}

func AbortUnauthenticate(c *gin.Context) {
	AbortError(status.Error(codes.Unauthenticated, "Request Unauthenticate"), c)
}

func AbortUnauthorized(c *gin.Context) {
	AbortError(status.Error(codes.Unauthenticated, "Request Unauthorized"), c)
}
