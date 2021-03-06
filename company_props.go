/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CompanyProps struct {

	ApiKey string `json:"apiKey,omitempty"`

	CallbackUrl string `json:"callbackUrl,omitempty"`

	Company *Company `json:"company,omitempty"`

	Id int64 `json:"id,omitempty"`

	PaymentGatewayApiKey string `json:"paymentGatewayApiKey,omitempty"`

	PaymentGatewaySecurityKey string `json:"paymentGatewaySecurityKey,omitempty"`

	PaymentMethod string `json:"paymentMethod,omitempty"`

	SecureKey string `json:"secureKey,omitempty"`
}
