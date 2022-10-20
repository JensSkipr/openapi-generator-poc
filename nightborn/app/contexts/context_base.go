/* This file is auto-generated, manual edits in this file will be overwritten! */
package contexts

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type ContextBase struct {
	Sub    *string
	Roles  []string
	TestId *string
}

func (context Context) HasRole(role string) bool {
	return context.Roles != nil && lo.Contains(context.Roles, role)
}

func (context Context) HasSomeRoles(roles []string) bool {
	return context.Roles != nil && lo.Some(context.Roles, roles)
}

func GetContext(c *gin.Context) (*Context, error) {
	contextFromGin, ok := c.Get("context")
	var context Context
	if ok {
		context, ok = contextFromGin.(Context)
		if !ok {
			return nil, errors.New(INVALID_CONTEXT)
		}
	}

	return &context, nil
}

func SaveContext(c *gin.Context, newContext func(context *Context)) error {
	context, err := GetContext(c)
	if err != nil {
		return err
	}

	// Get new context
	newContext(context)

	// Save new context
	c.Set("context", *context)

	c.Next()

	return nil
}
