package usecase

import (
	"context"
	"errors"

	"github.com/kou12345/go-restapi-example/api/internal/domain/administrator"
	"github.com/kou12345/go-restapi-example/api/internal/domain/user"
)

type SaveUserUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewSaveUserUsecase(uds user.UserDomainService, ads administrator.AdminDomainService) *SaveUserUsecase {
	return &SaveUserUsecase{uds: uds, ads: ads}
}

func (us *SaveUserUsecase) Run(ctx context.Context, param *user.User) error {
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		return ErrInvalidUserID
	}

	adminUser, err := us.ads.FindUser(ctx, value)
	if err != nil {
		return err
	}

	if adminUser.GetAdmin() == 1 || param.GetUUID() == value {
		if err := us.uds.EditUser(ctx, param); err != nil {
			return err
		}
		return nil
	}

	if adminUser.GetAdmin() != 1 {
		return ErrInvalidAdmin
	}

	return errors.New("unexpected error occurred")
}
