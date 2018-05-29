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

type ManagedUserVm struct {

	Activated bool `json:"activated,omitempty"`

	Authorities []string `json:"authorities,omitempty"`

	Company *Company `json:"company,omitempty"`

	CreatedBy string `json:"createdBy,omitempty"`

	CreatedDate time.Time `json:"createdDate,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Id int64 `json:"id,omitempty"`

	ImageUrl string `json:"imageUrl,omitempty"`

	LangKey string `json:"langKey,omitempty"`

	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Login string `json:"login"`

	Password string `json:"password,omitempty"`

	ResetDate string `json:"resetDate,omitempty"`
}