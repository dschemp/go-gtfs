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
	ID                `csv:"route_id"`
	AgencyID          ID     `csv:"agency_id"`
	ShortName         string `csv:"route_short_name"`
	LongName          string `csv:"route_long_name"`
	Description       string `csv:"route_desc"`
	Type              Enum   `csv:"route_type"`
	URL               `csv:"route_url"`
	Color             HexColor              `csv:"route_color"`
	TextColor         HexColor              `csv:"route_text_color"`
	ContinuousPickup  ContinuousPickupEnum  `csv:"continuous_pickup"`
	ContinuousDropOff ContinuousDropOffEnum `csv:"continuous_drop_off"`
}
