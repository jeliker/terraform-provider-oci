// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// Function A function resource defines the code (Docker image) and configuration for a specific function. Functions are defined in applications. Avoid entering confidential information.
type Function struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the function.
	Id *string `mandatory:"false" json:"id"`

	// The display name of the function. The display name is unique within the application containing the function.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the function.
	LifecycleState FunctionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCID of the application the function belongs to.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The OCID of the compartment that contains the function.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The qualified name of the Docker image to use in the function, including the image tag.
	// The image should be in the OCI Registry that is in the same region as the function itself.
	// Example: `phx.ocir.io/ten/functions/function:0.0.1`
	Image *string `mandatory:"false" json:"image"`

	// The image digest for the version of the image that will be pulled when invoking this function.
	// If no value is specified, the digest currently associated with the image in the OCI Registry will be used.
	// Example: `sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7`
	ImageDigest *string `mandatory:"false" json:"imageDigest"`

	// Maximum usable memory for the function (MiB).
	MemoryInMBs *int64 `mandatory:"false" json:"memoryInMBs"`

	// Function configuration. Overrides application configuration.
	// Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.
	// Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`
	// The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
	Config map[string]string `mandatory:"false" json:"config"`

	// Timeout for executions of the function. Value in seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached.
	InvokeEndpoint *string `mandatory:"false" json:"invokeEndpoint"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The time the function was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the function was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2018-09-12T22:47:12.613Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m Function) String() string {
	return common.PointerString(m)
}

// FunctionLifecycleStateEnum Enum with underlying type: string
type FunctionLifecycleStateEnum string

// Set of constants representing the allowable values for FunctionLifecycleStateEnum
const (
	FunctionLifecycleStateCreating FunctionLifecycleStateEnum = "CREATING"
	FunctionLifecycleStateActive   FunctionLifecycleStateEnum = "ACTIVE"
	FunctionLifecycleStateInactive FunctionLifecycleStateEnum = "INACTIVE"
	FunctionLifecycleStateUpdating FunctionLifecycleStateEnum = "UPDATING"
	FunctionLifecycleStateDeleting FunctionLifecycleStateEnum = "DELETING"
	FunctionLifecycleStateDeleted  FunctionLifecycleStateEnum = "DELETED"
	FunctionLifecycleStateFailed   FunctionLifecycleStateEnum = "FAILED"
)

var mappingFunctionLifecycleState = map[string]FunctionLifecycleStateEnum{
	"CREATING": FunctionLifecycleStateCreating,
	"ACTIVE":   FunctionLifecycleStateActive,
	"INACTIVE": FunctionLifecycleStateInactive,
	"UPDATING": FunctionLifecycleStateUpdating,
	"DELETING": FunctionLifecycleStateDeleting,
	"DELETED":  FunctionLifecycleStateDeleted,
	"FAILED":   FunctionLifecycleStateFailed,
}

// GetFunctionLifecycleStateEnumValues Enumerates the set of values for FunctionLifecycleStateEnum
func GetFunctionLifecycleStateEnumValues() []FunctionLifecycleStateEnum {
	values := make([]FunctionLifecycleStateEnum, 0)
	for _, v := range mappingFunctionLifecycleState {
		values = append(values, v)
	}
	return values
}
