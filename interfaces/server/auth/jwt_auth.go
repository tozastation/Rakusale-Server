package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// CreateJWT is
func CreateJWT(name string, email string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	// [Set] Claim
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "tozastation"
	claims["sub"] = "AccessToken"
	claims["https://idp.example.com/claim-types/email"] = email
	claims["https://idp.example.com/claim-types/user-name"] = name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["company"] = "Rakusale, Inc"
	// [SIGN]
	signedString, _ := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
	fmt.Println("Credential: " + name)
	fmt.Println(string([]byte(signedString)))
	return string([]byte(signedString))
}
