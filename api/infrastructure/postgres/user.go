package postgres

import (
	"context"
	"message-board/domain/model/user"
)

type userRepository struct {
	db DB
}

func NewUserRepository(db DB) user.IRepository {
	return &userRepository{db: db}
}

type UserDTO struct {
	ID    uint64 `gorm:"primaryKey,autoIncrement,column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

type userDTOs []UserDTO

func newUserDTO(u *user.User) *UserDTO {
	return &UserDTO{
		Name:  u.Name.String(),
		Email: u.Email.String(),
	}
}

func (u *UserDTO) toUser() *user.User {
	email, err := user.NewEmail(u.Email)
	if err != nil {
		return nil
	}
	return &user.User{
		ID:    user.NewID(u.ID),
		Name:  user.NewName(u.Name),
		Email: email,
	}
}

func (u userDTOs) toUsers() user.Users {
	users := make(user.Users, len(u))
	for i, dto := range u {
		users[i] = dto.toUser()
	}
	return users
}

func (r *userRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	udto := newUserDTO(u)
	db := r.db.Conn(ctx)
	err := db.Table("users").Create(udto).Error
	if err != nil {
		return nil, err
	}
	return udto.toUser(), nil
}

func (r *userRepository) Find(ctx context.Context) (user.Users, error) {
	var udtos userDTOs
	db := r.db.Conn(ctx)
	err := db.Table("users").Find(&udtos).Error
	if err != nil {
		return nil, err
	}
	return udtos.toUsers(), nil
}
