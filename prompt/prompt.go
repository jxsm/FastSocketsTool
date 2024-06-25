package prompt

import "github.com/cloudfoundry/jibber_jabber"

// GetLanguage returns the user's language setting.
func GetLanguage() string {
	userLocale, _ := jibber_jabber.DetectIETF()
	return userLocale
}
