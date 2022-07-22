package types

type DirectionEnum Enum

const (
	DirectionOutbound DirectionEnum = iota
	DirectionInbound
)

type WheelchairAccessibleEnum Enum

const (
	WheelchairAccessibleNoInfo WheelchairAccessibleEnum = iota
	WheelchairAccessibleAtLeastOneRider
	WheelchairAccessibleNoWheelchairs
)

type BikesAllowedEnum Enum

const (
	BikesAllowedNoInfo BikesAllowedEnum = iota
	BikesAllowedAtLeastOneBicycle
	BikesAllowedNoBicyclesAllowed
)

type Trip struct {
	ID
	RouteID              ID
	ServiceID            ID
	Destination          string // headsign
	ShortName            string
	Direction            DirectionEnum
	BlockID              ID
	ShapeID              ID
	WheelchairAccessible WheelchairAccessibleEnum
	BikesAllowed         BikesAllowedEnum
}
