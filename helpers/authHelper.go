package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("UserType")
	uid := c.GetString("User_id")
	if userType == "USER" && uid != userId {
		err := errors.New("unautherized access")
		return err
	}
	err:=CheckUserType(c,userType)
	return err
}
func CheckUserType(c *gin.Context,t string) error{
	usertype:=c.GetString("usertype")
	if usertype!=t{
		return errors.New("unauth access")
	}
	return nil
}