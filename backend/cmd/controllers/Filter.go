package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context) {
	tokenCookie, _ := c.Request.Cookie("token")
	fmt.Println(tokenCookie.Value)
	//if err != nil {
	//	if err == http.ErrNoCookie {
	//		// If the cookie is not set, return an unauthorized status
	//		c.Writer.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	// For any other type of error, return a bad request status
	//	c.Writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//
	//// Get the JWT string from the cookie
	//tokenString := tokenCookie.Value
	//
	//// Initialize a new instance of `Claims`
	//claims := jwt.MapClaims{}
	//
	//
	//// Parse the JWT string and store the result in `claims`.
	//// Note that we are passing the key in this method as well. This method will return an error
	//// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	//// or if the signature does not match
	//token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	//	return "huy", nil
	//})
	//if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		c.Writer.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	c.Writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//if !token.Valid {
	//	c.Writer.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	//
	//// Finally, return the welcome message to the user, along with their
	//// username given in the token
	//c.Writer.Write([]byte(fmt.Sprintf("Welcome %s!", claims)))
}
