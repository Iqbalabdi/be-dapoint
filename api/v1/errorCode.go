package v1

import (
	dapoint_api "dapoint-api/error"
	"net/http"
)

func GetErrorStatus(err error) int {
	switch err {
	case dapoint_api.ErrBadRequest:
		return http.StatusBadRequest
	case dapoint_api.ErrInternalServer:
		return http.StatusInternalServerError
	case dapoint_api.ErrNotFound:
		return http.StatusNotFound
	case dapoint_api.ErrUnauthorized:
		return http.StatusUnauthorized
	}
	return http.StatusOK
}
