package postgres

import (
	"context"
	"message-board/domain/model/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db DB
}

func NewUserRepository(db DB) user.IRepository {
	return &userRepository{db: db}
}

type userDTO struct {
	gorm.Model
	ID    uint64 `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

func newUserDTO(u *user.User) userDTO {
	return userDTO{
		ID:    u.ID.Uint64(),
		Name:  u.Name.String(),
		Email: u.Email.String(),
	}
}

func (u userDTO) toDomain() *user.User {
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

func (r *userRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	udto := newUserDTO(u)
	db := r.db.Conn(ctx)
	err := db.Create(udto).Error
	if err != nil {
		return nil, err
	}
	return udto.toDomain(), nil
}
