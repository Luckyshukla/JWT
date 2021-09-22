package Controller

import (
	"context"
	"net/http"
	"os/user"
	"time"

	"github.com/aquasecurity/libbpfgo/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/database"
	"main.go/models"
)

var userCollection mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUser()

func GetUsers() gin.HandlerFunc{
	return func (c *gin.Context)  {
		userID:=c.param("user_id")
		if err:=helpers.MatchUserTypeToUid(c, userID);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error})
		}
		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		err:=userCollection.FindOne(ctx, bson.H{"user_id":userID}).Decode(&user)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON((http.StatusOK, user)
	}

}
