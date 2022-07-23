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
	ID                 `csv:"stop_id"`
	Code               string `csv:"stop_code"`
	Name               string `csv:"stop_name"`
	Description        string `csv:"stop_desc"`
	Latitude           `csv:"stop_lat"`
	Longitude          `csv:"stop_lon"`
	ZoneID             ID `csv:"zone_id"`
	URL                `csv:"stop_url"`
	LocationType       LocationTypeEnum `csv:"location_type"`
	ParentStation      ID               `csv:"parent_station"`
	Timezone           `csv:"stop_timezone"`
	WheelchairBoarding WheelchairBoardingEnum `csv:"wheelchair_boarding"`
	LevelID            ID                     `csv:"level_id"`
	PlatformCode       string                 `csv:"platform_code"`
}
