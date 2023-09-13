package middleware

import(
	"fmt"
	"net/http"
	helper "github.com/a-viraj/golang-auth/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken:=c.Request.Header.Get("token")
		if clientToken==""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":fmt.Sprintf("Auth error provided")})
			c.Abort()
			return
		}
		claims,err:=helper.ValidateToken(clientToken)
		if err!=""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			c.Abort()
			return
		}
		c.Set("email",claims.Email)
		c.Set("firstname",claims.FirstName)
		c.Set("lastname",claims.LastName)
		c.Set("userid",claims.UserId)
		c.Set("usertype",claims.UserType)
		c.Next()
	}

}