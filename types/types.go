package types

import jwt "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password,omitempty"`
}

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
