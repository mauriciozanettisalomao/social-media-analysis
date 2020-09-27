package twitter

import (
	"fmt"
	"net/url"
	"os"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"

	"github.com/ChimeraCoder/anaconda"
)

const (
	// Name of social media
	Name = "twitter"
)

var (
	consumerKey       = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

type twitter struct{}

func (t twitter) Search(param string, filter map[string]string) ([]app.Mention, error) {

	var (
		searchResult anaconda.SearchResponse
		err          error
	)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	v := url.Values{}

	for key, value := range filter {
		v.Set(key, value)
	}

	for i := 0; i < 5; i++ {

		searchResult, err = api.GetSearch(param, v)
		if err != nil {
			return nil, fmt.Errorf("could not get twitter mentions: %v", err)
		}

		if len(searchResult.Statuses) > 0 {
			break
		}

	}

	mentions := make([]app.Mention, 0)
	for _, tweet := range searchResult.Statuses {

		if tweet.Text == "" {
			continue
		}

		mention := app.Mention{
			Name:      Name,
			Text:      tweet.Text,
			ID:        tweet.User.ScreenName,
			User:      tweet.User.Name,
			Lang:      tweet.Lang,
			CreatedAt: tweet.CreatedAt,
		}

		mentions = append(mentions, mention)

	}

	return mentions, nil
}

// New returns a new instance of a twitter social media
func New() app.SocialMedia {
	return twitter{}
}
