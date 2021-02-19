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
// OrderType Order type, like: Limit (LMT), Stop (STP), and Stop Limit (STP LMT).
type OrderType string

// List of OrderType
const (
	LMT OrderType = "LMT"
	STP OrderType = "STP"
	STP_LMT OrderType = "STP LMT"
	MKT OrderType = "MKT"
)
