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
// AddAlertRequestWarningInput struct for AddAlertRequestWarningInput
type AddAlertRequestWarningInput struct {
	Rules []SmartRule `json:"rules,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	// 1 is once a day, 2 is once a minute
	WarningFrequency int32 `json:"warningFrequency,omitempty"`
	// 1 is once, 0 is repeating
	WarningInterval int32 `json:"warningInterval,omitempty"`
}
