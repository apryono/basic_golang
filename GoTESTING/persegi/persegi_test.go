package persegi_test

import (
	"materiGolang/GoTESTING/persegi"
	"testing"
)

func TestHitungLuasPersegi(t *testing.T) {

	testSisi := 5
	expectedLuas := 25

	t.Run("Test Normal Case", func(t *testing.T) {
		hasilLuas, _ := persegi.HitungLuasPersegi(testSisi)
		if hasilLuas != expectedLuas {
			t.Errorf("HitungLuasPersegi return wrong value got %v, expected %v", hasilLuas, expectedLuas)
		}
	})

	t.Run("Test Negative Case", func(t *testing.T) {
		testSisi := -1
		_, err := persegi.HitungLuasPersegi(testSisi)
		if err != persegi.ErrSisiTidakBolehMinus {
			t.Errorf("HitungLuasPersegi not returning error ErrSisiTidakBolehMinus, it return err : %v", err)
		}
	})

}
