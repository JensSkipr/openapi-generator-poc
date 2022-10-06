package middlewares

import (
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func InitialiseAuthMiddleware() IAuthMiddleware {
	return AuthMiddleware{}
}

func (middleware AuthMiddleware) Authorize(f func(c *gin.Context), roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// // Gets token from header
		// token := c.GetHeader("Authorization")

		// // Initialize default app
		// opt := option.WithCredentialsFile("./firebase.json")
		// app, err := firebase.NewApp(context.Background(), nil, opt)
		// if err != nil {
		// 	c.IndentedJSON(http.StatusInternalServerError, SERVER_ERROR)
		// 	return
		// }

		// // Access auth service from the default app
		// client, err := app.Auth(context.Background())
		// if err != nil {
		// 	c.IndentedJSON(http.StatusUnauthorized, SERVER_ERROR)
		// 	return
		// }

		// // Verifies that the bearer token is included
		// token, err = verifyToken(token)
		// if err != nil {
		// 	c.IndentedJSON(http.StatusUnauthorized, err.Error())
		// 	return
		// }

		// // Validates token
		// parsedToken, err := client.VerifyIDToken(context.Background(), token)
		// if err != nil {
		// 	c.IndentedJSON(http.StatusUnauthorized, AUTH_INVALID_TOKEN)
		// 	return
		// }

		// // Verifies that the token is not nil
		// if parsedToken == nil {
		// 	c.IndentedJSON(http.StatusUnauthorized, AUTH_INVALID_TOKEN)
		// 	return
		// }

		// // Verifies claims
		// if len(roles) > 0 {
		// 	roleClaims := parsedToken.Claims["roles"]

		// 	if roleClaims == nil {
		// 		c.IndentedJSON(http.StatusForbidden, AUTH_MISSING_PERMISSIONS)
		// 		return
		// 	}

		// 	roleString := fmt.Sprintf("%v", roleClaims)
		// 	tokenRoles := strings.Split(roleString, ",")

		// 	// Verifies that at least one role is in common
		// 	if !lo.Some(tokenRoles, roles) {
		// 		c.IndentedJSON(http.StatusForbidden, AUTH_MISSING_PERMISSIONS)
		// 		return
		// 	}

		// 	// Verifies that the account has been confirmed
		// 	emailVerified := parsedToken.Claims["email_verified"].(bool)
		// 	if !emailVerified {
		// 		c.IndentedJSON(http.StatusForbidden, AUTH_UNVERIFIED_EMAIL)
		// 		return
		// 	}

		// 	// Save role in context
		// 	if err := contexts.SaveContext(c, func(context *contexts.Context) {
		// 		context.Roles = tokenRoles
		// 	}); err != nil {
		// 		c.IndentedJSON(http.StatusInternalServerError, err)
		// 		return
		// 	}
		// }

		// // Adds the sub to the context
		// if err := contexts.SaveContext(c, func(context *contexts.Context) {
		// 	context.Sub = &parsedToken.Subject
		// }); err != nil {
		// 	c.IndentedJSON(http.StatusInternalServerError, err)
		// 	return
		// }

		// Calls the next handler in chain
		f(c)
	}
}

// Verifies that given token is not empty and is of correct format
// func verifyToken(token string) (resp string, error error) {
// 	// Verifies that a token has been given
// 	if token == "" {
// 		return "", errors.New(AUTH_MISSING_TOKEN)
// 	}

// 	// Verifies that the token starts with bearer
// 	if !strings.Contains(strings.ToLower(token), "bearer") {
// 		return "", errors.New(AUTH_WRONG_TOKEN_FORMAT)
// 	}

// 	// Returns the substring containing only the token
// 	return token[7:], nil
// }
