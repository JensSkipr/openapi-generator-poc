/* This file is auto-generated, manual edits in this file will be overwritten! */
package middlewares

import "github.com/gin-gonic/gin"

type SortMiddleware struct {
	// Order of the execution
	Order int
	// The handler to call
	handler gin.HandlerFunc
}

// List of base middlewares
var MiddlewaresBase = []SortMiddleware{
	{
		Order:   100,
		handler: GetTestId,
	},
}
