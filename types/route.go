package types

type Route struct {
	ID
	AgencyID    ID
	ShortName   string
	LongName    string
	Description string
	Type        Enum
	URL
	Color             HexColor
	TextColor         HexColor
	ContinuousPickup  Enum
	ContinuousDropOff Enum
}
