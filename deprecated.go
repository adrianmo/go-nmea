package nmea

type (
	// GLGSV represents the GPS Satellites in view http://aprs.gids.nl/nmea/#glgsv
	// Deprecated
	GLGSV = GSV

	// GLGSVInfo represents information about a visible satellite
	// Deprecated
	GLGSVInfo = GSVInfo

	// GNGGA is the Time, position, and fix related data of the receiver.
	// Deprecated
	GNGGA = GGA

	// GNGNS is standard GNSS sentance that combined multiple constellations
	// Deprecated
	GNGNS = GNS

	// GNRMC is the Recommended Minimum Specific GNSS data. http://aprs.gids.nl/nmea/#rmc
	// Deprecated
	GNRMC = RMC

	// GPGGA represents fix data. http://aprs.gids.nl/nmea/#gga
	// Deprecated
	GPGGA = GGA

	// GPGLL is Geographic Position, Latitude / Longitude and time. http://aprs.gids.nl/nmea/#gll
	// Deprecated
	GPGLL = GLL

	// GPGSA represents overview satellite data. http://aprs.gids.nl/nmea/#gsa
	// Deprecated
	GPGSA = GSA

	// GPGSV represents the GPS Satellites in view http://aprs.gids.nl/nmea/#gpgsv
	// Deprecated
	GPGSV = GSV

	// GPGSVInfo represents information about a visible satellite
	// Deprecated
	GPGSVInfo = GSVInfo

	// GPHDT is the Actual vessel heading in degrees True. http://aprs.gids.nl/nmea/#hdt
	// Deprecated
	GPHDT = HDT

	// GPRMC is the Recommended Minimum Specific GNSS data. http://aprs.gids.nl/nmea/#rmc
	// Deprecated
	GPRMC = RMC

	// GPVTG represents track & speed data. http://aprs.gids.nl/nmea/#vtg
	// Deprecated
	GPVTG = VTG

	// GPZDA represents date & time data. http://aprs.gids.nl/nmea/#zda
	// Deprecated
	GPZDA = ZDA
)

const (
	// PrefixGNGNS prefix
	// Deprecated
	PrefixGNGNS = "GNGNS"

	// PrefixGPGGA prefix
	// Deprecated
	PrefixGPGGA = "GPGGA"

	// PrefixGPGLL prefix for GPGLL sentence type
	// Deprecated
	PrefixGPGLL = "GPGLL"

	// PrefixGPGSA prefix of GPGSA sentence type
	// Deprecated
	PrefixGPGSA = "GPGSA"

	// PrefixGPRMC prefix of GPRMC sentence type
	// Deprecated
	PrefixGPRMC = "GPRMC"

	// PrefixPGRME prefix for PGRME sentence type
	// Deprecated
	PrefixPGRME = "PGRME"

	// PrefixGLGSV prefix
	// Deprecated
	PrefixGLGSV = "GLGSV"

	// PrefixGNGGA prefix
	// Deprecated
	PrefixGNGGA = "GNGGA"

	// PrefixGNRMC prefix of GNRMC sentence type
	// Deprecated
	PrefixGNRMC = "GNRMC"

	// PrefixGPGSV prefix
	// Deprecated
	PrefixGPGSV = "GPGSV"

	// PrefixGPHDT prefix of GPHDT sentence type
	// Deprecated
	PrefixGPHDT = "GPHDT"

	// PrefixGPVTG prefix
	// Deprecated
	PrefixGPVTG = "GPVTG"

	// PrefixGPZDA prefix
	// Deprecated
	PrefixGPZDA = "GPZDA"
)
