package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

func ValidateUser(rawJson map[string]string) string {
	if rawJson["name"] == "" {
		return "Name is not an empty field"
	}

	if rawJson["email"] == "" {
		return "Email is not an empty field"
	}

	if rawJson["password"] == "" {
		return "Password is not an empty field"
	}

	return "ok"
}
