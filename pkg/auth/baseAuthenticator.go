package auth

func (adapter *baseAuthenticator) Login(user User) (bool, error) {

	return false, nil
}

func (adapter *baseAuthenticator) CreateUser(user User) (bool, error) {
	return false, nil
}

func (adapter *baseAuthenticator) UpdatePassword(user User, newPassword string) (bool, User, error) {
	return false, User{}, nil
}
