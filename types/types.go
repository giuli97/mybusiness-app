package types

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Session struct {
	Id         int64  `json:"id"`
	Token      string `json:"token"`
	ExpiresIn  int64  `json:"expiresIn"`
	LastUpdate string `json:"lastUpdate"`
	IdUser     int64  `json:"idUser"`
}
