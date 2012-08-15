package asign

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
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

func New(rw io.ReadWriter) *ASign {
	return NewSign(rw, SA_BROADCAST, TC_ALLSIGNS)
}

func NewSign(rw io.ReadWriter, address string, typeCode byte) *ASign {
	return &ASign{
		rw,
		address,
		typeCode,
		false,
		100,
		100,
	}
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

var textHeader = []byte{STX, CC_WRITE_TEXT}
var dotHeader = []byte{STX, CC_WRITE_SMALL_DOTS_PIC}
var stringHeader = []byte{STX, CC_WRITE_STRING}

func findLabel(cmd []byte, labelHeader []byte) (p []byte, err error) {
	start := bytes.Index(cmd, labelHeader)
	if start != -1 {
		etx := bytes.Index(cmd[start:], []byte{ETX})
		if etx != -1 {
			p = cmd[start : start+etx]
			return
		}
	}
	err = errors.New(fmt.Sprintf("label not found: %s", labelHeader))
	return
}

func AutoMemory(cmd []byte) (p []byte, err error) {
	p = []byte{
		STX,
		CC_WRITE_SPECIAL,
		SF_SetMemoryConfig,
	}
	found := false
	for _, v := range ValidLabel {
		labelHeader := append(textHeader, v)
		subCmd, notfound := findLabel(cmd, labelHeader)
		if notfound != nil {
			continue
		}
		found = true
		p = append(p, TextMemoryCmd(v, len(subCmd))...)
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

/*
func clearMemory() []byte {
	stp := sot(TC_ALLSIGNS, SA_BROADCAST)
	stp = append(stp, 'E', '$', EOT)
	return stp
}
func textMemoryConfig(fileLabel byte, size uint) []byte {
	stp := append(
		sot(TC_ALLSIGNS, SA_BROADCAST),
		CC_WRITE_SPECIAL,
		SF_SetMemoryConfig,
		fileLabel,
		M_FT_TEXT,
		M_KP_LOCKED,
	)
	stp = append(stp, fmt.Sprintf("%04X", size)...)
	stp = append(stp, "FF00"...)
	stp = append(stp, EOT)
	return stp
}

func makeTextPacket(ccp []byte, message string) []byte {
	stp := append(sot(TC_ALLSIGNS, SA_BROADCAST), ccp...)
	stp = append(stp, message...)
	stp = append(stp, EOT)
	return stp
}

func TextFile(label byte, msg string) []byte {
	ccp := []byte{
		CC_WRITE_TEXT,
		label,
	}
	return makeTextPacket(ccp, msg)
}

func writeTextMode(fileLabel byte, displayPos byte, mode byte, special byte, message string) []byte {
	ccp := []byte{
		CC_WRITE_TEXT,
		fileLabel,
		ESC,
		displayPos,
		mode,
		special,
	}
	return makeTextPacket(ccp, message)
}

func writeTextSpecial(fileLabel byte, displayPos byte, special byte, message string) []byte {
	ccp := []byte{
		CC_WRITE_TEXT,
		fileLabel,
		ESC,
		displayPos,
		MC_SPECIAL,
		special,
	}
	return makeTextPacket(ccp, message)
}

*/
/*
func (self *ReadWriter) WriteTextFile(label byte, msg string) (n int, err error) {
	p := self.SOT()
	p = append(p, WriteTextFile(label, msg)...)
	p = append(p, EOT)
	return self.writeText(label, p)
}

func (self *ReadWriter) WriteTextFileMode(label byte, mode byte, msg string) (n int, err error) {
	p := self.SOT()
	p = append(p, WriteTextFileMode(label, mode, msg)...)
	p = append(p, EOT)
	return self.writeText(label, p)
}
func (self *ReadWriter) WriteTextFileSpecial(label byte, special byte, msg string) (n int, err error) {
	p := self.SOT()
	p = append(p, WriteTextFileSpecial(label, special, msg)...)
	p = append(p, EOT)
	return self.writeText(label, p)
}
func (self *ReadWriter) WriteTextFilePosMode(label byte, displayPos byte, mode byte, msg string) (n int, err error) {
	p := self.SOT()
	p = append(p, WriteTextFilePosMode(label, displayPos, mode, msg)...)
	p = append(p, EOT)
	return self.writeText(label, p)
}
func (self *ReadWriter) WriteTextFilePosSpecial(label byte, displayPos byte, special byte, msg string) (n int, err error) {
	p := self.SOT()
	p = append(p, WriteTextFilePosSpecial(label, displayPos, special, msg)...)
	p = append(p, EOT)
	return self.writeText(label, p)
}
func (self *ReadWriter) TextMemoryConfig(label byte, size int) (n int, err error) {
	p := self.SOT()
	p = append(p, TextMemoryConfig(label, size)...)
	p = append(p, EOT)
	return self.rw.Write(p)
}

func WriteTextFileMode(label byte, mode byte, msg string) (p []byte) {
	p = []byte{
		CC_WRITE_TEXT,
		label,
		ESC,
		DP_MIDDLE_LINE,
		mode,
	}
	p = append(p, msg...)
	return
}

func WriteTextFilePosMode(label byte, displayPos byte, mode byte, msg string) (p []byte) {
	p = []byte{
		CC_WRITE_TEXT,
		label,
		ESC,
		displayPos,
		mode,
	}
	p = append(p, msg...)
	return
}

func WriteTextFileSpecial(label byte, special byte, msg string) (p []byte) {
	p = []byte{
		CC_WRITE_TEXT,
		label,
		ESC,
		DP_MIDDLE_LINE,
		MC_SPECIAL,
		special,
	}
	p = append(p, msg...)
	return
}

func WriteTextFilePosSpecial(label byte, displayPos byte, special byte, msg string) (p []byte) {
	p = []byte{
		CC_WRITE_TEXT,
		label,
		ESC,
		displayPos,
		MC_SPECIAL,
		special,
	}
	p = append(p, msg...)
	return
}

func WriteTextFile(label byte, msg string) (p []byte) {
	p = []byte{
		CC_WRITE_TEXT,
		label,
	}
	p = append(p, msg...)
	return
}
*/
