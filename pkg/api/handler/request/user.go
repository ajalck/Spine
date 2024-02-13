package request

type User struct{
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}