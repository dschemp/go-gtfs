package types

type ServiceEnum Enum

const (
	ServiceUnavailable ServiceEnum = iota
	ServiceAvailable
)

type Calendar struct {
	ServiceID ID
	Monday    ServiceEnum
	Tuesday   ServiceEnum
	Wednesday ServiceEnum
	Thursday  ServiceEnum
	Friday    ServiceEnum
	Saturday  ServiceEnum
	Sunday    ServiceEnum
	StartDate Date
	EndDate   Date
}
