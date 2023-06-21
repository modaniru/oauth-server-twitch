package model

type User struct {
	Id               int
	Username         string
	TwitchId         int
	ImageLink        string
	RegistrationDate string
}

type UserInfo struct {
	Username         string
	TwitchId         int
	ImageLink        string
	RegistrationTime string
}

type AccessToken struct {
	AccessToken string `json:"access_token" binding:"required"`
}
