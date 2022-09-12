package user

type User struct {
	ID    ID
	Name  Name
	Email Email
}

func NewUser(id ID, name Name, email Email) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
