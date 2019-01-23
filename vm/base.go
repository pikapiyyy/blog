package vm

type BaseViewModel struct {
	Title       string
	Errs        []string
	CurrentUser string
}

func (this *BaseViewModel) SetTitle(title string) {
	this.Title = title
}

func (this *BaseViewModel) AddError(err ...string) {
	this.Errs = append(this.Errs, err...)
}

func (this *BaseViewModel) SetCurrentUser(username string) {
	this.CurrentUser = username
}
