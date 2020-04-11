package internal

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

// blackscholes calculates option prices
// Arguments:
// float time_to_mat -> time to maturity
// float spot -> spot price of underlying asset
// float strike -> strike price
// float risk_free -> risk free rate
// float sig -> volatility of underlying asset
// Returns:
// float c -> call price
func Blackscholes(time_to_mat float64, spot float64, strike float64, risk_free float64, sig float64) (float64, error) {

	//check if any negativce values were provided
	if time_to_mat < 0 || spot < 0 || strike < 0 || risk_free < 0 || sig < 0 {
		return -1, errors.New("Provided negaitve number as arsgument to blackscholes")
	}
	//check if risk_free rate and volatility value percentage between 0,1
	if risk_free > 1.0 || sig > 1.0 {
		return -1, errors.New("percentage value not in the range between 0 and 1")
	}

	//define standart normal distribution
	norm_dist := distuv.Normal{
		Mu:    0,
		Sigma: 1,
	}
	// black scholes formulas SEE: https://en.wikipedia.org/wiki/Black%E2%80%93Scholes_model
	d1 := 1 / (sig * math.Sqrt(time_to_mat)) * (math.Log(spot/strike) + (risk_free+math.Pow(sig, 2)/2)*time_to_mat)
	d2 := d1 - sig*math.Sqrt(time_to_mat)
	pv := strike * math.Exp(-risk_free*time_to_mat)
	call_price := norm_dist.CDF(d1)*spot - norm_dist.CDF(d2)*pv

	return call_price, nil

}
