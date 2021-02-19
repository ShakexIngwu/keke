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
// PostRefreshTokenResponse struct for PostRefreshTokenResponse
type PostRefreshTokenResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	TokenExpireTime string `json:"tokenExpireTime,omitempty"`
}
