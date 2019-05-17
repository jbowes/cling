package skip

import "testing"

func TestAllNils(t *testing.T) {
	if err := Wrap(nil, 0, "wrap"); err != nil {
		t.Error("nil mishandled")
	}

	if err := Wrapf(nil, 0, "wrap"); err != nil {
		t.Error("nil mishandled")
	}

	if err := Seal(nil, 0, "seal"); err != nil {
		t.Error("nil mishandled")
	}

	if err := Sealf(nil, 0, "seal"); err != nil {
		t.Error("nil mishandled")
	}
}
