package prompt

import "github.com/cloudfoundry/jibber_jabber"

type promptFunc interface {
	Prompt(presuppose string)
	InitPrompt()
}

var UserLanguage string
var promptMethod promptFunc

var languagePromptMap = make(map[string]promptFunc)

func initLanguagePromptMap() {
	languagePromptMap["zh-CN"] = &ChinesePrompt{}
	languagePromptMap["zh-TW"] = &ChinesePrompt{}
	languagePromptMap["en-US"] = &EnglishPrompt{}
}

// GetLanguage returns the user's language setting.
func GetLanguage() string {
	userLocale, _ := jibber_jabber.DetectIETF()
	return userLocale
}

// Prompt displays a prompt message to the user.
func Prompt(presuppose string) {
	promptMethod.Prompt(presuppose)
}

func init() {
	UserLanguage = GetLanguage()
	initLanguagePromptMap()
	promptMethod = languagePromptMap[UserLanguage]
	if promptMethod == nil {
		promptMethod = languagePromptMap["en-US"]
	}
	promptMethod.InitPrompt()

}
