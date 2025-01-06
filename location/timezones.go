package location

import (
	"time"
)

type Timezone struct{}

func Timezones() *Timezone {
	return &Timezone{}
}

const (
	UTC   = "UTC"   // Coordinated Universal Time
	Local = "Local" // Local Time

	// European Time Zones

	GMT  = "GMT"  // Greenwich Mean Time
	CET  = "CET"  // Central European Time
	CEST = "CEST" // Central European Summer Time

	// US Time Zones
	EST  = "EST"  // Eastern Standard Time
	EDT  = "EDT"  // Eastern Daylight Time
	CST  = "CST"  // Central Standard Time
	CDT  = "CDT"  // Central Daylight Time
	MST  = "MST"  // Mountain Standard Time
	MDT  = "MDT"  // Mountain Daylight Time
	PST  = "PST"  // Pacific Standard Time
	PDT  = "PDT"  // Pacific Daylight Time
	AST  = "AST"  // Atlantic Standard Time
	ADT  = "ADT"  // Atlantic Daylight Time
	HST  = "HST"  // Hawaii Standard Time
	AKST = "AKST" // Alaska Standard Time
	AKDT = "AKDT" // Alaska Daylight Time

	// Australian Time Zones

	AEST = "AEST" // Australian Eastern Standard Time
	AEDT = "AEDT" // Australian Eastern Daylight Time
	ACST = "ACST" // Australian Central Standard Time
	ACDT = "ACDT" // Australian Central Daylight Time
	AWST = "AWST" // Australian Western Standard Time

	// New Zealand

	NZST = "NZST" // New Zealand Standard Time
	NZDT = "NZDT" // New Zealand Daylight Time

	// Asian Time Zones

	IST = "IST" // India Standard Time
	JST = "JST" // Japan Standard Time
	KST = "KST" // Korea Standard Time
	HKT = "HKT" // Hong Kong Time
	SGT = "SGT" // Singapore Time
	PHT = "PHT" // Philippine Time

	// African Time Zones

	SAST = "SAST" // South Africa Standard Time
	EAT  = "EAT"  // East Africa Time

	// Russia and Eastern Europe

	MSK = "MSK" // Moscow Standard Time

	// Indonesian Time Zones
	WIB  = "WIB"  // Western Indonesia Time
	WITA = "WITA" // Central Indonesia Time
	WIT  = "WIT"  // Eastern Indonesia Time

	// Argentina
	ART = "ART" // Argentina Time
)

// GetAll returns a slice of all available time zones as *time.Location objects.
func (t *Timezone) GetAll() []*time.Location {
	return []*time.Location{
		time.UTC,                              // UTC (Universal Coordinated Time)
		time.Local,                            // Local Time
		time.FixedZone("GMT", 0),              // Greenwich Mean Time (UTC+0)
		time.FixedZone("CET", 1*3600),         // Central European Time (UTC+1)
		time.FixedZone("CEST", 2*3600),        // Central European Summer Time (UTC+2)
		time.FixedZone("EST", -5*3600),        // Eastern Standard Time (UTC-5)
		time.FixedZone("EDT", -4*3600),        // Eastern Daylight Time (UTC-4)
		time.FixedZone("CST", -6*3600),        // Central Standard Time (UTC-6)
		time.FixedZone("CDT", -5*3600),        // Central Daylight Time (UTC-5)
		time.FixedZone("MST", -7*3600),        // Mountain Standard Time (UTC-7)
		time.FixedZone("MDT", -6*3600),        // Mountain Daylight Time (UTC-6)
		time.FixedZone("PST", -8*3600),        // Pacific Standard Time (UTC-8)
		time.FixedZone("PDT", -7*3600),        // Pacific Daylight Time (UTC-7)
		time.FixedZone("AST", -4*3600),        // Atlantic Standard Time (UTC-4)
		time.FixedZone("ADT", -3*3600),        // Atlantic Daylight Time (UTC-3)
		time.FixedZone("HST", -10*3600),       // Hawaii Standard Time (UTC-10)
		time.FixedZone("AKST", -9*3600),       // Alaska Standard Time (UTC-9)
		time.FixedZone("AKDT", -8*3600),       // Alaska Daylight Time (UTC-8)
		time.FixedZone("AEST", 10*3600),       // Australian Eastern Standard Time (UTC+10)
		time.FixedZone("AEDT", 11*3600),       // Australian Eastern Daylight Time (UTC+11)
		time.FixedZone("ACST", 9*3600+30*60),  // Australian Central Standard Time (UTC+9:30)
		time.FixedZone("ACDT", 10*3600+30*60), // Australian Central Daylight Time (UTC+10:30)
		time.FixedZone("AWST", 8*3600),        // Australian Western Standard Time (UTC+8)
		time.FixedZone("NZST", 12*3600),       // New Zealand Standard Time (UTC+12)
		time.FixedZone("NZDT", 13*3600),       // New Zealand Daylight Time (UTC+13)
		time.FixedZone("IST", 5*3600+30*60),   // India Standard Time (UTC+5:30)
		time.FixedZone("JST", 9*3600),         // Japan Standard Time (UTC+9)
		time.FixedZone("KST", 9*3600),         // Korea Standard Time (UTC+9)
		time.FixedZone("SAST", 2*3600),        // South Africa Standard Time (UTC+2)
		time.FixedZone("WIB", 7*3600),         // Western Indonesia Time (UTC+7)
		time.FixedZone("WITA", 8*3600),        // Central Indonesia Time (UTC+8)
		time.FixedZone("WIT", 9*3600),         // Eastern Indonesia Time (UTC+9)
		time.FixedZone("ART", -3*3600),        // Argentina Time (UTC-3)
		time.FixedZone("EAT", 3*3600),         // East Africa Time (UTC+3)
		time.FixedZone("MSK", 3*3600),         // Moscow Standard Time (UTC+3)
		time.FixedZone("HKT", 8*3600),         // Hong Kong Time (UTC+8)
		time.FixedZone("SGT", 8*3600),         // Singapore Time (UTC+8)
		time.FixedZone("PHT", 8*3600),         // Philippine Time (UTC+8)
	}
}

// GetSingular returns a *time.Location corresponding to the given timezone abbreviation.
// The loc parameter should be a string representing a timezone abbreviation (e.g., "UTC", "GMT", "EST").
// If the abbreviation is recognized, the function returns a fixed zone *time.Location with the appropriate offset.
// If the abbreviation is not recognized, it defaults to returning the local time zone.
func (tz *Timezone) GetSingular(loc string) *time.Location {
	switch loc {
	case UTC:
		return time.UTC
	case Local:
		return time.Local
	case GMT:
		return time.FixedZone("GMT", 0)
	case CET:
		return time.FixedZone("CET", 1*3600)
	case CEST:
		return time.FixedZone("CEST", 2*3600)
	case EST:
		return time.FixedZone("EST", -5*3600)
	case EDT:
		return time.FixedZone("EDT", -4*3600)
	case CST:
		return time.FixedZone("CST", -6*3600)
	case CDT:
		return time.FixedZone("CDT", -5*3600)
	case MST:
		return time.FixedZone("MST", -7*3600)
	case MDT:
		return time.FixedZone("MDT", -6*3600)
	case PST:
		return time.FixedZone("PST", -8*3600)
	case PDT:
		return time.FixedZone("PDT", -7*3600)
	case AST:
		return time.FixedZone("AST", -4*3600)
	case ADT:
		return time.FixedZone("ADT", -3*3600)
	case HST:
		return time.FixedZone("HST", -10*3600)
	case AKST:
		return time.FixedZone("AKST", -9*3600)
	case AKDT:
		return time.FixedZone("AKDT", -8*3600)
	case AEST:
		return time.FixedZone("AEST", 10*3600)
	case AEDT:
		return time.FixedZone("AEDT", 11*3600)
	case ACST:
		return time.FixedZone("ACST", 9*3600+30*60)
	case ACDT:
		return time.FixedZone("ACDT", 10*3600+30*60)
	case AWST:
		return time.FixedZone("AWST", 8*3600)
	case NZST:
		return time.FixedZone("NZST", 12*3600)
	case NZDT:
		return time.FixedZone("NZDT", 13*3600)
	case IST:
		return time.FixedZone("IST", 5*3600+30*60)
	case JST:
		return time.FixedZone("JST", 9*3600)
	case KST:
		return time.FixedZone("KST", 9*3600)
	case HKT:
		return time.FixedZone("HKT", 8*3600)
	case SGT:
		return time.FixedZone("SGT", 8*3600)
	case PHT:
		return time.FixedZone("PHT", 8*3600)
	case SAST:
		return time.FixedZone("SAST", 2*3600)
	case EAT:
		return time.FixedZone("EAT", 3*3600)
	case MSK:
		return time.FixedZone("MSK", 3*3600)
	case WIB:
		return time.FixedZone("WIB", 7*3600)
	case WITA:
		return time.FixedZone("WITA", 8*3600)
	case WIT:
		return time.FixedZone("WIT", 9*3600)
	case ART:
		return time.FixedZone("ART", -3*3600)
	default:
		return time.Local
	}
}
