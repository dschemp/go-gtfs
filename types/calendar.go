package types

type ServiceEnum Enum

const (
	ServiceUnavailable ServiceEnum = iota
	ServiceAvailable
)

type Calendar struct {
	ServiceID ID          `csv:"service_id"`
	Monday    ServiceEnum `csv:"monday"`
	Tuesday   ServiceEnum `csv:"tuesday"`
	Wednesday ServiceEnum `csv:"wednesday"`
	Thursday  ServiceEnum `csv:"thursday"`
	Friday    ServiceEnum `csv:"friday"`
	Saturday  ServiceEnum `csv:"saturday"`
	Sunday    ServiceEnum `csv:"sunday"`
	StartDate Date        `csv:"start_date"`
	EndDate   Date        `csv:"end_date"`
}
