package analysis

import (
	"fmt"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"
)

// Search gets the social media mentions
func Search(socialMediaParam, param, targetLang string, filters map[string]string, socialMedia app.SocialMedia) (app.ResponseMentions, error) {

	mentions, err := socialMedia.Search(param, filters)
	if err != nil {
		return app.ResponseMentions{}, fmt.Errorf("error retrieving mentions: %v", err)
	}

	filtersResp := make([]app.Filter, 0)
	for k, v := range filters {
		filtersResp = append(filtersResp, app.Filter{
			Key:   k,
			Value: v,
		})
	}

	socialMediaResp := make(map[string][]app.Mention)
	socialMediaResp[socialMediaParam] = mentions

	if targetLang != "" {
		// TODO implement translation
	}

	return app.ResponseMentions{
		Parameter:   param,
		Filters:     filtersResp,
		SocialMedia: socialMediaResp,
	}, nil

}

// SearchUser gets user data
func SearchUser(username string, filters map[string]string, socialMedia app.SocialMedia) (app.ResponseUser, error) {

	user, err := socialMedia.SearchUser(username, filters)
	if err != nil {
		return app.ResponseUser{}, fmt.Errorf("error retrieving user data: %v", err)
	}

	filtersResp := make([]app.Filter, 0)
	for k, v := range filters {
		filtersResp = append(filtersResp, app.Filter{
			Key:   k,
			Value: v,
		})
	}

	return app.ResponseUser{
		User: user,
	}, nil

}

// SearchFollowers gets folloers of a user
func SearchFollowers(username string, filters map[string]string, socialMedia app.SocialMedia) (app.Metrics, error) {

	metrics, err := socialMedia.SearchFollowers(username, filters)
	if err != nil {
		return app.Metrics{}, fmt.Errorf("error retrieving user data: %v", err)
	}
	return metrics, nil

}
