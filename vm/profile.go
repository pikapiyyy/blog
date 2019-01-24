// 个人主页数据
package vm

import (
	"blog/model"
)

type ProfileViewModel struct {
	BaseViewModel
	ProfileUser model.User
	Posts       []model.Post
}

type ProfileVM struct{}

func (ProfileVM) GetVM(pUsername, sUsername string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUsername)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostsByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Posts = *posts
	v.SetCurrentUser(sUsername)

	return v, nil
}
