package endpoints

import "github.com/gin-gonic/gin"

func Login(c *gin.Context)                {}
func Register(c *gin.Context)             {}
func SendConfirmationMail(c *gin.Context) {}
func Logout(c *gin.Context)               {}

func InitializeAuthEndpoints(engine *gin.Engine) {

}
