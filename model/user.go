package model

//User model
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// FullName will get the user's full name
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}
