package helper

import (
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(errCode codes.Code, label string, err error) error {
	// make sure that we didnt have nested error codes
	if st, ok := status.FromError(err); !ok {
		var errMsg string

		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" { // Duplicate data violation
				errMsg = "Duplicate data violation"

				if err.Table != "" {
					errMsg += " when trying to insert or update to : " + err.Table
				}
				return status.Error(errCode, errMsg+". Error details : "+label+err.Error())
			}
		}

		if err != nil {
			return status.Error(errCode, label+err.Error())
		} else {
			return status.Error(errCode, label)
		}
	} else {
		return status.Error(errCode, st.Message())
	}
}
