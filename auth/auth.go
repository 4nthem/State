package auth

import (
	"crypto/rsa"
	"github.com/4nthem/State/models"

	log "github.com/cihub/seelog"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func InitTokens() (err error) {
	signBytes, err := ioutil.ReadFile("config/rsaKey")
	if err != nil {
		log.Error("Error reading private key from file: ", err)
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Error("Error parsing private key from file: ", err)
		return
	}

	verifyBytes, err := ioutil.ReadFile("config/pubKey")
	if err != nil {
		log.Error("Error reading public key from file: ", err)
		return
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Error("Error parsing public key from file: ", err)
		return
	}
	return
}

func AuthenticateToken() (err error) {
	return
}

func NewToken(accessLevel string, theUser models.User) (tokenString string, err error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims["AccessToken"] = accessLevel
	t.Claims["UserInfo"] = theUser

	tokenString, err = t.SignedString(signKey)
	if err != nil {
		log.Error("Error creating token for user: " + theUser.Id)
	}
	return
}
