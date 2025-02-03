package models

type userLoginFormData interface {
	GetEmail() string
	GetPassword() string
}

func (m *UserModel) Authenticate(f userLoginFormData) (int, error) {
	return 0, nil
}
