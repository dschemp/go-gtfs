package types

type Stop struct {
	ID
	Code        string
	Name        string
	Description string
	Latitude
	Longitude
	ZoneID        ID
	StopURL       URL
	LocationType  Enum
	ParentStation ID
	Timezone
	WheelchairBoarding Enum
	LevelID            ID
	PlatformCode       string
}
