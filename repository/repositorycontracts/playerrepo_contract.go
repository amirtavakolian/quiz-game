package repositorycontracts

type PlayerRepoContract interface {
	Store(phoneNumber string) error
}
