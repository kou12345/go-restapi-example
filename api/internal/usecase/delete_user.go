package usecase

import (
	"context"
	"errors"

	"github.com/kou12345/go-restapi-example/api/internal/domain/administrator"
	"github.com/kou12345/go-restapi-example/api/internal/domain/user"
)

type DeleteUserUsecase struct {
	uds user.UserDomainService
	ads administrator.AdminDomainService
}

func NewDeleteUserUsecase(uds user.UserDomainService, ads administrator.AdminDomainService) *DeleteUserUsecase {
	return &DeleteUserUsecase{uds: uds, ads: ads}
}

// ユーザー削除
// 引数のidは削除対象のユーザーID
// ctxから取得したuser_idはリクエストをしているユーザーID
func (us *DeleteUserUsecase) Run(ctx context.Context, id string) error {
	// ユーザーID取得
	value, ok := ctx.Value("user_id").(string)
	if !ok {
		// ユーザーIDが取得できない場合はエラー
		return ErrInvalidUserID
	}

	// 管理者ユーザー取得
	adminUser, err := us.ads.FindUser(ctx, value)
	if err != nil {
		return err
	}

	// 管理者ユーザーまたは、削除対象ユーザーのidが自身と一致する場合
	if adminUser.GetAdmin() == 1 || id == value {
		// ユーザー削除
		if err := us.uds.DeleteUser(ctx, id); err != nil {
			return err
		}
		return nil
	}

	// ここに到達する場合は、
	// 管理者ユーザーではない 且つ 削除対象ユーザーのidが自身と一致しない場合

	if adminUser.GetAdmin() != 1 {
		return ErrInvalidAdmin
	}

	// ここに到達する場合はない気がする
	// 到達することが想定されないが、全ての実行パスで値を返す必要があるためのエラー
	return errors.New("unexpected error occurred")

}
