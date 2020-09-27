package app

// SocialMedia is an interface to handle social medias as twitter, facebook, instagram, etc
type SocialMedia interface {
	Search(param string, filter map[string]string) ([]Mention, error)
}

// Translator is an interface to handle the transaltion when requested
type Translator interface {
	Translate(text string) (Transalation, error)
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

// Response is used to return all the informations of social media mentions
type Response struct {
	Parameter   string               `json:"parameter"`
	Filters     []Filter             `json:"filters"`
	SocialMedia map[string][]Mention `json:"socialMedia"`
}

// Transalation is used to return details of the text translated
type Transalation struct {
	SourceLanguageCode string
	TargetLanguageCode string
	Text               string
}
