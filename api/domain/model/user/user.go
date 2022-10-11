package user

type User struct {
	ID    ID
	Name  Name
	Email Email
}

type Users []*User

func NewUser(id ID, name Name, email Email) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func NewUserForCreate(name Name, email Email) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}
