package configs
 
 import (
	"io/ioutil"
	"confignodoapi/models"

	jwt "github.com/dgrijalva/jwt-go"

	"crypto/sha256"
	"io"
	"fmt"

)
 
 var (
 	PrivateKey []byte
 )
 
 func Init() {
 	PrivateKey, _ = ioutil.ReadFile("conf/privateKey.txt")
 }

func ValidateToken(tokenString string) (models.MessageReturn, interface{}){

	tokenVal, errVal := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	    return PrivateKey, nil
	})

	if claims, ok := tokenVal.Claims.(jwt.MapClaims); ok && tokenVal.Valid {
	    
		return models.GetNewMessage(0, "Token Valido"), claims["sesionData"]

	} else if ve, ok := errVal.(*jwt.ValidationError); ok {
	    if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	        return models.GetNewMessage(1, "Token Invalido"), nil
	    } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	        return models.GetNewMessage(1, "Token Vencido"), models.SessionData{}
	    } else {
	        return models.GetNewMessage(1, "Error" + ve.Error()), models.SessionData{}
	    }
	} else {
	    return models.GetNewMessage(1, "Error" + ve.Error()), models.SessionData{}
	}

}

func GenerateKeysHash(password string) (string){

	h := sha256.New()
	io.WriteString(h,password)
	fmt.Println(fmt.Sprint(h.Sum(nil)))

	return fmt.Sprint(h.Sum(nil))

}
