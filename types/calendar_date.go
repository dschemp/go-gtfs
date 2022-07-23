package types

type ExceptionTypeEnum Enum

const (
	ExceptionTypeServiceAdded ExceptionTypeEnum = iota + 1
	ExceptionTypeServiceRemoved
)

type CalendarDate struct {
	ServiceID     ID `csv:"service_id"`
	Date          `csv:"date"`
	ExceptionType ExceptionTypeEnum `csv:"exception_type"`
}
