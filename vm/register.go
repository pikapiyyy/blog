package vm

type RegisterViewModel struct {
	BaseViewModel
}

type RegisterVM struct{}

func (RegisterVM) GetVM() RegisterViewModel {
	v := RegisterViewModel{BaseViewModel{Title: "Register"}}
	return v
}
