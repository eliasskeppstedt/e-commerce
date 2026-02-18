package customer

type user struct {
	UserID   int    `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"emailAddress"`
	// FirstName    string `json:"firstName"`
	// LastName     string `json:"lastName"`
	// Address      string `json:"address"`
	// ZipCode      string `json:"zipCode"`
	// PhoneNumber  string `json:"phoneNumber"`
}
