package usecase

import (
	"context"

	"github.com/kou12345/go-restapi-example/api/internal/domain/administrator"
	"github.com/kou12345/go-restapi-example/api/internal/domain/user"
)

type FindUsersUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewFindUsersUsecase(uds user.UserDomainService, ads administrator.AdminDomainService) *FindUsersUsecase {
	return &FindUsersUsecase{uds: uds, ads: ads}
}

func (us *FindUsersUsecase) Run(ctx context.Context) ([]*user.User, error) {
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, ErrInvalidUserID
	}

	adminUser, err := us.ads.FindUser(ctx, value)
	if err != nil {
		return nil, err
	}

	if adminUser.GetAdmin() != 1 {
		return nil, ErrInvalidAdmin
	}

	users, err := us.uds.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
