//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// VirtualMachineImageTemplatesCancelPoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesCancelPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesCancelPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesCancelPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesCancelResponse will be returned.
func (p *VirtualMachineImageTemplatesCancelPoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesCancelResponse, error) {
	respType := VirtualMachineImageTemplatesCancelResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachineImageTemplatesCancelResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesCancelPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesCreateOrUpdateResponse will be returned.
func (p *VirtualMachineImageTemplatesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesCreateOrUpdateResponse, error) {
	respType := VirtualMachineImageTemplatesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ImageTemplate)
	if err != nil {
		return VirtualMachineImageTemplatesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesDeletePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesDeleteResponse will be returned.
func (p *VirtualMachineImageTemplatesDeletePoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesDeleteResponse, error) {
	respType := VirtualMachineImageTemplatesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachineImageTemplatesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesRunPoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesRunPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesRunPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesRunPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesRunResponse will be returned.
func (p *VirtualMachineImageTemplatesRunPoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesRunResponse, error) {
	respType := VirtualMachineImageTemplatesRunResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachineImageTemplatesRunResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesRunPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachineImageTemplatesUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachineImageTemplatesUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachineImageTemplatesUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachineImageTemplatesUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachineImageTemplatesUpdateResponse will be returned.
func (p *VirtualMachineImageTemplatesUpdatePoller) FinalResponse(ctx context.Context) (VirtualMachineImageTemplatesUpdateResponse, error) {
	respType := VirtualMachineImageTemplatesUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ImageTemplate)
	if err != nil {
		return VirtualMachineImageTemplatesUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachineImageTemplatesUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}