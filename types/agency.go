package types

type Agency struct {
	ID
	Name string
	URL
	Timezone
	Language LanguageCode
	PhoneNumber
	FareURL URL
	Email   EmailAddress
}
