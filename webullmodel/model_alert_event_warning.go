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
// AlertEventWarning struct for AlertEventWarning
type AlertEventWarning struct {
	Remove bool `json:"remove,omitempty"`
	Rules []SmartRule `json:"rules,omitempty"`
}
