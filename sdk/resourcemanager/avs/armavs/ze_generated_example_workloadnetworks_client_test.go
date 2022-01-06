//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListSegments.json
func ExampleWorkloadNetworksClient_ListSegments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListSegments("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkSegment.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetSegments.json
func ExampleWorkloadNetworksClient_GetSegment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetSegment(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<segment-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkSegment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateSegments.json
func ExampleWorkloadNetworksClient_BeginCreateSegments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateSegments(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<segment-id>",
		armavs.WorkloadNetworkSegment{
			Properties: &armavs.WorkloadNetworkSegmentProperties{
				ConnectedGateway: to.StringPtr("<connected-gateway>"),
				DisplayName:      to.StringPtr("<display-name>"),
				Revision:         to.Int64Ptr(1),
				Subnet: &armavs.WorkloadNetworkSegmentSubnet{
					DhcpRanges: []*string{
						to.StringPtr("40.20.0.0-40.20.0.1")},
					GatewayAddress: to.StringPtr("<gateway-address>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkSegment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateSegments.json
func ExampleWorkloadNetworksClient_BeginUpdateSegments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateSegments(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<segment-id>",
		armavs.WorkloadNetworkSegment{
			Properties: &armavs.WorkloadNetworkSegmentProperties{
				ConnectedGateway: to.StringPtr("<connected-gateway>"),
				Revision:         to.Int64Ptr(1),
				Subnet: &armavs.WorkloadNetworkSegmentSubnet{
					DhcpRanges: []*string{
						to.StringPtr("40.20.0.0-40.20.0.1")},
					GatewayAddress: to.StringPtr("<gateway-address>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkSegment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteSegments.json
func ExampleWorkloadNetworksClient_BeginDeleteSegment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteSegment(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<segment-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListDhcpConfigurations.json
func ExampleWorkloadNetworksClient_ListDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListDhcp("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkDhcp.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetDhcpConfigurations.json
func ExampleWorkloadNetworksClient_GetDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetDhcp(ctx,
		"<resource-group-name>",
		"<dhcp-id>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDhcp.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateDhcpConfigurations.json
func ExampleWorkloadNetworksClient_BeginCreateDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateDhcp(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dhcp-id>",
		armavs.WorkloadNetworkDhcp{
			Properties: &armavs.WorkloadNetworkDhcpServer{
				WorkloadNetworkDhcpEntity: armavs.WorkloadNetworkDhcpEntity{
					DhcpType:    armavs.DhcpTypeEnumSERVER.ToPtr(),
					DisplayName: to.StringPtr("<display-name>"),
					Revision:    to.Int64Ptr(1),
				},
				LeaseTime:     to.Int64Ptr(86400),
				ServerAddress: to.StringPtr("<server-address>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDhcp.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDhcpConfigurations.json
func ExampleWorkloadNetworksClient_BeginUpdateDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDhcp(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dhcp-id>",
		armavs.WorkloadNetworkDhcp{
			Properties: &armavs.WorkloadNetworkDhcpServer{
				WorkloadNetworkDhcpEntity: armavs.WorkloadNetworkDhcpEntity{
					DhcpType: armavs.DhcpTypeEnumSERVER.ToPtr(),
					Revision: to.Int64Ptr(1),
				},
				LeaseTime:     to.Int64Ptr(86400),
				ServerAddress: to.StringPtr("<server-address>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDhcp.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteDhcpConfigurations.json
func ExampleWorkloadNetworksClient_BeginDeleteDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteDhcp(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dhcp-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListGateways.json
func ExampleWorkloadNetworksClient_ListGateways() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListGateways("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkGateway.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetGateway.json
func ExampleWorkloadNetworksClient_GetGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetGateway(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<gateway-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkGateway.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListPortMirroringProfiles.json
func ExampleWorkloadNetworksClient_ListPortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListPortMirroring("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkPortMirroring.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetPortMirroringProfiles.json
func ExampleWorkloadNetworksClient_GetPortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetPortMirroring(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<port-mirroring-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkPortMirroring.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreatePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginCreatePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreatePortMirroring(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<port-mirroring-id>",
		armavs.WorkloadNetworkPortMirroring{
			Properties: &armavs.WorkloadNetworkPortMirroringProperties{
				Destination: to.StringPtr("<destination>"),
				Direction:   armavs.PortMirroringDirectionEnumBIDIRECTIONAL.ToPtr(),
				DisplayName: to.StringPtr("<display-name>"),
				Revision:    to.Int64Ptr(1),
				Source:      to.StringPtr("<source>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkPortMirroring.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdatePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginUpdatePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdatePortMirroring(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<port-mirroring-id>",
		armavs.WorkloadNetworkPortMirroring{
			Properties: &armavs.WorkloadNetworkPortMirroringProperties{
				Destination: to.StringPtr("<destination>"),
				Direction:   armavs.PortMirroringDirectionEnumBIDIRECTIONAL.ToPtr(),
				Revision:    to.Int64Ptr(1),
				Source:      to.StringPtr("<source>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkPortMirroring.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeletePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginDeletePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeletePortMirroring(ctx,
		"<resource-group-name>",
		"<port-mirroring-id>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListVMGroups.json
func ExampleWorkloadNetworksClient_ListVMGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListVMGroups("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkVMGroup.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetVMGroups.json
func ExampleWorkloadNetworksClient_GetVMGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetVMGroup(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<vm-group-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkVMGroup.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateVMGroups.json
func ExampleWorkloadNetworksClient_BeginCreateVMGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateVMGroup(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<vm-group-id>",
		armavs.WorkloadNetworkVMGroup{
			Properties: &armavs.WorkloadNetworkVMGroupProperties{
				DisplayName: to.StringPtr("<display-name>"),
				Members: []*string{
					to.StringPtr("564d43da-fefc-2a3b-1d92-42855622fa50")},
				Revision: to.Int64Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkVMGroup.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateVMGroups.json
func ExampleWorkloadNetworksClient_BeginUpdateVMGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateVMGroup(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<vm-group-id>",
		armavs.WorkloadNetworkVMGroup{
			Properties: &armavs.WorkloadNetworkVMGroupProperties{
				Members: []*string{
					to.StringPtr("564d43da-fefc-2a3b-1d92-42855622fa50")},
				Revision: to.Int64Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkVMGroup.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteVMGroups.json
func ExampleWorkloadNetworksClient_BeginDeleteVMGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteVMGroup(ctx,
		"<resource-group-name>",
		"<vm-group-id>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListVirtualMachines.json
func ExampleWorkloadNetworksClient_ListVirtualMachines() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListVirtualMachines("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkVirtualMachine.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetVirtualMachine.json
func ExampleWorkloadNetworksClient_GetVirtualMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetVirtualMachine(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<virtual-machine-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkVirtualMachine.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListDnsServices.json
func ExampleWorkloadNetworksClient_ListDNSServices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListDNSServices("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkDNSService.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetDnsServices.json
func ExampleWorkloadNetworksClient_GetDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetDNSService(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-service-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateDnsServices.json
func ExampleWorkloadNetworksClient_BeginCreateDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateDNSService(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-service-id>",
		armavs.WorkloadNetworkDNSService{
			Properties: &armavs.WorkloadNetworkDNSServiceProperties{
				DefaultDNSZone: to.StringPtr("<default-dnszone>"),
				DisplayName:    to.StringPtr("<display-name>"),
				DNSServiceIP:   to.StringPtr("<dnsservice-ip>"),
				FqdnZones: []*string{
					to.StringPtr("fqdnZone1")},
				LogLevel: armavs.DNSServiceLogLevelEnumINFO.ToPtr(),
				Revision: to.Int64Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDnsServices.json
func ExampleWorkloadNetworksClient_BeginUpdateDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDNSService(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-service-id>",
		armavs.WorkloadNetworkDNSService{
			Properties: &armavs.WorkloadNetworkDNSServiceProperties{
				DefaultDNSZone: to.StringPtr("<default-dnszone>"),
				DisplayName:    to.StringPtr("<display-name>"),
				DNSServiceIP:   to.StringPtr("<dnsservice-ip>"),
				FqdnZones: []*string{
					to.StringPtr("fqdnZone1")},
				LogLevel: armavs.DNSServiceLogLevelEnumINFO.ToPtr(),
				Revision: to.Int64Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteDnsServices.json
func ExampleWorkloadNetworksClient_BeginDeleteDNSService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteDNSService(ctx,
		"<resource-group-name>",
		"<dns-service-id>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListDnsZones.json
func ExampleWorkloadNetworksClient_ListDNSZones() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListDNSZones("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkDNSZone.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetDnsZones.json
func ExampleWorkloadNetworksClient_GetDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetDNSZone(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-zone-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSZone.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateDnsZones.json
func ExampleWorkloadNetworksClient_BeginCreateDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateDNSZone(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-zone-id>",
		armavs.WorkloadNetworkDNSZone{
			Properties: &armavs.WorkloadNetworkDNSZoneProperties{
				DisplayName: to.StringPtr("<display-name>"),
				DNSServerIPs: []*string{
					to.StringPtr("1.1.1.1")},
				Domain:   []*string{},
				Revision: to.Int64Ptr(1),
				SourceIP: to.StringPtr("<source-ip>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSZone.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDnsZones.json
func ExampleWorkloadNetworksClient_BeginUpdateDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDNSZone(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dns-zone-id>",
		armavs.WorkloadNetworkDNSZone{
			Properties: &armavs.WorkloadNetworkDNSZoneProperties{
				DisplayName: to.StringPtr("<display-name>"),
				DNSServerIPs: []*string{
					to.StringPtr("1.1.1.1")},
				Domain:   []*string{},
				Revision: to.Int64Ptr(1),
				SourceIP: to.StringPtr("<source-ip>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkDNSZone.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeleteDnsZones.json
func ExampleWorkloadNetworksClient_BeginDeleteDNSZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeleteDNSZone(ctx,
		"<resource-group-name>",
		"<dns-zone-id>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_ListPublicIPs.json
func ExampleWorkloadNetworksClient_ListPublicIPs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	pager := client.ListPublicIPs("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("WorkloadNetworkPublicIP.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetPublicIPs.json
func ExampleWorkloadNetworksClient_GetPublicIP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetPublicIP(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<public-ipid>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkPublicIP.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreatePublicIPs.json
func ExampleWorkloadNetworksClient_BeginCreatePublicIP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreatePublicIP(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<public-ipid>",
		armavs.WorkloadNetworkPublicIP{
			Properties: &armavs.WorkloadNetworkPublicIPProperties{
				DisplayName:       to.StringPtr("<display-name>"),
				NumberOfPublicIPs: to.Int64Ptr(32),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkPublicIP.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_DeletePublicIPs.json
func ExampleWorkloadNetworksClient_BeginDeletePublicIP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDeletePublicIP(ctx,
		"<resource-group-name>",
		"<public-ipid>",
		"<private-cloud-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}