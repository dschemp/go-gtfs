package types

type Agency struct {
	ID          `csv:"agency_id"`
	Name        string `csv:"agency_name"`
	URL         `csv:"agency_url"`
	Timezone    `csv:"agency_timezone"`
	Language    LanguageCode `csv:"agency_lang"`
	PhoneNumber `csv:"agency_phone"`
	FareURL     URL          `csv:"agency_fare_url"`
	Email       EmailAddress `csv:"agency_email"`
}
