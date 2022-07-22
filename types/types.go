// Package types Definitions from https://developers.google.com/transit/gtfs/reference#field_types
//
// Defines some base types used in GTFS.
package types

// HexColor is a textual representation of a hex color without the leading '#'.
//
// Example: "FFFFFF" for white.
type HexColor string

// CurrencyCode is a currency code as defined by ISO 4217 standard.
//
// Example: "EUR" for Euro.
type CurrencyCode string

// Date is a single day as represented in the YYYYMMDD format.
//
// Example: 20180913 for September 13th, 2018.
type Date string

// EmailAddress is a simple email address.
//
// Example: jane.doe@example.com
type EmailAddress string

// Enum represents any sort of enumeration type with a defined set of possible values.
type Enum uint

// ID refers to any internal identifier used by GTFS.
// They can be any sequence of UTF-8 characters but is recommended to be a human-readable ASCII sequence.
// They are typically used to identify entries between tables.
type ID string

// LanguageCode refers to a language code as defined by IETF BCP 47.
//
// Examples: 'en' for English, 'en-US' for American English, 'de' for German.
type LanguageCode string

// Latitude is a WGS84 latitude in decimal degrees.
// It can range between -90.0 up to +90.0.
type Latitude float64

// Longitude is a WGS84 longitude in decimal degrees.
// It can range between -180.0 up to +180.0.
type Longitude float64

// PhoneNumber is a simple phone number is no specified format.
type PhoneNumber string

// Time is a time in the HH:MM:SS (or H:MM:SS) format.
// The time does not necessarily need to be between 00:00:00 and 23:59:59 but can refers to time on other days.
//
// Examples: 14:30:00 or 25:35:00 (01:35:00 on the next day).
type Time string

// Timezone is a string containing a valid timezone as defined by IANA.
// Spaces are escaped with underscores.
// List of valid timezones: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//
// Examples: Asia/Tokyo, America/Los_Angeles, Africa/Cairo
type Timezone string

// URL refers to any URI starting with http:// or https://.
// The URL has to be escaped properly.
type URL string
