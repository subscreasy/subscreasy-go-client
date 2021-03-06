/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StartSubscriptionRequest struct {

	ApiKey string `json:"apiKey,omitempty"`

	CallbackUrl string `json:"callbackUrl,omitempty"`

	CouponCode string `json:"couponCode,omitempty"`

	Offer *SubscriptionPlan `json:"offer,omitempty"`

	PaymentCard *PaymentCard `json:"paymentCard,omitempty"`

	Subscriber *Subscriber `json:"subscriber,omitempty"`
}
