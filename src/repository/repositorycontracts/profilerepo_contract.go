package repositorycontracts
import "github.com/amirtavakolian/quiz-game/param/profileparams"

type ProfileRepoContract interface {
	Update(profileRepo profileparams.UpdateProfile) error
}
