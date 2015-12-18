package tempconv

import "testing"

func TestConversions(t *testing.T) {
	if FToC(CToF(AbsoluteZeroC)) != AbsoluteZeroC {
		t.Error("C/F/C Conversion failed on absolute zero")
	}
	if FToC(CToF(FreezingC)) != FreezingC {
		t.Error("C/F/C Conversion failed on freezing point")
	}
	if FToC(CToF(BoilingC)) != BoilingC {
		t.Error("C/F/C Conversion failed on boiling point")
	}

	if FToC(KToF(CToK(AbsoluteZeroC))) != AbsoluteZeroC {
		t.Error("C/K/F/C Conversion failed on absolute zero")
	}
	if FToC(KToF(CToK(FreezingC))) != FreezingC {
		t.Error("C/K/F/C Conversion failed on freezing point")
	}
	if FToC(KToF(CToK(BoilingC))) != BoilingC {
		t.Error("C/K/F/C Conversion failed on boiling point")
	}

	if KToC(CToK(AbsoluteZeroC)) != AbsoluteZeroC {
		t.Error("C/K/C Conversion failed on absolute zero")
	}
	if KToC(CToK(FreezingC)) != FreezingC {
		t.Error("C/K/C Conversion failed on freezing point")
	}
	if KToC(CToK(FreezingC)) != FreezingC {
		t.Error("C/K/C Conversion failed on freezing point")
	}

	if FToC(32) != 0 {
		t.Error("F/C Conversion on 32 failed")
	}
	if float64(CToK(0)) != -float64(AbsoluteZeroC) {
		t.Error("C/K Conversion on 0 failed")
	}

	if FreezingC.String() != "0°C" {
		t.Error("C String() failed")
	}
	if Kelvin(0).String() != "0°K" {
		t.Error("K String() failed")
	}
	if Fahrenheit(0).String() != "0°F" {
		t.Error("F String() failed")
	}
}
