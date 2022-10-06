/* This file is auto-generated, manual edits in this file will be overwritten! */
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/blink/skipr-test/app/contexts"
)

func GetTestId(c *gin.Context) {
	testId := c.GetHeader("testId")

	if testId != "" {
		// Save context
		err := contexts.SaveContext(c, func(context *contexts.Context) {
			context.TestId = &testId
		})
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}
	}

	c.Next()
}
