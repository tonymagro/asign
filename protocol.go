package asign

const (
	NUL    = 0x00 // NULL
	SOH    = 0x01 // Start of Header
	STX    = 0x02 // Start of TeXt
	ETX    = 0x03 // End of TeXt
	EOT    = 0x04 // End of Transmission
	ESC    = 0x1b
	COL    = 0x1c
	STRING = 0x10 // String
	DOT    = 0x14
)

// Type Codes
const (
	TC_VISUAL_VERIFICATION = 0x21
	TC_SERIAL_CLOCK        = 0x22
	TC_ALLSIGNS            = 0x5a
	TC_BETABRITE           = 0x5e
)

// Sign Address
const (
	SA_BROADCAST = "00"
)

// Command Code
const (
	CC_WRITE_TEXT           = 0x41
	CC_READ_TEXT            = 0x42
	CC_WRITE_SPECIAL        = 0x45
	CC_READ_SPECIAL         = 0x46
	CC_WRITE_STRING         = 0x47
	CC_READ_STRING          = 0x48
	CC_WRITE_SMALL_DOTS_PIC = 0x49
	CC_READ_SMALL_DOTS_PIC  = 0x4a
	CC_WRITE_RGB_PIC        = 0x4b
	CC_READ_RGB_PIC         = 0x4c
	CC_SET_TIMEOOUT_MSG     = 0x54
)

// Display Position
const (
	DP_MIDDLE_LINE = 0x20
	DP_TOP_LINE    = 0X22
	DP_BOTTOM_LINE = 0x26
	DP_FILL        = 0x30
	DP_LEFT        = 0x31
	DP_RIGHT       = 0x32
)

// Mode Code
const (
	MC_ROTATE = 0x61
	MC_HOLD   = 0x62
	MC_FLASH  = 0x63
	//MC_reserved = 0x64
	MC_ROLL_UP           = 0x65
	MC_ROLL_DOWN         = 0x66
	MC_ROLL_LEFT         = 0x67
	MC_ROLL_RIGHT        = 0x68
	MC_WIPE_UP           = 0x69
	MC_WIPE_DOWN         = 0x6a
	MC_WIPE_LEFT         = 0x6b
	MC_WIPE_RIGHT        = 0x6c
	MC_SCROLL            = 0x6d
	MC_AUTO_MODE         = 0x6f
	MC_ROLL_IN           = 0x70
	MC_ROLL_OUT          = 0x71
	MC_WIPE_IN           = 0x72
	MC_WIPE_OUT          = 0x73
	MC_COMPRESSED_ROTATE = 0x74
	MC_EXPLODE           = 0x75
	MC_CLOCK             = 0x76
	MC_SPECIAL           = 0x6e
)

// Special Modes
const (
	MC_S_TWINKLE            = 0x30
	MC_S_SPARKLE            = 0x31
	MC_S_SNOW               = 0x32
	MC_S_INTERLOCK          = 0x33
	MC_S_SWITCH             = 0x34
	MC_S_SLIDECOLOR         = 0x35
	MC_S_SPRAY              = 0x36
	MC_S_STARBURST          = 0x37
	MC_S_WELCOME            = 0x38
	MC_S_SLOTMACHINE        = 0x39
	MC_S_NEWSFLASH          = 0x3a
	MC_S_TRUMPETANIMATION   = 0x3b
	MC_S_CYCLECOLORS        = 0x43
	MC_SG_THANKYOU          = 0x53
	MC_SG_NOSMOKING         = 0x55
	MC_SG_DONTDRINKANDDRIVE = 0x56
	MC_SG_RUNNINGANIMAL     = 0x57
	MC_SG_FIREWORKS         = 0x58
	MC_SG_TURBOCAR          = 0x59
	MC_SG_CHERRYBOMB        = 0x5a
)

// Color
const (
	C_RED       = 0x31
	C_GREEN     = 0x32
	C_AMBER     = 0x33
	C_DIMRED    = 0x34
	C_DIMGREEN  = 0x35
	C_BROWN     = 0x36
	C_ORANGE    = 0x37
	C_YELLOW    = 0x38
	C_RAINBOW1  = 0x39
	C_RAINBOW2  = 0x41
	C_COLORMIX  = 0x42
	C_AUTOCOLOR = 0x43
)

// Extended Character Set - combined with 0x08 ^H
const (
	EC_EURO          = 0x63
	EC_YPUNC         = 0x63
	EC_UPARROW       = 0x64
	EC_DOWNARROW     = 0x65
	EC_PACKMAN       = 0x68
	EC_SAILBOAT      = 0x69
	EC_BALL          = 0x6a
	EC_TELEPHONE     = 0x6b
	EC_HEART         = 0x6c
	EC_CAR           = 0x6d
	EC_HANDICAP      = 0x6e
	EC_RHINO         = 0x6f
	EC_MUG           = 0x70
	EC_SATELLITEDISH = 0x71
	EC_COPYRIGHT     = 0x72
	EC_MALE          = 0x73
	EC_FEMALE        = 0x74
	EC_BOTTLE        = 0x75
	EC_DISKETTE      = 0x76
	EC_PRINTER       = 0x77
	EC_MUSICALNOTE   = 0x78
	EC_INFINITY      = 0x79
	EC_SC_TEMP1      = 0x1c
	EC_SC_TEMP2      = 0x1d
	EC_SC_COUNTER1   = 0x7a
	EC_SC_COUNTER2   = 0x7b
	EC_SC_COUNTER3   = 0x7c
	EC_SC_COUNTER4   = 0x7d
	EC_SC_COUNTER5   = 0x7e
)

// Special Functions
const (
	SF_SetTimeOfDay    = 0x20
	SF_Speaker         = 0x21
	SF_SetMemoryConfig = 0x24
)

// Memory Keyboard Protection
const (
	M_KP_LOCKED   = 0x4c
	M_KP_UNLOCKED = 0x55
)

// Memory File Type
const (
	M_FT_TEXT   = 0x41
	M_FT_STRING = 0x42
	M_FT_DOTS   = 0x43
)
