package middleware

import (
	"github.com/gin-gonic/gin"
)

var whitelist = [][2]string{
	{"GET", "^/api/v1/_hc"},
	{"POST", "^/api/v1/auth/login"},
	{"POST", "^/api/v1/tentant/getdomain"},
	{"POST", "^/api/v1/auth/refresh"},
}

type auth struct {
	Token          string `json:"token"`
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Role           int    `json:"role"`
	ProfilePicture string `json:"profile_picture"`
}

//AuthUserMiddleware - For authenticating user
func AuthUserMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {

	}
}

//AuthAdminMiddleware - For authenticating admin
func AuthAdminMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {

	}
}

//AuthMiddleware - For authenticating admin and user
func AuthMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {

	}

}

//AuthMiddleware - For authenticating admin and user
func AuthIsSubscribedMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {

	}

}

//AuthMiddleware - For authenticating admin and user and public
func AuthPublicMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	return func(c *gin.Context) {

	}

}
