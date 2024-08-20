package gorm

import (
	"errors"
	"github.com/samber/mo"
	"gorm.io/gorm"
	"sudoku/model"
	"sudoku/repository"
)

var _ repository.IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByGitHubUserID(githubUserID uint) (mo.Option[*model.User], error) {
	var user User
	if err := r.db.Where("github_user_id = ?", githubUserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mo.None[*model.User](), nil
		}
		return mo.None[*model.User](), err
	}

	return mo.Some(model.NewUser(user.ID, user.GitHubUserID, user.Name)), nil
}

func (r *UserRepository) Save(user *model.User) error {
	if !user.IsIDPresent() {
		// 新規作成
		return r.db.Create(&User{
			Name:         user.Name(),
			GitHubUserID: user.GithubID(),
		}).Error
	}

	// 既存ユーザーの更新
	return r.db.Model(&User{}).Where("id = ?", user.ID()).Updates(&User{
		Name:         user.Name(),
		GitHubUserID: user.GithubID(),
	}).Error
}
