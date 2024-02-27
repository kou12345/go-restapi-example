package usecase

import (
	"context"
	"errors"

	"github.com/kou12345/go-restapi-example/api/internal/domain/administrator"
	"github.com/kou12345/go-restapi-example/api/internal/domain/user"
)

type FindUserByIdUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewFindUserByIdUsecase(uds user.UserDomainService, ads administrator.AdminDomainService) *FindUserByIdUsecase {
	return &FindUserByIdUsecase{uds: uds, ads: ads}
}

// ユーザーIDでユーザーを検索
func (us *FindUserByIdUsecase) Run(ctx context.Context, id string) (*user.User, error) {
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, ErrInvalidUserID
	}

	adminUser, err := us.ads.FindUser(ctx, value)
	if err != nil {
		return nil, err
	}

	if adminUser.GetAdmin() == 1 || id == value {
		user, err := us.uds.FindUserById(ctx, id)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	if adminUser.GetAdmin() != 1 {
		return nil, ErrInvalidAdmin
	}

	return nil, errors.New("unexpected error occurred")
}
