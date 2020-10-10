package twitter

import (
	"fmt"
	"net/url"
	"os"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/sirupsen/logrus"
)

const (
	// Name of social media
	Name = "twitter"

	// MetricFollowers name of the metric
	MetricFollowers = "Followers"
)

var (
	consumerKey       = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	api *anaconda.TwitterApi

	userFallback = app.User{
		ID:                 599689553,
		Name:               "DXC Technology",
		CreatedAt:          "Tue Jun 05 00:40:02 +0000 2012",
		Description:        "We help our customers across the entire enterprise tech stack with differentiated industry solutions. Connect with our award-winning teams across the world.",
		FollowersCount:     32742,
		Location:           "Global Technology Company",
		ScreenName:         "DXCTechnology",
		StatusesCount:      21647,
		MentionsCount:      31818,
		ProfileVisitsCount: 88412,
		URL:                "https://t.co/rxQgOzkYut",
	}

	metricFallback = app.Metric{
		Name:        MetricFollowers,
		Description: "percentage of followers by country",
	}
)

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}

type twitter struct{}

func (t twitter) Search(param string, filter map[string]string) ([]app.Mention, error) {

	var (
		searchResult anaconda.SearchResponse
		err          error
	)

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

	t.SearchUser(param, nil)

	return mentions, nil
}

func (t twitter) SearchUser(username string, filter map[string]string) (app.User, error) {

	var (
		searchUserResult anaconda.User
		err              error
	)

	v := url.Values{}

	for key, value := range filter {
		v.Set(key, value)
	}

	searchUserResult, err = api.GetUsersShow(username, v)
	if err != nil {
		log.WithFields(log.Fields{
			"username": username,
			"v":        username,
		}).Errorf("error on search user: %v", err)
		userFallback.Error = err.Error()
		return userFallback, nil
	}

	user := app.User{
		ID:                 searchUserResult.Id,
		Name:               searchUserResult.Name,
		ScreenName:         searchUserResult.ScreenName,
		CreatedAt:          searchUserResult.CreatedAt,
		Description:        searchUserResult.Description,
		Email:              searchUserResult.Email,
		FollowersCount:     searchUserResult.FollowersCount,
		StatusesCount:      searchUserResult.StatusesCount,
		MentionsCount:      int(searchUserResult.StatusesCount) * 4,
		ProfileVisitsCount: searchUserResult.FollowersCount * 3,
		Location:           searchUserResult.Location,
		URL:                searchUserResult.URL,
	}

	return user, nil
}

func (t twitter) SearchFollowers(username string, filter map[string]string) (app.Metrics, error) {

	metric := make(map[string]float32)
	metric["India"] = 19.9
	metric["Japan"] = 17.8
	metric["SouthKorea"] = 26.4
	metric["Brazil"] = 13.5
	metric["USA"] = 32.5
	metric["UK"] = 25.6
	metricFallback.Values = metric

	metrics := app.Metrics{
		Metrics: []app.Metric{
			metricFallback,
		},
	}

	return metrics, nil
}

// New returns a new instance of a twitter social media
func New() app.SocialMedia {
	return twitter{}
}
