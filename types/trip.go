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
	ID                   `csv:"trip_id"`
	RouteID              ID                       `csv:"route_id"`
	ServiceID            ID                       `csv:"service_id"`
	Destination          string                   `csv:"trip_headsign"` // headsign
	ShortName            string                   `csv:"trip_short_name"`
	Direction            DirectionEnum            `csv:"direction_id"`
	BlockID              ID                       `csv:"block_id"`
	ShapeID              ID                       `csv:"shape_id"`
	WheelchairAccessible WheelchairAccessibleEnum `csv:"wheelchair_accessible"`
	BikesAllowed         BikesAllowedEnum         `csv:"bikes_allowed"`
}
