package internal

import (
	"math"
	"reflect"
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
func TestMovingaverage1(t *testing.T) {
	data := []float64{2, 4, 6, 8, 12, 14, 16, 18, 20}
	avg, _ := internal.Get_movingaverage(data, 2)
	expected := []float64{3, 5, 7, 10, 13, 15, 17, 19}
	if !reflect.DeepEqual(expected, avg) {
		t.Errorf("Failed to calculate correct moving average")
	}

}

func TestMovingaverage2(t *testing.T) {
	data := []float64{2, 4, 6, 8, 12, 14, 16, 18, 20}
	avg, _ := internal.Get_movingaverage(data, 1)
	expected := []float64{2, 4, 6, 8, 12, 14, 16, 18, 20}
	if !reflect.DeepEqual(expected, avg) {
		t.Errorf("Failed to calculate correct moving average")
	}

}

func TestMovingaverageError(t *testing.T) {
	data := []float64{2, 4, 6, 8, 12, 14, 16, 18, 20}
	_, err := internal.Get_movingaverage(data, -1)
	if err == nil {
		t.Errorf("Should raise error")
	}

}
