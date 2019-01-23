package vm

type LoginViewModel struct {
	BaseViewModel
}

type LoginVM struct{}

func (LoginVM) GetVM() LoginViewModel {
	v := LoginViewModel{BaseViewModel{Title: "Login"}}

	return v
}
