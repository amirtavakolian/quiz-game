package userrepo

type UserRepositoryContract interface {
	FindByPhoneNumber(phoneNumber string) (bool, error)
}
