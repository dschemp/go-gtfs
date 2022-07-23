package types

type PickupTypeEnum Enum

const (
	PickupTypeRegularlyScheduled PickupTypeEnum = iota
	PickupTypeNoPickupAvailable
	PickupTypeArrangeWithAgency
	PickupTypeArrangeWithDriver
)

type DropOffTypeEnum Enum

const (
	DropOffTypeRegularlyScheduled DropOffTypeEnum = iota
	DropOffTypeNoDropOffAvailable
	DropOffTypeArrangeWithAgency
	DropOffTypeArrangeWithDriver
)

type TimepointEnum Enum

const (
	TimepointApproximate TimepointEnum = iota
	TimepointExact
)

type StopTime struct {
	TripID            ID                    `csv:"trip_id"`
	ArrivalTime       Time                  `csv:"arrival_time"`
	DepartureTime     Time                  `csv:"departure_time"`
	StopID            ID                    `csv:"stop_id"`
	Sequence          uint                  `csv:"stop_sequence"`
	Destination       string                `csv:"stop_headsign"` // headsign
	PickupType        PickupTypeEnum        `csv:"pickup_type"`
	DropOffType       DropOffTypeEnum       `csv:"drop_off_type"`
	ContinuousPickup  ContinuousPickupEnum  `csv:"continuous_pickup"`
	ContinuousDropOff ContinuousDropOffEnum `csv:"continuous_drop_off"`
	DistanceTraveled  float64               `csv:"shape_dist_traveled"`
	Timepoint         Enum                  `csv:"timepoint"`
}
