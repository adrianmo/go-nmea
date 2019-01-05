package nmea

type (
	// GLGSV represents the GPS Satellites in view http://aprs.gids.nl/nmea/#glgsv
	//
	// Deprecated: Use GSV instead
	GLGSV = GSV

	// GLGSVInfo represents information about a visible satellite
	//
	// Deprecated: Use GSVInfo instead
	GLGSVInfo = GSVInfo

	// GNGGA is the Time, position, and fix related data of the receiver.
	//
	// Deprecated: Use GGA instead
	GNGGA = GGA

	// GNGNS is standard GNSS sentance that combined multiple constellations
	//
	// Deprecated: Use GNS instead
	GNGNS = GNS

	// GNRMC is the Recommended Minimum Specific GNSS data. http://aprs.gids.nl/nmea/#rmc
	//
	// Deprecated: Use RCM instead
	GNRMC = RMC

	// GPGGA represents fix data. http://aprs.gids.nl/nmea/#gga
	//
	// Deprecated: Use GGA instead
	GPGGA = GGA

	// GPGLL is Geographic Position, Latitude / Longitude and time. http://aprs.gids.nl/nmea/#gll
	//
	// Deprecated: Use GLL instead
	GPGLL = GLL

	// GPGSA represents overview satellite data. http://aprs.gids.nl/nmea/#gsa
	//
	// Deprecated: Use GSA instead
	GPGSA = GSA

	// GPGSV represents the GPS Satellites in view http://aprs.gids.nl/nmea/#gpgsv
	//
	// Deprecated: Use GSV instead
	GPGSV = GSV

	// GPGSVInfo represents information about a visible satellite
	//
	// Deprecated: Use GSVInfo instead
	GPGSVInfo = GSVInfo

	// GPHDT is the Actual vessel heading in degrees True. http://aprs.gids.nl/nmea/#hdt
	//
	// Deprecated: Use HDT instead
	GPHDT = HDT

	// GPRMC is the Recommended Minimum Specific GNSS data. http://aprs.gids.nl/nmea/#rmc
	//
	// Deprecated: Use RMC instead
	GPRMC = RMC

	// GPVTG represents track & speed data. http://aprs.gids.nl/nmea/#vtg
	//
	// Deprecated: Use VTG instead
	GPVTG = VTG

	// GPZDA represents date & time data. http://aprs.gids.nl/nmea/#zda
	//
	// Deprecated: Use ZDA instead
	GPZDA = ZDA
)

const (
	// PrefixGNGNS prefix
	//
	// Deprecated: Use TypeGNS instead
	PrefixGNGNS = "GNGNS"

	// PrefixGPGGA prefix
	//
	// Deprecated: Use TypeGGA instead
	PrefixGPGGA = "GPGGA"

	// PrefixGPGLL prefix for GPGLL sentence type
	//
	// Deprecated: Use TypeGLL instead
	PrefixGPGLL = "GPGLL"

	// PrefixGPGSA prefix of GPGSA sentence type
	//
	// Deprecated: Use TypeGSA instead
	PrefixGPGSA = "GPGSA"

	// PrefixGPRMC prefix of GPRMC sentence type
	//
	// Deprecated: Use TypeRMC instead
	PrefixGPRMC = "GPRMC"

	// PrefixPGRME prefix for PGRME sentence type
	//
	// Deprecated: Use TypePGRME instead
	PrefixPGRME = "PGRME"

	// PrefixGLGSV prefix
	//
	// Deprecated: Use TypeGSV instead
	PrefixGLGSV = "GLGSV"

	// PrefixGNGGA prefix
	//
	// Deprecated: Use TypeGGA instead
	PrefixGNGGA = "GNGGA"

	// PrefixGNRMC prefix of GNRMC sentence type
	//
	// Deprecated: Use TypeRMC instead
	PrefixGNRMC = "GNRMC"

	// PrefixGPGSV prefix
	//
	// Deprecated: Use TypeGSV instead
	PrefixGPGSV = "GPGSV"

	// PrefixGPHDT prefix of GPHDT sentence type
	//
	// Deprecated: Use TypeHDT instead
	PrefixGPHDT = "GPHDT"

	// PrefixGPVTG prefix
	//
	// Deprecated: Use TypeVTG instead
	PrefixGPVTG = "GPVTG"

	// PrefixGPZDA prefix
	//
	// Deprecated: Use TypeZDA instead
	PrefixGPZDA = "GPZDA"
)
