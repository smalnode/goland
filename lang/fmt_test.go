package lang

import (
	"strconv"
	"testing"
)

func TestFmtFloat(t *testing.T) {
	fs := "1.234"
	f, err := strconv.ParseFloat(fs, 32)
	if err != nil {
		t.Errorf("convert %s to float failed %s", fs, err)
	}
	t.Log(fs)
	t.Logf("%f", f)
	t.Logf("%.2f", f)
}
