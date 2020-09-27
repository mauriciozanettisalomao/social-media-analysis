package translator

/*
https://docs.aws.amazon.com/translate/latest/dg/what-is.html
https://aws.amazon.com/pt/translate/pricing/
*/

import (
	"fmt"

	app "github.com/mauriciozanettisalomao/social-media-analysis/internal/app"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

const (
	awsRegion = "us-west-2"
)

type translator struct {
	SourceLanguageCode string
	TargetLanguageCode string
}

// Translate returns the text translated to target language
func (t translator) Translate(text string) (app.Transalation, error) {
	mySession := session.Must(session.NewSession())

	// Create a Translate client from just a session.
	svc := translate.New(mySession, aws.NewConfig().WithRegion(awsRegion))

	textInput := &translate.TextInput{
		SourceLanguageCode: &t.SourceLanguageCode,
		TargetLanguageCode: &t.TargetLanguageCode,
		Text:               &text,
	}

	output, err := svc.Text(textInput)
	if err != nil {
		return app.Transalation{}, fmt.Errorf("could not translate the text: %v", err)
	}

	return app.Transalation{
		SourceLanguageCode: *output.SourceLanguageCode,
		TargetLanguageCode: *output.TargetLanguageCode,
		Text:               *output.TranslatedText,
	}, nil
}

// New returns a new instance of a translator
func New(sourceLanguageCode, targetLanguageCode string) app.Translator {
	return translator{
		SourceLanguageCode: sourceLanguageCode,
		TargetLanguageCode: targetLanguageCode,
	}
}
