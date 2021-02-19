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
// PostOtocoStockOrderRequest struct for PostOtocoStockOrderRequest
type PostOtocoStockOrderRequest struct {
	Action OrderSide `json:"action,omitempty"`
	// limit price
	LmtPrice float32 `json:"lmtPrice,omitempty"`
	OrderType OrderType `json:"orderType,omitempty"`
	OutsideRegularTradingHour bool `json:"outsideRegularTradingHour,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	SerialId string `json:"serialId,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`
}
