package types

type StopTime struct {
	TripID            ID
	ArrivalTime       Time
	DepartureTime     Time
	StopID            ID
	Sequence          uint
	Destination       string // headsign
	PickupType        Enum
	DropOffType       Enum
	ContinuousPickup  Enum
	ContinuousDropOff Enum
	DistanceTraveled  float64
	Timepoint         Enum
}
