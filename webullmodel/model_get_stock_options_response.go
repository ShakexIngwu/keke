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
// GetStockOptionsResponse struct for GetStockOptionsResponse
type GetStockOptionsResponse struct {
	Change string `json:"change,omitempty"`
	ChangeRatio string `json:"changeRatio,omitempty"`
	Close string `json:"close,omitempty"`
	DisExchangeCode string `json:"disExchangeCode,omitempty"`
	DisSymbol string `json:"disSymbol,omitempty"`
	ExpireDate string `json:"expireDate,omitempty"`
	ExpireDateList []GetStockOptionsResponseExpireDateList `json:"expireDateList,omitempty"`
	Name string `json:"name,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	UnSymbol string `json:"unSymbol,omitempty"`
}
