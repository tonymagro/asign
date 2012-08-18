package asign

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	tmpl := `{SOT}
{STX}{WriteText}{A}{Rotate}Here is some {Red}red {Green}green and text{RollUp}{Orange}Some orange text!{ETX}

{STX}
{WriteText}{B}
{RollDown}{Rainbow1}This text is rolling down
{Rotate}{Brown}Here is some rotating {DimRed}text
{ETX}`

	p, _ := Parse([]byte(tmpl))
	println(PacketString(p))

}
