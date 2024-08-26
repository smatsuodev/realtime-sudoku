//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"github.com/samber/mo"
	"sudoku/model"
)

type IUserRepository interface {
	FindByGitHubUserID(githubUserID uint) (mo.Option[*model.User], error)
	// Save 新規レコード作成の場合, 採番された ID をセットする
	Save(user *model.User) error
}
