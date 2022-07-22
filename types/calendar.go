package types

type Calendar struct {
	ServiceID ID
	Monday    Enum
	Tuesday   Enum
	Wednesday Enum
	Thursday  Enum
	Friday    Enum
	Saturday  Enum
	Sunday    Enum
	StartDate Date
	EndDate   Date
}
