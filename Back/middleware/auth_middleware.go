package middleware

import (
	"api-dashboard/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc"

	"github.com/gin-gonic/gin"
)

//Authorize middleware checks JWT Token
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		provider, err := oidc.NewProvider(c, "https://login.microsoftonline.com/5ab9af9b-4534-4c31-8e50-1e098461481c/v2.0")
		if err != nil {
			log.Println((err))
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Auth": "Error getting provider",
			})
			return
		}

		rawIDToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

		verifier := provider.Verifier(&oidc.Config{ClientID: setting.AppSetting.ClientID})

		// Parse and verify ID Token payload.
		idToken, err := verifier.Verify(c, rawIDToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Token": "Invalid Token",
			})
			c.Abort()
			return
		}

		// Extract custom claims
		var claims struct {
			Name  string `json:"name"`
			Email string `json:"preferred_username"`
		}
		if err := idToken.Claims(&claims); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"Claims": "Error extracting custom claims",
			})

		}

		c.Set("userEmail", claims.Email)
		c.Set("userName", claims.Name)

		c.Next()
	}
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// We can obtain the session token from the requests cookies, which come with every request
		rawIDToken := strings.Trim(strings.TrimLeft(c.GetHeader("authorization"), "Bearer"), " ")

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.Parse(rawIDToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(setting.AppSetting.JwtSecret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()

	}

}
