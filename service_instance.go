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

type ServiceInstance struct {

	Capacity float32 `json:"capacity,omitempty"`

	CurrentUsage float32 `json:"currentUsage,omitempty"`

	EndDate time.Time `json:"endDate,omitempty"`

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	NumberOfUnits int32 `json:"numberOfUnits,omitempty"`

	Offer *Offer `json:"offer,omitempty"`

	OverUsage float32 `json:"overUsage,omitempty"`

	OverUsageQuota float32 `json:"overUsageQuota,omitempty"`

	QuotaOrigin string `json:"quotaOrigin,omitempty"`

	ServiceOffering *ServiceOffering `json:"serviceOffering,omitempty"`

	ServiceType string `json:"serviceType,omitempty"`

	StartDate time.Time `json:"startDate,omitempty"`

	Status string `json:"status,omitempty"`

	SubscriberId string `json:"subscriberId,omitempty"`

	Subscription *Subsription `json:"subscription,omitempty"`

	Type_ string `json:"type,omitempty"`

	Version int64 `json:"version,omitempty"`
}