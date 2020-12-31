package azureadb2c

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/azureadb2c/mgmt/2019-01-01-preview/azureadb2c"

// AsyncOperationStatus the async operation status.
type AsyncOperationStatus struct {
	autorest.Response `json:"-"`
	// SubscriptionID - Subscription ID that the resource belongs to.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ID - The GET resource path for the operation.
	ID *string `json:"id,omitempty"`
	// Name - The operation ID.
	Name *string `json:"name,omitempty"`
	// Status - The status of the operation. Possible values include: 'Succeeded', 'Pending', 'Failed'
	Status StatusType `json:"status,omitempty"`
	// StartTime - Start time of the async operation.
	StartTime *string `json:"startTime,omitempty"`
	// EndTime - End time of the async operation.
	EndTime *string `json:"endTime,omitempty"`
	// B2CTenantResourceProperties - The Azure AD B2C tenant resource properties
	*B2CTenantResourceProperties `json:"properties,omitempty"`
	// Error - Error response if async operation failed.
	Error *AsyncOperationStatusError `json:"error,omitempty"`
}

// MarshalJSON is the custom marshaler for AsyncOperationStatus.
func (aos AsyncOperationStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aos.SubscriptionID != nil {
		objectMap["subscriptionId"] = aos.SubscriptionID
	}
	if aos.ID != nil {
		objectMap["id"] = aos.ID
	}
	if aos.Name != nil {
		objectMap["name"] = aos.Name
	}
	if aos.Status != "" {
		objectMap["status"] = aos.Status
	}
	if aos.StartTime != nil {
		objectMap["startTime"] = aos.StartTime
	}
	if aos.EndTime != nil {
		objectMap["endTime"] = aos.EndTime
	}
	if aos.B2CTenantResourceProperties != nil {
		objectMap["properties"] = aos.B2CTenantResourceProperties
	}
	if aos.Error != nil {
		objectMap["error"] = aos.Error
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AsyncOperationStatus struct.
func (aos *AsyncOperationStatus) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "subscriptionId":
			if v != nil {
				var subscriptionID string
				err = json.Unmarshal(*v, &subscriptionID)
				if err != nil {
					return err
				}
				aos.SubscriptionID = &subscriptionID
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				aos.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				aos.Name = &name
			}
		case "status":
			if v != nil {
				var status StatusType
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				aos.Status = status
			}
		case "startTime":
			if v != nil {
				var startTime string
				err = json.Unmarshal(*v, &startTime)
				if err != nil {
					return err
				}
				aos.StartTime = &startTime
			}
		case "endTime":
			if v != nil {
				var endTime string
				err = json.Unmarshal(*v, &endTime)
				if err != nil {
					return err
				}
				aos.EndTime = &endTime
			}
		case "properties":
			if v != nil {
				var b2CTenantResourceProperties B2CTenantResourceProperties
				err = json.Unmarshal(*v, &b2CTenantResourceProperties)
				if err != nil {
					return err
				}
				aos.B2CTenantResourceProperties = &b2CTenantResourceProperties
			}
		case "error":
			if v != nil {
				var errorVar AsyncOperationStatusError
				err = json.Unmarshal(*v, &errorVar)
				if err != nil {
					return err
				}
				aos.Error = &errorVar
			}
		}
	}

	return nil
}

// AsyncOperationStatusError error response if async operation failed.
type AsyncOperationStatusError struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message.
	Message *string `json:"message,omitempty"`
}

// B2CResourceSKU SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at
// [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
type B2CResourceSKU struct {
	// Name - The name of the SKU for the tenant. Possible values include: 'Standard', 'PremiumP1', 'PremiumP2'
	Name B2CResourceSKUName `json:"name,omitempty"`
	// Tier - The tier of the tenant. Possible values include: 'A0'
	Tier B2CResourceSKUTier `json:"tier,omitempty"`
}

// B2CTenantResource ...
type B2CTenantResource struct {
	autorest.Response `json:"-"`
	// Type - READ-ONLY; The type of the B2C tenant resource. Possible values include: 'MicrosoftAzureActiveDirectoryb2cDirectories'
	Type TypeValue       `json:"type,omitempty"`
	Sku  *B2CResourceSKU `json:"sku,omitempty"`
	// B2CTenantResourceProperties - The B2C tenant resource properties
	*B2CTenantResourceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; An identifier that represents the B2C tenant resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the B2C tenant resource.
	Name *string `json:"name,omitempty"`
	// Location - The location in which the resource is hosted and data resides. Refer to [this documentation](https://aka.ms/B2CDataResidency) to see valid data residency locations. Please choose one of 'United States', 'Europe', and 'Asia Pacific'.
	Location *string `json:"location,omitempty"`
	// Tags - Resource Tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for B2CTenantResource.
func (btr B2CTenantResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if btr.Sku != nil {
		objectMap["sku"] = btr.Sku
	}
	if btr.B2CTenantResourceProperties != nil {
		objectMap["properties"] = btr.B2CTenantResourceProperties
	}
	if btr.Location != nil {
		objectMap["location"] = btr.Location
	}
	if btr.Tags != nil {
		objectMap["tags"] = btr.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for B2CTenantResource struct.
func (btr *B2CTenantResource) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "type":
			if v != nil {
				var typeVar TypeValue
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				btr.Type = typeVar
			}
		case "sku":
			if v != nil {
				var sku B2CResourceSKU
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				btr.Sku = &sku
			}
		case "properties":
			if v != nil {
				var b2CTenantResourceProperties B2CTenantResourceProperties
				err = json.Unmarshal(*v, &b2CTenantResourceProperties)
				if err != nil {
					return err
				}
				btr.B2CTenantResourceProperties = &b2CTenantResourceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				btr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				btr.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				btr.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				btr.Tags = tags
			}
		}
	}

	return nil
}

// B2CTenantResourceList the collection of Azure AD B2C tenant resources
type B2CTenantResourceList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of guest usages resources
	Value *[]B2CTenantResource `json:"value,omitempty"`
}

// B2CTenantResourceProperties properties of the Azure AD B2C tenant Azure resource.
type B2CTenantResourceProperties struct {
	// BillingConfig - The billing configuration for the tenant.
	BillingConfig *B2CTenantResourcePropertiesBillingConfig `json:"billingConfig,omitempty"`
	// TenantID - An identifier of the B2C tenant.
	TenantID *string `json:"tenantId,omitempty"`
}

// B2CTenantResourcePropertiesBillingConfig the billing configuration for the tenant.
type B2CTenantResourcePropertiesBillingConfig struct {
	// BillingType - The type of billing. Will be MAU for all new customers. If 'Auths', it can be updated to 'MAU'. Cannot be changed if value is 'MAU'. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cbilling). Possible values include: 'MAU', 'Auths'
	BillingType BillingType `json:"billingType,omitempty"`
	// EffectiveStartDateUtc - READ-ONLY; The data from which the billing type took effect
	EffectiveStartDateUtc *string `json:"effectiveStartDateUtc,omitempty"`
}

// MarshalJSON is the custom marshaler for B2CTenantResourcePropertiesBillingConfig.
func (btrpC B2CTenantResourcePropertiesBillingConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if btrpC.BillingType != "" {
		objectMap["billingType"] = btrpC.BillingType
	}
	return json.Marshal(objectMap)
}

// B2CTenantsCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type B2CTenantsCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(B2CTenantsClient) (B2CTenantResource, error)
}

// B2CTenantsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type B2CTenantsDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(B2CTenantsClient) (autorest.Response, error)
}

// B2CTenantUpdateRequest the request body to update the Azure AD B2C tenant resource.
type B2CTenantUpdateRequest struct {
	Sku *B2CResourceSKU `json:"sku,omitempty"`
	// B2CTenantResourceProperties - The Azure AD B2C tenant resource properties.
	*B2CTenantResourceProperties `json:"properties,omitempty"`
	// Tags - Resource Tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for B2CTenantUpdateRequest.
func (btur B2CTenantUpdateRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if btur.Sku != nil {
		objectMap["sku"] = btur.Sku
	}
	if btur.B2CTenantResourceProperties != nil {
		objectMap["properties"] = btur.B2CTenantResourceProperties
	}
	if btur.Tags != nil {
		objectMap["tags"] = btur.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for B2CTenantUpdateRequest struct.
func (btur *B2CTenantUpdateRequest) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku B2CResourceSKU
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				btur.Sku = &sku
			}
		case "properties":
			if v != nil {
				var b2CTenantResourceProperties B2CTenantResourceProperties
				err = json.Unmarshal(*v, &b2CTenantResourceProperties)
				if err != nil {
					return err
				}
				btur.B2CTenantResourceProperties = &b2CTenantResourceProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				btur.Tags = tags
			}
		}
	}

	return nil
}

// CheckNameAvailabilityRequestBody the information required to check the availability of the name for the
// tenant.
type CheckNameAvailabilityRequestBody struct {
	// Name - The domain name to check for availability.
	Name        *string `json:"name,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
}

// CloudError an error response for a resource management request.
type CloudError struct {
	Error *ErrorResponse `json:"error,omitempty"`
}

// CreateTenantProperties these properties are used to create the Azure AD B2C tenant. These properties are
// not part of the Azure resource.
type CreateTenantProperties struct {
	// DisplayName - The display name of the B2C tenant.
	DisplayName *string `json:"displayName,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
}

// CreateTenantRequestBody the information needed to create the Azure AD B2C tenant and corresponding Azure
// resource, which is used for billing purposes.
type CreateTenantRequestBody struct {
	// Location - The location in which the resource is hosted and data resides. Refer to [this documentation](https://aka.ms/B2CDataResidency) to see valid data residency locations. Please choose one of 'United States', 'Europe', and 'Asia Pacific'.
	Location   *string                            `json:"location,omitempty"`
	Properties *CreateTenantRequestBodyProperties `json:"properties,omitempty"`
	Sku        *B2CResourceSKU                    `json:"sku,omitempty"`
	// Tags - Resource Tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for CreateTenantRequestBody.
func (ctrb CreateTenantRequestBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ctrb.Location != nil {
		objectMap["location"] = ctrb.Location
	}
	if ctrb.Properties != nil {
		objectMap["properties"] = ctrb.Properties
	}
	if ctrb.Sku != nil {
		objectMap["sku"] = ctrb.Sku
	}
	if ctrb.Tags != nil {
		objectMap["tags"] = ctrb.Tags
	}
	return json.Marshal(objectMap)
}

// CreateTenantRequestBodyProperties ...
type CreateTenantRequestBodyProperties struct {
	*CreateTenantProperties `json:"createTenantProperties,omitempty"`
}

// MarshalJSON is the custom marshaler for CreateTenantRequestBodyProperties.
func (ctrb CreateTenantRequestBodyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ctrb.CreateTenantProperties != nil {
		objectMap["createTenantProperties"] = ctrb.CreateTenantProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CreateTenantRequestBodyProperties struct.
func (ctrb *CreateTenantRequestBodyProperties) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "createTenantProperties":
			if v != nil {
				var createTenantProperties CreateTenantProperties
				err = json.Unmarshal(*v, &createTenantProperties)
				if err != nil {
					return err
				}
				ctrb.CreateTenantProperties = &createTenantProperties
			}
		}
	}

	return nil
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.)
type ErrorResponse struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorResponse `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// NameAvailabilityResponse response of the CheckNameAvailability operation.
type NameAvailabilityResponse struct {
	autorest.Response `json:"-"`
	// Message - Description of the reason if name is not available.
	Message *string `json:"message,omitempty"`
	// NameAvailable - True if the name is available and can be used to create a new tenant. Otherwise false.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - Possible values include: 'AlreadyExists', 'Invalid'
	Reason NameAvailabilityReasonType `json:"reason,omitempty"`
}

// Operation microsoft.AzureActiveDirectory REST API operation.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	return json.Marshal(objectMap)
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft.AzureActiveDirectory.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed: GuestUsages, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Friendly name of the operation
	Description *string `json:"description,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Description != nil {
		objectMap["description"] = o.Description
	}
	return json.Marshal(objectMap)
}

// OperationListResult result of listing operations for the resourceProvider
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of cpim service operations supported by the Microsoft.AzureActiveDirectory resource provider.
	Value *[]Operation `json:"value,omitempty"`
}