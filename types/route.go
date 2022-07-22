package types

type RouteTypeEnum Enum

const (
	RouteTypeLightRail RouteTypeEnum = iota
	RouteTypeUnderground
	RouteTypeRail
	RouteTypeBus
	RouteTypeFerry
	RouteTypeCableTram
	RouteTypeAerialLift
	RouteTypeFunicular
	RouteTypeTrolleybus = 11
	RouteTypeMonorail   = 12
)

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
	ContinuousPickup  ContinuousPickupEnum
	ContinuousDropOff ContinuousDropOffEnum
}
