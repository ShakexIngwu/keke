/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webullmodel
// GetStockQuoteResponse struct for GetStockQuoteResponse
type GetStockQuoteResponse struct {
	AvgVol10D string `json:"avgVol10D,omitempty"`
	AvgVol3M string `json:"avgVol3M,omitempty"`
	Beta3Y string `json:"beta3Y,omitempty"`
	Bps string `json:"bps,omitempty"`
	Change string `json:"change,omitempty"`
	ChangeRatio string `json:"changeRatio,omitempty"`
	Close string `json:"close,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	DerivativeSupport int32 `json:"derivativeSupport,omitempty"`
	Dividend string `json:"dividend,omitempty"`
	Eps string `json:"eps,omitempty"`
	EpsTtm string `json:"epsTtm,omitempty"`
	FiftyTwoWkHigh string `json:"fiftyTwoWkHigh,omitempty"`
	FiftyTwoWkLow string `json:"fiftyTwoWkLow,omitempty"`
	High string `json:"high,omitempty"`
	InceptionDate string `json:"inceptionDate,omitempty"`
	LatestDividendDate string `json:"latestDividendDate,omitempty"`
	LimitDown string `json:"limitDown,omitempty"`
	LimitUp string `json:"limitUp,omitempty"`
	LotSize string `json:"lotSize,omitempty"`
	Low string `json:"low,omitempty"`
	NetAsset string `json:"netAsset,omitempty"`
	NetExpenseRatio string `json:"netExpenseRatio,omitempty"`
	NetValue string `json:"netValue,omitempty"`
	NtvSize int32 `json:"ntvSize,omitempty"`
	Open string `json:"open,omitempty"`
	PChRatio string `json:"pChRatio,omitempty"`
	PChange string `json:"pChange,omitempty"`
	PPrice string `json:"pPrice,omitempty"`
	Pe string `json:"pe,omitempty"`
	PeTtm string `json:"peTtm,omitempty"`
	PreClose string `json:"preClose,omitempty"`
	Rating int32 `json:"rating,omitempty"`
	ReturnThisYear string `json:"returnThisYear,omitempty"`
	Status string `json:"status,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	TimeZone string `json:"timeZone,omitempty"`
	TotalAsset string `json:"totalAsset,omitempty"`
	TradeTime string `json:"tradeTime,omitempty"`
	TzName string `json:"tzName,omitempty"`
	VibrateRatio string `json:"vibrateRatio,omitempty"`
	Volume string `json:"volume,omitempty"`
	Yield string `json:"yield,omitempty"`
	Yield1Y string `json:"yield1Y,omitempty"`
}
