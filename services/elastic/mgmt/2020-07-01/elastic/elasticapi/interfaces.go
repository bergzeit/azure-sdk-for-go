package elasticapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/elastic/mgmt/2020-07-01/elastic"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result elastic.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result elastic.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*elastic.OperationsClient)(nil)

// MonitorsClientAPI contains the set of methods on the MonitorsClient type.
type MonitorsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, monitorName string, body *elastic.MonitorResource) (result elastic.MonitorsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.MonitorsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.MonitorResource, err error)
	List(ctx context.Context) (result elastic.MonitorResourceListResponsePage, err error)
	ListComplete(ctx context.Context) (result elastic.MonitorResourceListResponseIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result elastic.MonitorResourceListResponsePage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result elastic.MonitorResourceListResponseIterator, err error)
	Update(ctx context.Context, resourceGroupName string, monitorName string, body *elastic.MonitorResourceUpdateParameters) (result elastic.MonitorResource, err error)
}

var _ MonitorsClientAPI = (*elastic.MonitorsClient)(nil)

// MonitoredResourcesClientAPI contains the set of methods on the MonitoredResourcesClient type.
type MonitoredResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.MonitoredResourceListResponsePage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.MonitoredResourceListResponseIterator, err error)
}

var _ MonitoredResourcesClientAPI = (*elastic.MonitoredResourcesClient)(nil)

// DeploymentInfoClientAPI contains the set of methods on the DeploymentInfoClient type.
type DeploymentInfoClientAPI interface {
	List(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.DeploymentInfoResponse, err error)
}

var _ DeploymentInfoClientAPI = (*elastic.DeploymentInfoClient)(nil)

// TagRulesClientAPI contains the set of methods on the TagRulesClient type.
type TagRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, body *elastic.MonitoringTagRules) (result elastic.MonitoringTagRules, err error)
	Delete(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string) (result elastic.TagRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string) (result elastic.MonitoringTagRules, err error)
	List(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.MonitoringTagRulesListResponsePage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.MonitoringTagRulesListResponseIterator, err error)
}

var _ TagRulesClientAPI = (*elastic.TagRulesClient)(nil)

// VMHostClientAPI contains the set of methods on the VMHostClient type.
type VMHostClientAPI interface {
	List(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.VMHostListResponsePage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.VMHostListResponseIterator, err error)
}

var _ VMHostClientAPI = (*elastic.VMHostClient)(nil)

// VMIngestionClientAPI contains the set of methods on the VMIngestionClient type.
type VMIngestionClientAPI interface {
	Details(ctx context.Context, resourceGroupName string, monitorName string) (result elastic.VMIngestionDetailsResponse, err error)
}

var _ VMIngestionClientAPI = (*elastic.VMIngestionClient)(nil)

// VMCollectionClientAPI contains the set of methods on the VMCollectionClient type.
type VMCollectionClientAPI interface {
	Update(ctx context.Context, resourceGroupName string, monitorName string, body *elastic.VMCollectionUpdate) (result autorest.Response, err error)
}

var _ VMCollectionClientAPI = (*elastic.VMCollectionClient)(nil)
