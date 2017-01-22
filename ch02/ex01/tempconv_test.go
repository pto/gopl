package tempconv

import "testing"

func TestConversions(t *testing.T) {
	testCFC := func(c Celsius) {
		if FToC(CToF(c)) != c {
			t.Error("C/F/C conversion failed on", c)
		}
	}
	testCKFC := func(c Celsius) {
		if FToC(KToF(CToK(c))) != c {
			t.Error("C/K/F/C conversion failed on", c)
		}
	}
	testCFKC := func(c Celsius) {
		if KToC(FToK(CToF(c))) != c {
			t.Error("C/F/K/C conversion failed on", c)
		}
	}
	testCKC := func(c Celsius) {
		if KToC(CToK(c)) != c {
			t.Error("C/K/C conversion failed on", c)
		}
	}

	for _, f := range []func(c Celsius){testCFC, testCKFC, testCFKC, testCKC} {
		f(AbsoluteZeroC)
		f(FreezingC)
		f(BoilingC)
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
