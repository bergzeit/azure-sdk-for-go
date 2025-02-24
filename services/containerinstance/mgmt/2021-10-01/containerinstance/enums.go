package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AutoGeneratedDomainNameLabelScope enumerates the values for auto generated domain name label scope.
type AutoGeneratedDomainNameLabelScope string

const (
	// AutoGeneratedDomainNameLabelScopeNoreuse ...
	AutoGeneratedDomainNameLabelScopeNoreuse AutoGeneratedDomainNameLabelScope = "Noreuse"
	// AutoGeneratedDomainNameLabelScopeResourceGroupReuse ...
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	// AutoGeneratedDomainNameLabelScopeSubscriptionReuse ...
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	// AutoGeneratedDomainNameLabelScopeTenantReuse ...
	AutoGeneratedDomainNameLabelScopeTenantReuse AutoGeneratedDomainNameLabelScope = "TenantReuse"
	// AutoGeneratedDomainNameLabelScopeUnsecure ...
	AutoGeneratedDomainNameLabelScopeUnsecure AutoGeneratedDomainNameLabelScope = "Unsecure"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns an array of possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{AutoGeneratedDomainNameLabelScopeNoreuse, AutoGeneratedDomainNameLabelScopeResourceGroupReuse, AutoGeneratedDomainNameLabelScopeSubscriptionReuse, AutoGeneratedDomainNameLabelScopeTenantReuse, AutoGeneratedDomainNameLabelScopeUnsecure}
}

// ContainerGroupIPAddressType enumerates the values for container group ip address type.
type ContainerGroupIPAddressType string

const (
	// ContainerGroupIPAddressTypePrivate ...
	ContainerGroupIPAddressTypePrivate ContainerGroupIPAddressType = "Private"
	// ContainerGroupIPAddressTypePublic ...
	ContainerGroupIPAddressTypePublic ContainerGroupIPAddressType = "Public"
)

// PossibleContainerGroupIPAddressTypeValues returns an array of possible values for the ContainerGroupIPAddressType const type.
func PossibleContainerGroupIPAddressTypeValues() []ContainerGroupIPAddressType {
	return []ContainerGroupIPAddressType{ContainerGroupIPAddressTypePrivate, ContainerGroupIPAddressTypePublic}
}

// ContainerGroupNetworkProtocol enumerates the values for container group network protocol.
type ContainerGroupNetworkProtocol string

const (
	// ContainerGroupNetworkProtocolTCP ...
	ContainerGroupNetworkProtocolTCP ContainerGroupNetworkProtocol = "TCP"
	// ContainerGroupNetworkProtocolUDP ...
	ContainerGroupNetworkProtocolUDP ContainerGroupNetworkProtocol = "UDP"
)

// PossibleContainerGroupNetworkProtocolValues returns an array of possible values for the ContainerGroupNetworkProtocol const type.
func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return []ContainerGroupNetworkProtocol{ContainerGroupNetworkProtocolTCP, ContainerGroupNetworkProtocolUDP}
}

// ContainerGroupRestartPolicy enumerates the values for container group restart policy.
type ContainerGroupRestartPolicy string

const (
	// ContainerGroupRestartPolicyAlways ...
	ContainerGroupRestartPolicyAlways ContainerGroupRestartPolicy = "Always"
	// ContainerGroupRestartPolicyNever ...
	ContainerGroupRestartPolicyNever ContainerGroupRestartPolicy = "Never"
	// ContainerGroupRestartPolicyOnFailure ...
	ContainerGroupRestartPolicyOnFailure ContainerGroupRestartPolicy = "OnFailure"
)

// PossibleContainerGroupRestartPolicyValues returns an array of possible values for the ContainerGroupRestartPolicy const type.
func PossibleContainerGroupRestartPolicyValues() []ContainerGroupRestartPolicy {
	return []ContainerGroupRestartPolicy{ContainerGroupRestartPolicyAlways, ContainerGroupRestartPolicyNever, ContainerGroupRestartPolicyOnFailure}
}

// ContainerGroupSku enumerates the values for container group sku.
type ContainerGroupSku string

const (
	// ContainerGroupSkuDedicated ...
	ContainerGroupSkuDedicated ContainerGroupSku = "Dedicated"
	// ContainerGroupSkuStandard ...
	ContainerGroupSkuStandard ContainerGroupSku = "Standard"
)

// PossibleContainerGroupSkuValues returns an array of possible values for the ContainerGroupSku const type.
func PossibleContainerGroupSkuValues() []ContainerGroupSku {
	return []ContainerGroupSku{ContainerGroupSkuDedicated, ContainerGroupSkuStandard}
}

// ContainerNetworkProtocol enumerates the values for container network protocol.
type ContainerNetworkProtocol string

const (
	// ContainerNetworkProtocolTCP ...
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = "TCP"
	// ContainerNetworkProtocolUDP ...
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = "UDP"
)

// PossibleContainerNetworkProtocolValues returns an array of possible values for the ContainerNetworkProtocol const type.
func PossibleContainerNetworkProtocolValues() []ContainerNetworkProtocol {
	return []ContainerNetworkProtocol{ContainerNetworkProtocolTCP, ContainerNetworkProtocolUDP}
}

// GpuSku enumerates the values for gpu sku.
type GpuSku string

const (
	// GpuSkuK80 ...
	GpuSkuK80 GpuSku = "K80"
	// GpuSkuP100 ...
	GpuSkuP100 GpuSku = "P100"
	// GpuSkuV100 ...
	GpuSkuV100 GpuSku = "V100"
)

// PossibleGpuSkuValues returns an array of possible values for the GpuSku const type.
func PossibleGpuSkuValues() []GpuSku {
	return []GpuSku{GpuSkuK80, GpuSkuP100, GpuSkuV100}
}

// LogAnalyticsLogType enumerates the values for log analytics log type.
type LogAnalyticsLogType string

const (
	// LogAnalyticsLogTypeContainerInsights ...
	LogAnalyticsLogTypeContainerInsights LogAnalyticsLogType = "ContainerInsights"
	// LogAnalyticsLogTypeContainerInstanceLogs ...
	LogAnalyticsLogTypeContainerInstanceLogs LogAnalyticsLogType = "ContainerInstanceLogs"
)

// PossibleLogAnalyticsLogTypeValues returns an array of possible values for the LogAnalyticsLogType const type.
func PossibleLogAnalyticsLogTypeValues() []LogAnalyticsLogType {
	return []LogAnalyticsLogType{LogAnalyticsLogTypeContainerInsights, LogAnalyticsLogTypeContainerInstanceLogs}
}

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// OperatingSystemTypesLinux ...
	OperatingSystemTypesLinux OperatingSystemTypes = "Linux"
	// OperatingSystemTypesWindows ...
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns an array of possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{OperatingSystemTypesLinux, OperatingSystemTypesWindows}
}

// OperationsOrigin enumerates the values for operations origin.
type OperationsOrigin string

const (
	// OperationsOriginSystem ...
	OperationsOriginSystem OperationsOrigin = "System"
	// OperationsOriginUser ...
	OperationsOriginUser OperationsOrigin = "User"
)

// PossibleOperationsOriginValues returns an array of possible values for the OperationsOrigin const type.
func PossibleOperationsOriginValues() []OperationsOrigin {
	return []OperationsOrigin{OperationsOriginSystem, OperationsOriginUser}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// Scheme enumerates the values for scheme.
type Scheme string

const (
	// SchemeHTTP ...
	SchemeHTTP Scheme = "http"
	// SchemeHTTPS ...
	SchemeHTTPS Scheme = "https"
)

// PossibleSchemeValues returns an array of possible values for the Scheme const type.
func PossibleSchemeValues() []Scheme {
	return []Scheme{SchemeHTTP, SchemeHTTPS}
}
