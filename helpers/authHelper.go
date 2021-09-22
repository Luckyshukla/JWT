package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)
func MatchUserTypeToUid(c *gin.Context, user_id string)(err error){
	userType :=c.GetString("user_type")
	uid:= c.GetString("uid")
	err = nil
	if userType == "USER" && uid!= userID{
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}