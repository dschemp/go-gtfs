package types

type ExceptionTypeEnum Enum

const (
	ExceptionTypeServiceAdded ExceptionTypeEnum = iota + 1
	ExceptionTypeServiceRemoved
)

type CalendarDate struct {
	ServiceID ID
	Date
	ExceptionType ExceptionTypeEnum
}
