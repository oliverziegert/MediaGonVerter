package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
)

var httpStatusCodeMapping = map[int]int{
	e.AuthUnknown:            http.StatusUnauthorized,
	e.AuthDataInvalid:        http.StatusUnauthorized,
	e.AuthCredentialsInvalid: http.StatusUnauthorized,
	e.AuthUserNotActivated:   http.StatusPreconditionFailed,

	e.PermUnknown:             http.StatusForbidden,
	e.PermGetUserData:         http.StatusForbidden,
	e.PermChangeUserData:      http.StatusForbidden,
	e.PermGetUserAccount:      http.StatusForbidden,
	e.PermChangeUserAccount:   http.StatusForbidden,
	e.PermGetEntryCharacts:    http.StatusForbidden,
	e.PermChangeEntryCharacts: http.StatusForbidden,
	e.PermGetAllEntries:       http.StatusForbidden,
	e.PermChangeAllEntries:    http.StatusForbidden,
	e.PermGetOwnEntries:       http.StatusForbidden,
	e.PermChangeOwnEntries:    http.StatusForbidden,

	e.ValUnknown:              http.StatusBadRequest,
	e.ValJsonInvalid:          http.StatusBadRequest,
	e.ValPageNumberInvalid:    http.StatusBadRequest,
	e.ValIdInvalid:            http.StatusBadRequest,
	e.ValFilterInvalid:        http.StatusBadRequest,
	e.ValSortInvalid:          http.StatusBadRequest,
	e.ValOffsetInvalid:        http.StatusBadRequest,
	e.ValLimitInvalid:         http.StatusBadRequest,
	e.ValFieldNil:             http.StatusBadRequest,
	e.ValNumberNegative:       http.StatusBadRequest,
	e.ValNumberNegativeOrZero: http.StatusBadRequest,
	e.ValStringEmpty:          http.StatusBadRequest,
	e.ValStringTooLong:        http.StatusBadRequest,
	e.ValDateInvalid:          http.StatusBadRequest,
	e.ValTimestampInvalid:     http.StatusBadRequest,
	e.ValArrayEmpty:           http.StatusBadRequest,
	e.ValRoleInvalid:          http.StatusBadRequest,
	e.ValUsernameInvalid:      http.StatusBadRequest,
	e.ValPasswordInvalid:      http.StatusBadRequest,

	e.LogicEntryNotFound:                  http.StatusNotFound,
	e.LogicEntryTypeNotFound:              http.StatusNotFound,
	e.LogicEntryActivityNotFound:          http.StatusNotFound,
	e.LogicEntryActivityDeleteNotAllowed:  http.StatusConflict,
	e.LogicEntryTimeIntervalInvalid:       http.StatusBadRequest,
	e.LogicEntryBreakDurationTooLong:      http.StatusBadRequest,
	e.LogicEntrySearchDateIntervalInvalid: http.StatusBadRequest,
	e.LogicRoleNotFound:                   http.StatusNotFound,
	e.LogicUserNotFound:                   http.StatusNotFound,
	e.LogicUserAlreadyExists:              http.StatusConflict,
}

func getHttpStatusCode(errorCode int) int {
	sc, ok := httpStatusCodeMapping[errorCode]
	if !ok {
		l.Debugf("Unexpected error code %d. Using status code 500.", errorCode)
		return http.StatusInternalServerError
	}
	return sc
}

// ErrorMiddleware creates the security context.
type ErrorMiddleware struct{}

// NewErrorMiddleware  create a new ErrorMiddleware
func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

// ServeHTTP processes requests.
func (m *ErrorMiddleware) ServeHTTP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l.Verb("Before API error check.")

		defer func() {
			if err := recover(); err != nil {
				l.Verb("Catching error.")
				m.handleError(ctx, err)
			}
		}()

		ctx.Next()

		l.Verb("After API error check.")
	}
}

func (m *ErrorMiddleware) handleError(ctx *gin.Context, err interface{}) {
	switch eType := err.(type) {
	case *e.Error:
		// We can retrieve the status here and write out a specific status code.
		sc := getHttpStatusCode(eType.Code)
		er := e.NewError(eType.Code, eType.Message)
		ctx.AbortWithError(sc, er)
	default:
		// Any error types we don't specifically look out for default to serving a HTTP 500
		l.Errorf("%s", eType)
		sc := http.StatusInternalServerError
		er := e.NewError(e.SysUnknown, "Internal Server Error")
		ctx.AbortWithError(sc, er)
	}
}
