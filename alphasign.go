package asign

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"time"
)

type ASign struct {
	Rw                io.ReadWriter
	Address           string
	TypeCode          byte
	DisableAutoMemory bool
	StxDelay          time.Duration
	PacketDelay       time.Duration
}

func SOT(typeCode byte, address string) []byte {
	// Start of Transmission
	sot := []byte{
		NUL, NUL, NUL, NUL, NUL, // Baud Rate
		SOH,      // Start of Header
		typeCode, // Sign Type Code
	}
	sot = append(sot, address...)
	return sot
}
func New(rw io.ReadWriter) *ASign {
	return &ASign{
		rw,
		SA_BROADCAST,
		TC_ALLSIGNS,
		false,
		100,
		100,
	}
}

func (s *ASign) SOT() []byte {
	return SOT(s.TypeCode, s.Address)
}

func (s *ASign) Write(p []byte) (n int, err error) {
	parts := bytes.SplitAfter(p, []byte{STX})
	var pn int
	for _, p := range parts {
		pn, err = s.Rw.Write(p)
		n += pn
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond * s.StxDelay)
	}
	return
}

var stxToEtx = regexp.MustCompile(fmt.Sprintf("%c[^%c]*%c", STX, ETX, ETX))
func (s *ASign) WriteTemplate(text []byte) (n int, err error) {
	cmd, err := Parse(text)
	cmd = bytes.Replace(cmd, []byte{'\n'}, []byte{}, -1)
	if err != nil {
		return
	}
	if !s.DisableAutoMemory {
		_, err = s.WriteAutoMemoryForCmd(cmd)
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond * s.PacketDelay)
	}
	return s.Write(cmd)
}

func (s *ASign) WriteAutoMemoryForCmd(cmd []byte) (n int, err error) {
	p := s.SOT()
	memCmd, err := AutoMemory(cmd)
	if err != nil {
		return
	}
	p = append(p, memCmd...)
	p = append(p, EOT)
	s.Write(p)
	return
}

func (s *ASign) Reset() (err error) {
	_, err = s.Write(ResetCmd())
	return
}

func ResetCmd() []byte {
	return []byte{
		NUL, NUL, NUL, NUL, NUL,
		NUL, NUL, NUL, NUL, NUL,
		NUL, NUL, NUL, NUL, NUL,
		NUL, NUL, NUL, NUL, NUL,
	}
}

func AutoMemory(cmd []byte) (p []byte, err error) {
	p = []byte{
		STX,
		CC_WRITE_SPECIAL,
		SF_SetMemoryConfig,
	}
	stxs := stxToEtx.FindAll(cmd, -1)
	found := false
	for _, stx := range stxs {
		size := len(stx)
		if size < 3 {
			continue
		}
		label := stx[2]
		cmdCode := stx[1]

		switch cmdCode {
		case CC_WRITE_TEXT:
			p = append(p, TextMemoryCmd(label, size)...)
			found = true
		case CC_WRITE_STRING:
		case CC_WRITE_SMALL_DOTS_PIC:
		default:
			continue
		}
	}
	if !found {
		err = errors.New("unable to auto allocate memory - no labels found in command")
	}
	return
}

func TextMemoryCmd(label byte, size int) (p []byte) {
	p = []byte{
		label,
		M_FT_TEXT,
		M_KP_LOCKED,
	}
	p = append(p, fmt.Sprintf("%04X", size)...)
	p = append(p, "FF00"...)
	return
}

/*
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"text/template"
	"time"
)
type ASign struct {
	Rw                io.ReadWriter
	Address           string
	TypeCode          byte
	DisableAutoMemory bool
	StxDelay          time.Duration
	PacketDelay       time.Duration
}

func (s *ASign) SOT() []byte {
	return SOT(s.TypeCode, s.Address)
}

func SOT(typeCode byte, address string) []byte {
	// Start of Transmission
	sot := []byte{
		NUL, NUL, NUL, NUL, NUL, // Baud Rate
		SOH,      // Start of Header
		typeCode, // Sign Type Code
	}
	sot = append(sot, address...)
	return sot
}

func (s *ASign) Read(p []byte) (n int, err error) {
	log.Println("READ not implemented")
	return
}

func (s *ASign) Close() {
	return
}

func (s *ASign) Reset() (err error) {
	_, err = s.Write(ResetCmd())
	return
}

func ResetCmd() []byte {
	return []byte{
		NUL, NUL, NUL, NUL, NUL,
		NUL, NUL, NUL, NUL, NUL,
		NUL, NUL, NUL, NUL, NUL,
		NUL, NUL, NUL, NUL, NUL,
	}
}

func (s *ASign) Write(p []byte) (n int, err error) {
	parts := bytes.SplitAfter(p, []byte{STX})
	var pn int
	for _, p := range parts {
		pn, err = s.Rw.Write(p)
		n += pn
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond * s.StxDelay)
	}
	return
}

func (s *ASign) WriteTemplate(text string) (n int, err error) {
	p := s.SOT()

	cmd, err := Parse(text)
	if err != nil {
		return
	}
	if !s.DisableAutoMemory {
		_, err = s.WriteAutoMemoryForCmd(cmd)
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond * s.PacketDelay)
	}
	p = append(p, cmd...)
	p = append(p, EOT)
	return s.Write(p)
}

func (s *ASign) WriteAutoMemoryForCmd(cmd []byte) (n int, err error) {
	p := s.SOT()
	memCmd, err := AutoMemory(cmd)
	if err != nil {
		return
	}
	p = append(p, memCmd...)
	p = append(p, EOT)
	s.Write(p)
	return
}

func Parse(text string) (p []byte, err error) {
	tmpl, err := template.New("sub").Funcs(funcMap).Parse(text)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, CmdLookup)
	if err != nil {
		return
	}
	p = buf.Bytes()
	return
}

var stxToEtx = regexp.MustCompile(fmt.Sprintf("%c[^%c]*%c", STX, ETX, ETX))

func AutoMemory(cmd []byte) (p []byte, err error) {
	p = []byte{
		STX,
		CC_WRITE_SPECIAL,
		SF_SetMemoryConfig,
	}
	stxs := stxToEtx.FindAll(cmd, -1)
	found := false
	for _, stx := range stxs {
		size := len(stx)
		if size < 3 {
			continue
		}
		label := stx[2]
		cmdCode := stx[1]

		switch cmdCode {
		case CC_WRITE_TEXT:
			p = append(p, TextMemoryCmd(label, size)...)
			found = true
		case CC_WRITE_STRING:
		case CC_WRITE_SMALL_DOTS_PIC:
		default:
			continue
		}
	}
	if !found {
		err = errors.New("unable to auto allocate memory - no labels found in command")
	}
	return
}

func TextMemoryCmd(label byte, size int) (p []byte) {
	p = []byte{
		label,
		M_FT_TEXT,
		M_KP_LOCKED,
	}
	p = append(p, fmt.Sprintf("%04X", size)...)
	p = append(p, "FF00"...)
	return
}


*/
func PacketString(pack []byte) string {
	str := fmt.Sprintf("===== Packet (Size: %v) =====\n", len(pack))
	str += fmt.Sprintln("DEC:", pack)
	str += "HEX: ["
	for i, c := range pack {
		str += fmt.Sprintf("%x", c)
		if i != len(pack)-1 {
			str += " "
		}
	}
	str += "]\n"
	var esc string
	for _, c := range pack {
		switch c {
		case NUL:
			esc += "<NUL>"
		case SOH:
			esc += "<SOH>"
		case STX:
			esc += "<STX>"
		case ETX:
			esc += "<ETX>"
		case EOT:
			esc += "<EOT>"
		case ESC:
			esc += "<ESC>"
		default:
			esc += fmt.Sprintf("%c", c)
		}
	}
	str += "ESC: " + strconv.Quote(esc)
	return str
}
