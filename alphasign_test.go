package asign

/*
import (
	"bytes"
	"testing"
)

func TestWriteTextFile(t *testing.T) {
	m := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, // auto baud
		0x01,     // SOH
		'Z',      // SA_BROADCAST
		'0', '0', // TC_ALLSIGNS
		0x02,               // STX
		'E',                // Special Function
		'$',                // Memory set
		'B',                // Label
		'A',                // Text File
		'L',                // Locked
		'0', '0', '0', 'F', // Size of text
		'F', 'F', // Time
		'0', '0',
		0x04, // EOT
	}

	p := append(m,
		0x00, 0x00, 0x00, 0x00, 0x00, // auto baud
		0x01,     // SOH
		'Z',      // SA_BROADCAST
		'0', '0', // TC_ALLSIGNS
		0x02, // STX
		'A',  // CC_WRITE_TEXT
		'B',  // label
	)
	p = append(p, "Hello World!"...)
	p = append(p, 0x03) // ETX
	p = append(p, 0x04) // EOT

	var b bytes.Buffer
	sign := NewReadWriter(&b, SA_BROADCAST, TC_ALLSIGNS)
	defer sign.Close()

	sign.WriteCmd("{{Text .B}}Hello World!{{.ETX}}")
	if bytes.Compare(p, b.Bytes()) != 0 {
		t.Log(PacketString(p))
		t.Log(PacketString(b.Bytes()))
		t.FailNow()
	}
}
*/

/*
func TestWriteFile(t *testing.T) {
	var label byte = 'B'
	msg := "Hello World!"
	p := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x01,
		'Z',
		'0', '0',
		0x02,
		'A',
		label,
	}
	p = append(p, msg...)
	p = append(p, 0x04)

	var b bytes.Buffer
	sign := NewReadWriter(&b, SA_BROADCAST, TC_ALLSIGNS)
	defer sign.Close()

	sign.WriteTextFile(label, msg)
	if bytes.Compare(p, b.Bytes()) != 0 {
		t.Log(PacketString(p))
		t.Log(PacketString(b.Bytes()))
		t.FailNow()
	}
}

func TestWriteFileMode(t *testing.T) {
	var label byte = 'B'
	var mode byte = 'a' // Rotate
	msg := "Hello World!"
	p := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x01,
		'Z',
		'0', '0',
		0x02,
		'A',
		label,
		ESC,
		0x20, // Middle Line
		mode,
	}
	p = append(p, msg...)
	p = append(p, 0x04)

	var b bytes.Buffer
	sign := NewReadWriter(&b, SA_BROADCAST, TC_ALLSIGNS)
	defer sign.Close()

	sign.WriteTextFileMode(label, MC_ROTATE, msg)
	if bytes.Compare(p, b.Bytes()) != 0 {
		t.Log(PacketString(p))
		t.Log(PacketString(b.Bytes()))
		t.FailNow()
	}
}

func TestWriteFileSpecial(t *testing.T) {
	var label byte = 'B'
	var special byte = '4' // SWITCH
	msg := "Hello World!"
	p := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x01,
		'Z',
		'0', '0',
		0x02,
		'A',
		label,
		ESC,
		0x20,
		'n',
		special,
	}
	p = append(p, msg...)
	p = append(p, 0x04)

	var b bytes.Buffer
	sign := NewReadWriter(&b, SA_BROADCAST, TC_ALLSIGNS)
	defer sign.Close()

	sign.WriteTextFileSpecial(label, MC_S_SWITCH, msg)
	if bytes.Compare(p, b.Bytes()) != 0 {
		t.Log(PacketString(p))
		t.Log(PacketString(b.Bytes()))
		t.FailNow()
	}
}

func TestWriteTextCmd(t *testing.T) {
	var label byte = 'B'
	var special byte = '4' // SWITCH
	msg := "Hello World!"
	p := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x01,
		'Z',
		'0', '0',
		0x02,
		'A',
		label,
		ESC,
		0x20,
		'n',
		special,
	}
	p = append(p, msg...)
	p = append(p, 0x04)

	var b bytes.Buffer
	sign := NewReadWriter(&b, SA_BROADCAST, TC_ALLSIGNS)
	defer sign.Close()

	sign.WriteTextFileCmd(label, msg)
	if bytes.Compare(p, b.Bytes()) != 0 {
		t.Log(PacketString(p))
		t.Log(PacketString(b.Bytes()))
		t.FailNow()
	}
}
*/
