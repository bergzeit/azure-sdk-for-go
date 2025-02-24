//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoleManagementPolicyAssignmentsClient contains the methods for the RoleManagementPolicyAssignments group.
// Don't use this type directly, use NewRoleManagementPolicyAssignmentsClient() instead.
type RoleManagementPolicyAssignmentsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRoleManagementPolicyAssignmentsClient creates a new instance of RoleManagementPolicyAssignmentsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleManagementPolicyAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleManagementPolicyAssignmentsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RoleManagementPolicyAssignmentsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Create - Create a role management policy assignment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role management policy assignment to upsert.
// roleManagementPolicyAssignmentName - The name of format {guid_guid} the role management policy assignment to upsert.
// parameters - Parameters for the role management policy assignment.
// options - RoleManagementPolicyAssignmentsClientCreateOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Create
// method.
func (client *RoleManagementPolicyAssignmentsClient) Create(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment, options *RoleManagementPolicyAssignmentsClientCreateOptions) (RoleManagementPolicyAssignmentsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, parameters, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleManagementPolicyAssignmentsClient) createCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, parameters RoleManagementPolicyAssignment, options *RoleManagementPolicyAssignmentsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleManagementPolicyAssignmentsClient) createHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsClientCreateResponse, error) {
	result := RoleManagementPolicyAssignmentsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignment); err != nil {
		return RoleManagementPolicyAssignmentsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a role management policy assignment
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role management policy assignment to delete.
// roleManagementPolicyAssignmentName - The name of format {guid_guid} the role management policy assignment to delete.
// options - RoleManagementPolicyAssignmentsClientDeleteOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Delete
// method.
func (client *RoleManagementPolicyAssignmentsClient) Delete(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientDeleteOptions) (RoleManagementPolicyAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleManagementPolicyAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RoleManagementPolicyAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoleManagementPolicyAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified role management policy assignment for a resource scope
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role management policy.
// roleManagementPolicyAssignmentName - The name of format {guid_guid} the role management policy assignment to get.
// options - RoleManagementPolicyAssignmentsClientGetOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Get
// method.
func (client *RoleManagementPolicyAssignmentsClient) Get(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientGetOptions) (RoleManagementPolicyAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleManagementPolicyAssignmentName, options)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleManagementPolicyAssignmentsClient) getCreateRequest(ctx context.Context, scope string, roleManagementPolicyAssignmentName string, options *RoleManagementPolicyAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{roleManagementPolicyAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyAssignmentName == "" {
		return nil, errors.New("parameter roleManagementPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyAssignmentName}", url.PathEscape(roleManagementPolicyAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleManagementPolicyAssignmentsClient) getHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsClientGetResponse, error) {
	result := RoleManagementPolicyAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignment); err != nil {
		return RoleManagementPolicyAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role management assignment policies for a resource scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role management policy.
// options - RoleManagementPolicyAssignmentsClientListForScopeOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.ListForScope
// method.
func (client *RoleManagementPolicyAssignmentsClient) NewListForScopePager(scope string, options *RoleManagementPolicyAssignmentsClientListForScopeOptions) *runtime.Pager[RoleManagementPolicyAssignmentsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleManagementPolicyAssignmentsClientListForScopeResponse]{
		More: func(page RoleManagementPolicyAssignmentsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleManagementPolicyAssignmentsClientListForScopeResponse) (RoleManagementPolicyAssignmentsClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleManagementPolicyAssignmentsClientListForScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleManagementPolicyAssignmentsClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleManagementPolicyAssignmentsClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleManagementPolicyAssignmentsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicyAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleManagementPolicyAssignmentsClient) listForScopeHandleResponse(resp *http.Response) (RoleManagementPolicyAssignmentsClientListForScopeResponse, error) {
	result := RoleManagementPolicyAssignmentsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyAssignmentListResult); err != nil {
		return RoleManagementPolicyAssignmentsClientListForScopeResponse{}, err
	}
	return result, nil
}
