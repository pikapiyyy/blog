package vm

import (
	"blog/model"
)

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexVM struct{}

func (IndexVM) GetVM(username string) IndexViewModel {
	u1, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(u1.ID)

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	v.SetCurrentUser(username)

	return v
}
