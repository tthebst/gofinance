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
func blackscholes(time_to_mat float64, spot float64, strike float64, risk_free float64, sig float64) (float64, error) {

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

func movingaverage(data []float64, days_to_avg int) ([]float64, error) {

	//return error invalid data
	if len(data) < 1 {
		return make([]float64, 1), errors.New("invalid time-data input")
	}
	// days_to_avg needs to be bigger than or 1
	if days_to_avg < 1 {
		return make([]float64, 1), errors.New("invalid days to average need to bigger then 1")
	}

	// create new array for movingaverage data, is smaller than data array because for first values can't calculates average
	mov_avg := make([]float64, len(data)-days_to_avg+1)

	// loop throung movering average loop and perform sliding window operation over data to calculate moving average
	// Time complexity O(n*d), n size of data array, d days to average
	for i := 0; i < len(mov_avg); i++ {
		for j := 0; j < days_to_avg; j++ {
			//position in data array is shifted by days_to_avg-1 in comparisan to move avg array
			mov_avg[i] += data[i+days_to_avg-1-j]
		}
		mov_avg[i] = mov_avg[i] / float64(days_to_avg)
	}
	return mov_avg, nil

}

//returns call price of an option
func Get_call_price(time_to_mat float64, spot float64, strike float64, risk_free float64, sig float64) (float64, error) {

	return blackscholes(time_to_mat, spot, strike, risk_free, sig)
}

//returns put price of an option
func Get_put_price(time_to_mat float64, spot float64, strike float64, risk_free float64, sig float64) (float64, error) {
	call_price, err := blackscholes(time_to_mat, spot, strike, risk_free, sig)

	//calcul
	put_price := strike*math.Exp(-risk_free*time_to_mat) - spot + call_price
	return put_price, err
}

func Get_movingaverage(data []float64, days_to_avg int) ([]float64, error) {
	return movingaverage(data, days_to_avg)
}
