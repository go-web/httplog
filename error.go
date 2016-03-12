package httplog

import (
	"fmt"
	"net/http"

	"github.com/go-web/httpctx"
	"golang.org/x/net/context"
)

// ErrorType is the type used for storing error in a context.Context.
type ErrorType int

// ErrorID is the key used for storing and retrieving errors to/from context.
var ErrorID ErrorType

// Error associates message v with the request context.
func Error(r *http.Request, v ...interface{}) {
	ctx := httpctx.Get(r)
	ctx = context.WithValue(ctx, ErrorID, fmt.Sprint(v...))
	httpctx.Set(r, ctx)
}

// Errorf associates message v with the request context.
func Errorf(r *http.Request, format string, v ...interface{}) {
	ctx := httpctx.Get(r)
	ctx = context.WithValue(ctx, ErrorID, fmt.Sprintf(format, v...))
	httpctx.Set(r, ctx)
}

// Errorln associates message v with the request context.
func Errorln(r *http.Request, v ...interface{}) {
	ctx := httpctx.Get(r)
	ctx = context.WithValue(ctx, ErrorID, fmt.Sprintln(v...))
	httpctx.Set(r, ctx)
}
