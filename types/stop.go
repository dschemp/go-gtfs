package types

type LocationTypeEnum Enum

const (
	LocationTypeStop LocationTypeEnum = iota
	LocationTypeStation
	LocationTypeEntranceOrExit
	LocationTypeGeneric
	LocationTypeBoardingArea
)

type WheelchairBoardingEnum Enum

const (
	WheelchairBoardingNoInfo WheelchairBoardingEnum = iota
	WheelchairBoardingSomeAccessiblePath
	WheelchairBoardingNoAccessiblePath
)

type Stop struct {
	ID
	Code        string
	Name        string
	Description string
	Latitude
	Longitude
	ZoneID        ID
	StopURL       URL
	LocationType  LocationTypeEnum
	ParentStation ID
	Timezone
	WheelchairBoarding WheelchairBoardingEnum
	LevelID            ID
	PlatformCode       string
}
