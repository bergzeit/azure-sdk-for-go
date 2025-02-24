//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DscConfigurationClient contains the methods for the DscConfiguration group.
// Don't use this type directly, use NewDscConfigurationClient() instead.
type DscConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDscConfigurationClient creates a new instance of DscConfigurationClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDscConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscConfigurationClient, error) {
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
	client := &DscConfigurationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdateWithJSON - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The create or update parameters for configuration.
// parameters - The create or update parameters for configuration.
// options - DscConfigurationClientCreateOrUpdateWithJSONOptions contains the optional parameters for the DscConfigurationClient.CreateOrUpdateWithJSON
// method.
func (client *DscConfigurationClient) CreateOrUpdateWithJSON(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationCreateOrUpdateParameters, options *DscConfigurationClientCreateOrUpdateWithJSONOptions) (DscConfigurationClientCreateOrUpdateWithJSONResponse, error) {
	req, err := client.createOrUpdateWithJSONCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateWithJSONHandleResponse(resp)
}

// createOrUpdateWithJSONCreateRequest creates the CreateOrUpdateWithJSON request.
func (client *DscConfigurationClient) createOrUpdateWithJSONCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationCreateOrUpdateParameters, options *DscConfigurationClientCreateOrUpdateWithJSONOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateWithJSONHandleResponse handles the CreateOrUpdateWithJSON response.
func (client *DscConfigurationClient) createOrUpdateWithJSONHandleResponse(resp *http.Response) (DscConfigurationClientCreateOrUpdateWithJSONResponse, error) {
	result := DscConfigurationClientCreateOrUpdateWithJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientCreateOrUpdateWithJSONResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateWithText - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The create or update parameters for configuration.
// parameters - The create or update parameters for configuration.
// options - DscConfigurationClientCreateOrUpdateWithTextOptions contains the optional parameters for the DscConfigurationClient.CreateOrUpdateWithText
// method.
func (client *DscConfigurationClient) CreateOrUpdateWithText(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationClientCreateOrUpdateWithTextOptions) (DscConfigurationClientCreateOrUpdateWithTextResponse, error) {
	req, err := client.createOrUpdateWithTextCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateWithTextHandleResponse(resp)
}

// createOrUpdateWithTextCreateRequest creates the CreateOrUpdateWithText request.
func (client *DscConfigurationClient) createOrUpdateWithTextCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationClientCreateOrUpdateWithTextOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	body := streaming.NopCloser(strings.NewReader(parameters))
	return req, req.SetBody(body, "text/plain; charset=utf-8")
}

// createOrUpdateWithTextHandleResponse handles the CreateOrUpdateWithText response.
func (client *DscConfigurationClient) createOrUpdateWithTextHandleResponse(resp *http.Response) (DscConfigurationClientCreateOrUpdateWithTextResponse, error) {
	result := DscConfigurationClientCreateOrUpdateWithTextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientCreateOrUpdateWithTextResponse{}, err
	}
	return result, nil
}

// Delete - Delete the dsc configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The configuration name.
// options - DscConfigurationClientDeleteOptions contains the optional parameters for the DscConfigurationClient.Delete method.
func (client *DscConfigurationClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientDeleteOptions) (DscConfigurationClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DscConfigurationClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DscConfigurationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The configuration name.
// options - DscConfigurationClientGetOptions contains the optional parameters for the DscConfigurationClient.Get method.
func (client *DscConfigurationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetOptions) (DscConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscConfigurationClient) getHandleResponse(resp *http.Response) (DscConfigurationClientGetResponse, error) {
	result := DscConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// GetContent - Retrieve the configuration script identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The configuration name.
// options - DscConfigurationClientGetContentOptions contains the optional parameters for the DscConfigurationClient.GetContent
// method.
func (client *DscConfigurationClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetContentOptions) (DscConfigurationClientGetContentResponse, error) {
	req, err := client.getContentCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationClientGetContentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientGetContentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationClientGetContentResponse{}, runtime.NewResponseError(resp)
	}
	return client.getContentHandleResponse(resp)
}

// getContentCreateRequest creates the GetContent request.
func (client *DscConfigurationClient) getContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationClientGetContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}/content"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"text/powershell"}
	return req, nil
}

// getContentHandleResponse handles the GetContent response.
func (client *DscConfigurationClient) getContentHandleResponse(resp *http.Response) (DscConfigurationClientGetContentResponse, error) {
	result := DscConfigurationClientGetContentResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return DscConfigurationClientGetContentResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of configurations.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - DscConfigurationClientListByAutomationAccountOptions contains the optional parameters for the DscConfigurationClient.ListByAutomationAccount
// method.
func (client *DscConfigurationClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *DscConfigurationClientListByAutomationAccountOptions) *runtime.Pager[DscConfigurationClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscConfigurationClientListByAutomationAccountResponse]{
		More: func(page DscConfigurationClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscConfigurationClientListByAutomationAccountResponse) (DscConfigurationClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscConfigurationClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DscConfigurationClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscConfigurationClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscConfigurationClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscConfigurationClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Inlinecount != nil {
		reqQP.Set("$inlinecount", *options.Inlinecount)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscConfigurationClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscConfigurationClientListByAutomationAccountResponse, error) {
	result := DscConfigurationClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfigurationListResult); err != nil {
		return DscConfigurationClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// UpdateWithJSON - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The create or update parameters for configuration.
// parameters - The create or update parameters for configuration.
// options - DscConfigurationClientUpdateWithJSONOptions contains the optional parameters for the DscConfigurationClient.UpdateWithJSON
// method.
func (client *DscConfigurationClient) UpdateWithJSON(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationUpdateParameters, options *DscConfigurationClientUpdateWithJSONOptions) (DscConfigurationClientUpdateWithJSONResponse, error) {
	req, err := client.updateWithJSONCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationClientUpdateWithJSONResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateWithJSONHandleResponse(resp)
}

// updateWithJSONCreateRequest creates the UpdateWithJSON request.
func (client *DscConfigurationClient) updateWithJSONCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationUpdateParameters, options *DscConfigurationClientUpdateWithJSONOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateWithJSONHandleResponse handles the UpdateWithJSON response.
func (client *DscConfigurationClient) updateWithJSONHandleResponse(resp *http.Response) (DscConfigurationClientUpdateWithJSONResponse, error) {
	result := DscConfigurationClientUpdateWithJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientUpdateWithJSONResponse{}, err
	}
	return result, nil
}

// UpdateWithText - Create the configuration identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// configurationName - The create or update parameters for configuration.
// parameters - The create or update parameters for configuration.
// options - DscConfigurationClientUpdateWithTextOptions contains the optional parameters for the DscConfigurationClient.UpdateWithText
// method.
func (client *DscConfigurationClient) UpdateWithText(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationClientUpdateWithTextOptions) (DscConfigurationClientUpdateWithTextResponse, error) {
	req, err := client.updateWithTextCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationClientUpdateWithTextResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateWithTextHandleResponse(resp)
}

// updateWithTextCreateRequest creates the UpdateWithText request.
func (client *DscConfigurationClient) updateWithTextCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationClientUpdateWithTextOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	body := streaming.NopCloser(strings.NewReader(parameters))
	return req, req.SetBody(body, "text/plain; charset=utf-8")
}

// updateWithTextHandleResponse handles the UpdateWithText response.
func (client *DscConfigurationClient) updateWithTextHandleResponse(resp *http.Response) (DscConfigurationClientUpdateWithTextResponse, error) {
	result := DscConfigurationClientUpdateWithTextResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationClientUpdateWithTextResponse{}, err
	}
	return result, nil
}
