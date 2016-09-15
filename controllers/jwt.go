package controllers
 
 import (
 	"time"
 	"fmt"
 
 	"github.com/astaxie/beego"
 	"github.com/astaxie/beego/logs"
 	jwt "github.com/dgrijalva/jwt-go"
 
 	"github.com/bizoru/nodo-central-oas/configs"
 	"github.com/bizoru/nodo-central-oas/models"
 )
 
 // oprations for Jwt
 type JwtController struct {
 	beego.Controller
 }
 
 func (this *JwtController) URLMapping() {
 	this.Mapping("IssueToken", this.IssueToken)
 }
 
 // @Title IssueToken
 // @Description Issue a Json Web Token
 // @Success 200 string
 // @Failure 403 no privilege to access
 // @Failure 500 server inner error
 // @router /issue-token [get]
 func (this *JwtController) IssueToken() {
 	log := logs.NewLogger(10000)
 	log.SetLogger("console", "")
 
 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	    "sesionData": models.GetNewSession("jjoc007") ,
	    "exp": time.Now().Unix() + 36000,
	})

 	// The claims object allows you to store information in the actual token.
 	tokenString, err := token.SignedString(configs.PrivateKey)

 	// tokenString Contains the actual token you should share with your client.
 	this.Data["json"] = map[string]string{"token": tokenString}
 	
 	fmt.Println(tokenString, err)

 	//	log.Debug(err.Error())
 	this.ServeJSON()

 	
 }