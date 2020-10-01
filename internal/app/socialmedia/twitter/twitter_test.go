package twitter

import (
	"errors"
	"testing"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	//Search(param string, filter map[string]string) ([]app.Mention, error) {

	cases := []struct {
		name     string
		param    string
		expected []app.Mention
		err      error
	}{
		{
			name:     "no credentials",
			param:    "teste",
			expected: []app.Mention{},
			err:      errors.New("unauthorized"),
		},
	}

	assertion := assert.New(t)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			twitter := twitter{}
			result, err := twitter.Search(tc.param, nil)
			if err != nil {
				assertion.NotNil(tc.err)
				return
			}

			assertion.Equal(tc.expected, result)
		})
	}

}
