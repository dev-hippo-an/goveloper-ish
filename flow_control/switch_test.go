package flow_control

import "testing"

func TestSwitchStatement(t *testing.T) {
	v := "hello"

	switch l := len(v); l {
	case 4:
		t.Log("Length is 4")
	case 5:
		t.Log("Length is 5") // length is 5
	case 6:
		t.Log("Length is 6")
	default:
		t.Log("this is default")
	}

	switch {
	case len(v) < 5:
		t.Log("Length is lower than 5")
	case len(v) >= 5:
		t.Log("Length is greater or equals than 5")
	}
}
