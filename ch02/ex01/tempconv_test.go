package tempconv

import "testing"

func TestConversions(t *testing.T) {
	if FToC(CToF(AbsoluteZeroC)) != AbsoluteZeroC {
		t.Error("F/C Conversion failed on absolute zero")
	}
	if FToC(CToF(FreezingC)) != FreezingC {
		t.Error("F/C Conversion failed on freezing point")
	}
	if FToC(CToF(BoilingC)) != BoilingC {
		t.Error("F/C Conversion failed on boiling point")
	}

	if FToK(KToF(CToK(AbsoluteZeroC))) != CToK(AbsoluteZeroC) {
		t.Error("F/C/K Conversion failed on absolute zero")
	}
	if FToK(KToF(CToK(FreezingC))) != CToK(FreezingC) {
		t.Error("F/C/K Conversion failed on freezing point")
	}
	if FToK(KToF(CToK(BoilingC))) != CToK(BoilingC) {
		t.Error("F/C/K Conversion failed on boiling point")
	}
}
