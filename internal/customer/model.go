package customer

type user struct {
	UserID      int    `json:"user_id"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	PhoneNumber string `json:"phone_number"`
	IsAdmin     string `json:"is_admin"`
}
