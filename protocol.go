package asign

const (
	NUL = 0x00 // NULL
	SOH = 0x01 // Start of Header
	STX = 0x02 // Start of TeXt
	ETX = 0x03 // End of TeXt
	EOT = 0x04 // End of Transmission
	ESC = 0x1b
	COL = 0x1c
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
	MC_ROLL_UP    = 0x65
	MC_ROLL_DOWN  = 0x66
	MC_ROLL_LEFT  = 0x67
	MC_ROLL_RIGHT = 0x68
	MC_AUTO_MODE  = 0x6f
	MC_SPECIAL    = 0x6e
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
