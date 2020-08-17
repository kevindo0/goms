package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Claims can have user id.. etc for Identification purpose
type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

var (
	privateKey []byte
	publicKey  []byte
	err        error
)

const (
	longForm = "Jan 2, 2006 at 3:04pm (MST)"
)

func errLog(err error) {
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}

func init() {
	privateKey, err = ioutil.ReadFile("private.pem")
	errLog(err)
	publicKey, err = ioutil.ReadFile("public.pem")
	errLog(err)
}

func jwtTokenGen() (interface{}, error) {
	privateRSA, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}
	claims := AppClaims{
		"RAJINIS*",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(privateRSA)
	return ss, err
}

func jwtTokenRead(inToken interface{}) (interface{}, error) {
	publicRSA, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(inToken.(string), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicRSA, err
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func getTokenRemainingValidity(timestamp interface{}) int {
	expireOffset := 0
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := tm.Sub(time.Now())
		if remainder > 0 {
			fmt.Println(remainder)
			return int(remainder.Seconds()) + expireOffset
		}
	}
	return expireOffset
}

func main() {
	signedString, err := jwtTokenGen()
	fmt.Println(signedString, err)
	claims, err := jwtTokenRead(signedString)
	if err != nil {
		errLog(err)
	}
	fmt.Println("claims:", claims)
	claimValue := claims.(jwt.MapClaims)
	fmt.Println("claimValue:", claimValue)
	fmt.Println(claimValue["iss"], claimValue["exp"], claimValue["userId"])
	//  t, _ := time.Parse(longForm, string(claimValue["exp"].(float64)))
	fmt.Println(getTokenRemainingValidity(claimValue["exp"]))
}
