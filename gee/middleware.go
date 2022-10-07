package gee

import (
	logger "github.com/amoghe/distillog"
	"time"
)

func Logger() HandleFunc {
	return func(ctx *Context) {
		t := time.Now()
		ctx.Next()
		// Calculate resolution time
		logger.Infof("[%d] %s in %v", ctx.StatusCode, ctx.Request.RequestURI, time.Since(t))
	}
}

func MiddlewareForV2() HandleFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		logger.Infof("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))

	}
}
