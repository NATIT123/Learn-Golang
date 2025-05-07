package biz

import (
	"context"
	"main/common"
	models "main/modules/user/models/postgreSQL"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*models.User, error)
	CreateUser(ctx context.Context, data *models.UserCreate) error
}

type Hasher interface {
	Hash(data string) (string, error)
	Compare(hashedPassword string, password string) error
}

// user a, pass 123456: pass in db abcdef
type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *models.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		if *user.Status == models.UserStatusDeleted {
			return common.ErrEntityDeleted(models.EntityName, models.ErrUserDeleted)

		}
		return models.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	hashedPassword, err := business.hasher.Hash(data.Password + salt)
	if err != nil {
		return common.ErrInternal(err)
	}

	data.Password = hashedPassword
	data.Salt = salt
	data.Role = "user" // hard code

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.EntityName, err)
	}

	return nil
}
