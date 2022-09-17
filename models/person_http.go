package models

type PersonResultHTTP struct {
	Status PersonStatus   `json:"status"`
	Result []PersonResult `json:"result"`
}

type PersonStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type PersonResult struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	UUID      string `json:"uuid"`
}
