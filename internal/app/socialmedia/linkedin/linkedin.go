package linkedin

import (
	"fmt"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"
)

const (
	// Name of social media
	Name = "linkedin"
)

var (
	consumerKey       = "" // os.Getenv("...")
	consumerSecret    = "" // os.Getenv("...")
	accessToken       = "" // os.Getenv("...")
	accessTokenSecret = "" // os.Getenv("...")
)

type linkedin struct{}

func (t linkedin) Search(param string, filter map[string]string) ([]app.Mention, error) {

	return nil, fmt.Errorf("social media %s not implemented yet", Name)
}

// New returns a new instance of a twitter social media
func New() app.SocialMedia {
	return linkedin{}
}
