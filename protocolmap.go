package asign

import (
	"regexp"
)

var (
	cmdLookup     = map[string][]byte{}
	cmdRegex      *regexp.Regexp
	cmdLabelRegex *regexp.Regexp
)

func init() {
	cmdMaps := []map[string]byte{
		Escape,
		ValidLabel,
		CommandCode,
	}

	for _, cmdMap := range cmdMaps {
		for k, v := range cmdMap {
			cmdLookup[k] = []byte{v}
		}
	}
	for k, v := range Color {
		cmdLookup[k] = []byte{COL, v}
	}
	for k, v := range ModeCode {
		cmdLookup[k] = []byte{ESC, DP_MIDDLE_LINE, v}
	}

	cmdLookup["SOT"] = []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, // auto baud
		SOH,      // SOH
		'Z',      // TC_ALLSIGNS
		'0', '0', // SA_BROADCAST
	}

	cmdRegex = regexp.MustCompile(cmdRegexpString(cmdLookup))
}

func cmdRegexpString(cmds map[string][]byte) string {
	s := `\{(`
	i := 0
	for k := range cmds {
		s += regexp.QuoteMeta(k)
		if i != len(cmds)-1 {
			s += "|"
		}
		i++
	}
	s += `)\}`
	return s
}

func Parse(tmpl []byte) (p []byte, err error) {
	p = cmdRegex.ReplaceAllFunc(tmpl, func(c []byte) (b []byte) {
		println(string(c))
		if hex, ok := cmdLookup[string(c[1:len(c)-1])]; ok {
			b = hex
		} else {
			b = c
		}
		return
	})
	return
}


var Escape = map[string]byte{
	"NUL": NUL,
	"SOH": SOH,
	"STX": STX,
	"ETX": ETX,
	"EOT": EOT,
	"ESC": ESC,
	"COL": COL,
}

var TypeCode = map[string]byte{
	"VisualVerification": TC_VISUAL_VERIFICATION,
	"SerialClock":        TC_SERIAL_CLOCK,
	"AllSigns":           TC_ALLSIGNS,
	"BetaBrite":          TC_BETABRITE,
}

var CommandCode = map[string]byte{
	"WriteText":         CC_WRITE_TEXT,
	"ReadText":          CC_READ_TEXT,
	"WriteSpecial":      CC_WRITE_SPECIAL,
	"ReadSpecial":       CC_READ_SPECIAL,
	"WriteString":       CC_WRITE_STRING,
	"ReadString":        CC_READ_STRING,
	"WriteSmallDotsPic": CC_WRITE_SMALL_DOTS_PIC,
	"ReadSmallDotsPic":  CC_READ_SMALL_DOTS_PIC,
	"WriteRGBPic":       CC_WRITE_RGB_PIC,
	"ReadRGBPic":        CC_READ_RGB_PIC,
	"SetTimeoutMsg":     CC_SET_TIMEOOUT_MSG,
}

var DisplayPosition = map[string]byte{
	"MiddleLine": DP_MIDDLE_LINE,
	"TopLine":    DP_TOP_LINE,
}

var ModeCode = map[string]byte{
	"Rotate":    MC_ROTATE,
	"Hold":      MC_HOLD,
	"Flash":     MC_FLASH,
	"RollUp":    MC_ROLL_UP,
	"RollDown":  MC_ROLL_DOWN,
	"RollLeft":  MC_ROLL_LEFT,
	"RollRight": MC_ROLL_RIGHT,
	"AutoMode":  MC_AUTO_MODE,
	"Special":   MC_SPECIAL,
}

var SpecialMode = map[string]byte{
	"Twinkle":           MC_S_TWINKLE,
	"Sparkle":           MC_S_SPARKLE,
	"Snow":              MC_S_SNOW,
	"Interlock":         MC_S_INTERLOCK,
	"Switch":            MC_S_SWITCH,
	"SlideColor":        MC_S_SLIDECOLOR,
	"Spray":             MC_S_SPRAY,
	"Starburst":         MC_S_STARBURST,
	"Welcome":           MC_S_WELCOME,
	"SlotMachine":       MC_S_SLOTMACHINE,
	"NewsFlash":         MC_S_NEWSFLASH,
	"TrumpetAnimation":  MC_S_TRUMPETANIMATION,
	"CycleColors":       MC_S_CYCLECOLORS,
	"NoSmoking":         MC_SG_NOSMOKING,
	"DontDrinkAndDrive": MC_SG_DONTDRINKANDDRIVE,
	"RunningAnimal":     MC_SG_RUNNINGANIMAL,
	"Fireworks":         MC_SG_FIREWORKS,
	"TurboCar":          MC_SG_TURBOCAR,
	"ThankYou":          MC_SG_THANKYOU,
}

var Color = map[string]byte{
	"Red":       0x31,
	"Green":     0x32,
	"Amber":     0x33,
	"DimRed":    0x34,
	"DimGreen":  0x35,
	"Brown":     0x36,
	"Orange":    0x37,
	"Yellow":    0x38,
	"Rainbow1":  0x39,
	"Rainbow2":  0x41,
	"ColorMix":  0x42,
	"AutoColor": 0x43,
}

var ExtendedCharacter = map[string]byte{
	"Euro":          EC_EURO,
	"YPunc":         EC_YPUNC,
	"UpArrow":       EC_UPARROW,
	"DownArrow":     EC_DOWNARROW,
	"PackMan":       EC_PACKMAN,
	"SailBoat":      EC_SAILBOAT,
	"Ball":          EC_BALL,
	"Telephone":     EC_TELEPHONE,
	"Heart":         EC_HEART,
	"Car":           EC_CAR,
	"Handicap":      EC_HANDICAP,
	"Rhino":         EC_RHINO,
	"Mug":           EC_MUG,
	"SatelliteDish": EC_SATELLITEDISH,
	"Copyright":     EC_COPYRIGHT,
	"Male":          EC_MALE,
	"Female":        EC_FEMALE,
	"Bottle":        EC_BOTTLE,
	"Diskette":      EC_DISKETTE,
	"Printer":       EC_PRINTER,
	"MusicalNote":   EC_MUSICALNOTE,
	"Infinity":      EC_INFINITY,
	"Temp1":         EC_SC_TEMP1,
	"Temp2":         EC_SC_TEMP2,
	"Counter1":      EC_SC_COUNTER1,
	"Counter2":      EC_SC_COUNTER2,
	"Counter3":      EC_SC_COUNTER3,
	"Counter4":      EC_SC_COUNTER4,
	"Counter5":      EC_SC_COUNTER5,
}

var ValidLabel = map[string]byte{
	" ":  0x20,
	"!":  0x21,
	"\"": 0x22,
	"#":  0x23,
	"$":  0x24,
	"%":  0x25,
	"&":  0x26,
	"'":  0x27,
	"(":  0x28,
	")":  0x29,
	"*":  0x2a,
	"+":  0x2b,
	",":  0x2c,
	"-":  0x2d,
	".":  0x2e,
	"/":  0x2f,
	"0":  0x30,
	"1":  0x31,
	"2":  0x32,
	"3":  0x33,
	"4":  0x34,
	"5":  0x35,
	"6":  0x36,
	"7":  0x37,
	"8":  0x38,
	"9":  0x39,
	":":  0x3a,
	";":  0x3b,
	"<":  0x3c,
	"=":  0x3d,
	">":  0x3e,
	"?":  0x3f,
	"@":  0x40,
	"A":  0x41,
	"B":  0x42,
	"C":  0x43,
	"D":  0x44,
	"E":  0x45,
	"F":  0x46,
	"G":  0x47,
	"H":  0x48,
	"I":  0x49,
	"J":  0x4a,
	"K":  0x4b,
	"L":  0x4c,
	"M":  0x4d,
	"N":  0x4e,
	"O":  0x4f,
	"P":  0x50,
	"Q":  0x51,
	"R":  0x52,
	"S":  0x53,
	"T":  0x54,
	"U":  0x55,
	"V":  0x56,
	"W":  0x57,
	"X":  0x58,
	"Y":  0x59,
	"Z":  0x5a,
	"[":  0x5b,
	"\\": 0x5c,
	"]":  0x5d,
	"^":  0x5e,
	"_":  0x5f,
	"`":  0x60,
	"a":  0x61,
	"b":  0x62,
	"c":  0x63,
	"d":  0x64,
	"e":  0x65,
	"f":  0x66,
	"g":  0x67,
	"h":  0x68,
	"i":  0x69,
	"j":  0x6a,
	"k":  0x6b,
	"l":  0x6c,
	"m":  0x6d,
	"n":  0x6e,
	"o":  0x6f,
	"p":  0x70,
	"q":  0x71,
	"r":  0x72,
	"s":  0x73,
	"t":  0x74,
	"u":  0x75,
	"v":  0x76,
	"w":  0x77,
	"x":  0x78,
	"y":  0x79,
	"z":  0x7a,
	"{":  0x7b,
	"|":  0x7c,
	"}":  0x7d,
	"~":  0x7e,
	// "" : 0x7f, // reserved
}
