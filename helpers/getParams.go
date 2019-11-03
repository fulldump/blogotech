package helpers

import (
	"context"
	"net/http"

	"github.com/fulldump/box"
)

// GetParams return parameters map ready to be used
func GetParams(ctx context.Context) map[string]string {

	c := ctx.Value("box_context")
	if c == nil {
		return map[string]string{}
	}

	return c.(*box.C).Parameters
}

// SetBoxContext set a box context (a struct stored inside a go Context)
func SetBoxContext(ctx context.Context, c *box.C) context.Context {
	return context.WithValue(ctx, "box_context", c)
}

// GetBoxContext return box context (a struct stored inside a go Context)
func GetBoxContext(ctx context.Context) *box.C {
	c := ctx.Value("box_context")
	if c == nil {
		panic("box context should be in context")
	}

	return c.(*box.C)
}

// GetResponse return current http response (from context)
func GetResponse(ctx context.Context) http.ResponseWriter {
	return GetBoxContext(ctx).Response
}
