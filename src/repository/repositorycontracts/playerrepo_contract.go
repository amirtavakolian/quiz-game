package repositorycontracts

type PlayerRepoContract interface {
	Store(phoneNumber string) (int64, error)
}
