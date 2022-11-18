package gin_

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerFunc used to make adapter from handler.
type HandlerFunc func(request Request) (response Response, err error)

type Request struct {
	*http.Request
	GinCtx *gin.Context
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
	Message    map[string]string
	Meta, Data interface{}
}

// CaptureError is used when calling gin_.CaptureErrors.
type CaptureError struct {
	StatusCode  int
	Msg         map[string]string
	Err         error
	UserMsg     string
	InternalMsg string
	MoreInfo    string
}

func (x *CaptureError) Unwrap() error { return x.Err }

func (x *CaptureError) Error() string {
	return fmt.Sprintf("scode: [%d] user:[%s] internal:[%s] more:[%s] unwrap:[%s]",
		x.StatusCode,
		x.UserMsg,
		x.InternalMsg,
		x.MoreInfo,
		x.Err.Error(),
	)
}

func NewResponseError(err error, sCode int, msg map[string]string, uMsg, iMsg, more string) (Response, error) {
	return Response{sCode, nil, nil, msg, nil, nil}, &CaptureError{sCode, msg, err, uMsg, iMsg, more}
}

func UnwrapFirstError(err error) string {
	return UnwrapAll(err).Error()
}

// UnwrapAll will unwrap the underlying error until we get the first wrapped error.
func UnwrapAll(err error) error {
	for err != nil && errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	return err
}
