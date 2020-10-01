package analysis

import (
	"fmt"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"
)

// Search gets the social media mentions
func Search(socialMediaParam, param, targetLang string, filters map[string]string, socialMedia app.SocialMedia) (app.Response, error) {

	mentions, err := socialMedia.Search(param, filters)
	if err != nil {
		return app.Response{}, fmt.Errorf("error retrieving mentions: %v", err)
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

	return app.Response{
		Parameter:   param,
		Filters:     filtersResp,
		SocialMedia: socialMediaResp,
	}, nil

}
