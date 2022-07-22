package types

type Trip struct {
	ID
	RouteID              ID
	ServiceID            ID
	Destination          string // headsign
	ShortName            string
	DirectionID          Enum
	BlockID              ID
	ShapeID              ID
	WheelchairAccessible Enum
	BikesAllowed         Enum
}
