package httplog

import (
	"net/http"
	"testing"

	"github.com/go-web/httpctx"
)

func TestError(t *testing.T) {
	r := &http.Request{}
	Error(r, "hello", "world")
	ctx := httpctx.Get(r)
	if ctx.Value(ErrorID) != "helloworld" {
		t.Fatalf("Unexpected value. Want \"helloworld\", have %v", ctx.Value("error"))
	}
}

func TestErrorf(t *testing.T) {
	r := &http.Request{}
	Errorf(r, "hello, world")
	ctx := httpctx.Get(r)
	if ctx.Value(ErrorID) != "hello, world" {
		t.Fatalf("Unexpected value. Want \"hello, world\", have %v", ctx.Value("error"))
	}
}

func TestErrorln(t *testing.T) {
	r := &http.Request{}
	Errorln(r, "hello", "world")
	ctx := httpctx.Get(r)
	if ctx.Value(ErrorID) != "hello world\n" {
		t.Fatalf("Unexpected value. Want \"hello world\\n\", have %v", ctx.Value("error"))
	}
}
