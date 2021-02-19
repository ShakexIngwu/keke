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
// GetAlertsResponseData struct for GetAlertsResponseData
type GetAlertsResponseData struct {
	DisExchangeCode string `json:"disExchangeCode,omitempty"`
	DisSymbol string `json:"disSymbol,omitempty"`
	EventWarning GetAlertsResponseEventWarning `json:"eventWarning,omitempty"`
	ExchangeCode string `json:"exchangeCode,omitempty"`
	ExchangeTrade bool `json:"exchangeTrade,omitempty"`
	RegionId int32 `json:"regionId,omitempty"`
	ShowCode string `json:"showCode,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	TickerName string `json:"tickerName,omitempty"`
	TickerSymbol string `json:"tickerSymbol,omitempty"`
	TickerType int32 `json:"tickerType,omitempty"`
	TickerWarning GetAlertsResponseTickerWarning `json:"tickerWarning,omitempty"`
	TinyName string `json:"tinyName,omitempty"`
}
