package authResponse

type ResultGetProfile struct {
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
