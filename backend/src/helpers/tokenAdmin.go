package helper

import (
	"codequest/src/configs"

	"codequest/src/models"
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_ADMIN_KEY = configs.JWTSecretKey()

func GenerateAllAdminTokens(email string, name string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &models.SignedAdminDetails{
		Email: email,
		Name: name,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &models.SignedAdminDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_ADMIN_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_ADMIN_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateAdminToken(signedToken string) (claims *models.SignedAdminDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.SignedAdminDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_ADMIN_KEY), nil
		},
	)

	if err != nil {
		msg = fmt.Sprintf("error parsing token: %s", err)
		return
	}

	claims, ok := token.Claims.(*models.SignedAdminDetails)
	if !ok {
		msg = "invalid token claims"
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		msg = "token has expired"
		return
	}

	return claims, msg
}

func HashAdminPassword(password string, cost int) ([]byte, error) {
	if len(password) == 0 {
		return nil, errors.New("password cannot be empty")
	}

	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		return nil, errors.New("invalid cost parameter")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

func VerifyAdminPassword(userPassword string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return errors.New("passwords do not match")
		default:
			return err // Return the original error for other cases
		}
	}

	return nil // Passwords match
}