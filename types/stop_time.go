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
	TripID            ID
	ArrivalTime       Time
	DepartureTime     Time
	StopID            ID
	Sequence          uint
	Destination       string // headsign
	PickupType        Enum
	DropOffType       Enum
	ContinuousPickup  ContinuousPickupEnum
	ContinuousDropOff ContinuousDropOffEnum
	DistanceTraveled  float64
	Timepoint         Enum
}
