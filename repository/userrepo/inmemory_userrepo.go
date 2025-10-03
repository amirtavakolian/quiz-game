package userrepo

import "github.com/amirtavakolian/quiz-game/entity"

type InMemoryUserRepo struct {
	Users map[string]entity.User
}

func NewInMemoryUserRepo() InMemoryUserRepo {
	return InMemoryUserRepo{
		Users: make(map[string]entity.User),
	}
}

func (r InMemoryUserRepo) FindByPhoneNumber(phoneNumber string) (bool, error) {
	for _, user := range r.Users {
		if user.PhoneNumber == phoneNumber {
			return true, nil
		}
	}
	return false, nil
}
