package repository

import (
	"context"
	"time"

	"geektime-basic-go/week04/webook/internal/domain"
	"geektime-basic-go/week04/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail    = dao.ErrUserDuplicateEmail
	ErrUserNotFound          = dao.ErrDataNotFound
	ErrDataTooLong           = dao.ErrDataTooLong
	ErrUserDuplicateNickname = dao.ErrUserDuplicateNickname
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(d *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: d}
}

func (ur *UserRepository) Create(ctx context.Context, u domain.User) error {
	return ur.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.dao.FindByEmail(ctx, email)
	return domain.User{Id: u.Id, Email: u.Email, Password: u.Password}, err
}

func (ur *UserRepository) Update(ctx context.Context, user domain.User) error {
	return ur.dao.Update(ctx, dao.User{
		Id:           user.Id,
		Nickname:     user.Nickname,
		Birthday:     user.Birthday,
		Introduction: user.Introduction,
	})
}

func (ur *UserRepository) FindByID(ctx context.Context, id int64) (domain.User, error) {
	u, err := ur.dao.FindByID(ctx, id)
	return domain.User{
		Id:           u.Id,
		Email:        u.Email,
		Nickname:     u.Nickname,
		Birthday:     u.Birthday,
		Introduction: u.Introduction,
		CreateAt:     time.UnixMilli(u.CreateAt),
	}, err
}