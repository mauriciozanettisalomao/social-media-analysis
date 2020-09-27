package factory

import (
	"github.com/mauriciozanettisalomao/social-media-analysis/internal/app"
	"github.com/mauriciozanettisalomao/social-media-analysis/internal/app/socialmedia/twitter"
	translator "github.com/mauriciozanettisalomao/social-media-analysis/internal/app/translation"
)

// NewSocialMedia returns a new instance of a social media
func NewSocialMedia(socialMedia string) app.SocialMedia {

	switch socialMedia {
	case twitter.Name:
		return twitter.New()
	default:
		return nil
	}

}

// NewTranslator returns a new instance of a translator
func NewTranslator(sourceLanguageCode, targetLanguageCode string) app.Translator {
	return translator.New(sourceLanguageCode, targetLanguageCode)
}
