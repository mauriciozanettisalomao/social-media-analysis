package app

// SocialMedia is an interface to handle social medias as twitter, facebook, instagram, etc
type SocialMedia interface {
	Search(param string, filter map[string]string) ([]Mention, error)
	SearchUser(username string, filter map[string]string) (User, error)
	SearchFollowers(username string, filter map[string]string) (Metrics, error)
}

// Translator is an interface to handle the transaltion when requested
type Translator interface {
	Translate(text string) (Transalation, error)
}

// User is used to return details of the user
type User struct {
	ID                 int64  `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	Description        string `json:"description,omitempty"`
	Email              string `json:"email,omitempty"`
	FollowersCount     int    `json:"followersCount,omitempty"` // The number of followers this account currently has. Under certain conditions of duress, this field will temporarily indicate
	Location           string `json:"location,omitempty"`       // User defined location
	ScreenName         string `json:"screenName,omitempty"`
	StatusesCount      int64  `json:"statusesCount,omitempty"` // The number of Tweets (including retweets) issued by the user
	MentionsCount      int    `json:"mentionsCount,omitempty"`
	ProfileVisitsCount int    `json:"profileVisitsCount,omitempty"`
	URL                string `json:"url,omitempty"`
	Error              string `json:"error,omitempty"`
}

// Metrics is used to return user's metrics
type Metrics struct {
	Metrics []Metric `json:"metrics,omitempty"`
}

// Metric is used to return user's metric
type Metric struct {
	Name        string             `json:"name,omitempty"`
	Values      map[string]float32 `json:"values,omitempty"`
	Description string             `json:"description,omitempty"`
}

// Mention is used to return details of social media mentions
type Mention struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	User      string `json:"user"`
	Lang      string `json:"lang"`
	Sentiment string `json:"sentiment,omitempty"`
	CreatedAt string `json:"createdAt"`
}

// Filter is used to return all the filters used on the search
type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ResponseMentions is used to return all the informations of social media mentions
type ResponseMentions struct {
	Parameter   string               `json:"parameter"`
	Filters     []Filter             `json:"filters"`
	SocialMedia map[string][]Mention `json:"socialMedia"`
}

// ResponseUser is used to return user data
type ResponseUser struct {
	User User `json:"user"`
}

// Transalation is used to return details of the text translated
type Transalation struct {
	SourceLanguageCode string
	TargetLanguageCode string
	Text               string
}
