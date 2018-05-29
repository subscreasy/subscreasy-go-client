/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Invoice struct {

	Amount float32 `json:"amount,omitempty"`

	BillingMonth int32 `json:"billingMonth,omitempty"`

	BillingYear int32 `json:"billingYear,omitempty"`

	Id int64 `json:"id,omitempty"`

	PeriodEnd time.Time `json:"periodEnd,omitempty"`

	PeriodStart time.Time `json:"periodStart,omitempty"`

	Status string `json:"status,omitempty"`

	SubscriberSecureId string `json:"subscriberSecureId,omitempty"`
}