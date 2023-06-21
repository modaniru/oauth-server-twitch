package utils

type ValidResponse struct{
	ClientId string `json:"client_id"`
	Login string `json:"login"`
	Scopes []string `json:"scopes"`
	UserId string `json:"user_id"`
	ExpiresIn int `json:"expires_in"`
}

type UsersInfo struct{
	Data []UserInfo `json:"data"`
}

type UserInfo struct{
	Id string `json:"id"`
	Login string `json:"login"`
	DisplayName string `json:"display_name"`
	Type string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	OfflineImageUrl string`json:"offline_image_url"`
	ViewCount int`json:"view_count"`
	Email string`json:"email"`
	CreatedAt string`json:"created_at"`
}