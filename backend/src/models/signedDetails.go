package models

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email     string
	Username  string
	FirstName string
	LastName  string
	Uid       string
	ClassRoom string
	jwt.StandardClaims
}

type SignedAdminDetails struct {
	Email string
	Name  string
	Uid   string
	jwt.StandardClaims
}
