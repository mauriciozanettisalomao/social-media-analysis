package analysis

import (
	"testing"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"

	"github.com/stretchr/testify/assert"
)

type socialMediaMock struct {
	searchMock func(param string, filter map[string]string) ([]app.Mention, error)
}

func (s socialMediaMock) Search(param string, filter map[string]string) ([]app.Mention, error) {
	return s.searchMock(param, filter)
}

func TestSearch(t *testing.T) {

	type filters struct {
		key   string
		value string
	}

	case1 := make(map[string][]app.Mention)
	case1["twitter"] = []app.Mention{
		{
			ID:        "NFL_DovKleiman",
			Name:      "twitter",
			Text:      "That team was just amazing. Let's not forget Nick Foles took over without much playing time with the starters and end up winning all the playoffs games and Super Bowl MVP",
			User:      "Dov Kleiman",
			Lang:      "en",
			CreatedAt: "Thu Oct 01 20:23:33 +0000 2020",
		},
	}

	cases := []struct {
		name             string
		socialMediaParam string
		param            string
		targetLang       string
		filters          filters
		mentions         []app.Mention
		expected         app.Response
		err              error
	}{
		{
			name:             "search completely with success",
			socialMediaParam: "twitter",
			param:            "nfl",
			targetLang:       "en",
			mentions: []app.Mention{
				{
					ID:        "NFL_DovKleiman",
					Name:      "twitter",
					Text:      "That team was just amazing. Let's not forget Nick Foles took over without much playing time with the starters and end up winning all the playoffs games and Super Bowl MVP",
					User:      "Dov Kleiman",
					Lang:      "en",
					CreatedAt: "Thu Oct 01 20:23:33 +0000 2020",
				},
			},
			expected: app.Response{
				Parameter:   "nfl",
				Filters:     []app.Filter{},
				SocialMedia: case1,
			},
			err: nil,
		},
	}

	assertion := assert.New(t)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			s := socialMediaMock{}
			s.searchMock = func(param string, filter map[string]string) ([]app.Mention, error) {
				return tc.mentions, nil
			}

			result, err := Search(tc.socialMediaParam, tc.param, tc.targetLang, nil, s)
			if err != nil {
				assertion.NotNil(tc.err)
				return
			}

			assertion.Equal(tc.expected, result)
		})
	}

}
