package internal

import (
	"math"
	"testing"

	"github.com/gofinance/internal"
)

func TestBlackScholesCall(t *testing.T) {
	got, err := internal.Get_call_price(1.0, 300.0, 250.0, 0.03, 0.15)
	if err != nil {
		t.Errorf("Unexpected Error when calling blackscholes function:  %v", err)
	}
	expected := 58.82
	diff := got - expected
	if math.Abs(diff) < 1e-9 {
		t.Errorf("Flalse blackschole calculation. Diffrence: %v", diff)
	}
}
func TestBlackScholesPut(t *testing.T) {
	got, err := internal.Get_put_price(1.0, 300.0, 250.0, 0.03, 0.15)
	if err != nil {
		t.Errorf("Unexpected Error when calling blackscholes function:  %v", err)
	}
	expected := 1.431
	diff := got - expected
	if math.Abs(diff) < 1e-5 {
		t.Errorf("Flalse blackschole calculation. Diffrence: %v", diff)
	}
}
func TestBlackScholesCallError(t *testing.T) {
	_, err := internal.Get_call_price(1.0, 300.0, -250.0, 0.03, 0.15)
	if err == nil {
		t.Errorf("Should raise error")
	}

}

func TestBlackScholesPutError(t *testing.T) {
	_, err := internal.Get_put_price(1.0, 300.0, -250.0, 0.03, 0.15)
	if err == nil {
		t.Errorf("Should raise error")
	}

}
