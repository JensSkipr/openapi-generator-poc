/* This file is auto-generated, manual edits in this file will be overwritten! */
package middlewares

type Middleware struct {
	AuthMiddleware IAuthMiddleware
}

func Initialise() Middleware {
	return Middleware{
		AuthMiddleware: InitialiseAuthMiddleware(),
	}
}
