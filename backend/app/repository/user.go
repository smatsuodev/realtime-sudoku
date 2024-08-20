//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"github.com/samber/mo"
	"sudoku/model"
)

type IUserRepository interface {
	FindByGitHubUserID(githubUserID uint) (mo.Option[*model.User], error)
	Save(user *model.User) error
}
