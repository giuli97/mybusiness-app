package models

import jwt "github.com/dgrijalva/jwt-go"

type Session struct {
	Id         int64  `json:"id"`
	Token      string `json:"token"`
	ExpiresIn  int64  `json:"expiresIn"`
	LastUpdate string `json:"lastUpdate"`
	IdUser     int64  `json:"idUser"`
}

type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}

type ResponseToken struct {
	Token string `json:"token"`
}

type JWTResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh_token"`
}
