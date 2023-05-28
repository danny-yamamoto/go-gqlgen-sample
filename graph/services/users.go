package services

import (
	"context"
	"fmt"

	"github.com/danny-personal/go-gqlgen-sample/graph/db"
	"github.com/danny-personal/go-gqlgen-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	// 1. SQLBoilerで生成されたORMコードを呼び出す
	user, err := db.Users( // from users
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name), // select id, name
		db.UserWhere.Name.EQ(name),                                  // where name = {引数nameの内容}
	).One(ctx, u.exec) // limit 1
	// 2. エラー処理
	if err != nil {
		fmt.Println("koko")
		return nil, err
	}
	// 3. 戻り値の*model.User型を作る
	return convertUser(user), nil
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
