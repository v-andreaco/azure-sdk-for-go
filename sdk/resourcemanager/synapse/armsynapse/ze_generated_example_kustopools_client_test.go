//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCheckNameAvailability.json
func ExampleKustoPoolsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		"<location>",
		armsynapse.KustoPoolCheckNameRequest{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientCheckNameAvailabilityResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListByWorkspace.json
func ExampleKustoPoolsClient_ListByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	res, err := client.ListByWorkspace(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientListByWorkspaceResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsGet.json
func ExampleKustoPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientGetResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCreateOrUpdate.json
func ExampleKustoPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<workspace-name>",
		"<resource-group-name>",
		"<kusto-pool-name>",
		armsynapse.KustoPool{
			Location: to.StringPtr("<location>"),
			Properties: &armsynapse.KustoPoolProperties{
				EnablePurge:           to.BoolPtr(true),
				EnableStreamingIngest: to.BoolPtr(true),
				WorkspaceUID:          to.StringPtr("<workspace-uid>"),
			},
			SKU: &armsynapse.AzureSKU{
				Name:     armsynapse.SKUName("Storage optimized").ToPtr(),
				Capacity: to.Int32Ptr(2),
				Size:     armsynapse.SKUSize("Medium").ToPtr(),
			},
		},
		&armsynapse.KustoPoolsClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsUpdate.json
func ExampleKustoPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<workspace-name>",
		"<resource-group-name>",
		"<kusto-pool-name>",
		armsynapse.KustoPoolUpdate{
			Properties: &armsynapse.KustoPoolProperties{
				EnablePurge:           to.BoolPtr(true),
				EnableStreamingIngest: to.BoolPtr(true),
				WorkspaceUID:          to.StringPtr("<workspace-uid>"),
			},
			SKU: &armsynapse.AzureSKU{
				Name:     armsynapse.SKUName("Storage optimized").ToPtr(),
				Capacity: to.Int32Ptr(2),
				Size:     armsynapse.SKUSize("Medium").ToPtr(),
			},
		},
		&armsynapse.KustoPoolsClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientUpdateResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsDelete.json
func ExampleKustoPoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<workspace-name>",
		"<resource-group-name>",
		"<kusto-pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStop.json
func ExampleKustoPoolsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStop(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStart.json
func ExampleKustoPoolsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsList.json
func ExampleKustoPoolsClient_ListLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	res, err := client.ListLanguageExtensions(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientListLanguageExtensionsResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsAdd.json
func ExampleKustoPoolsClient_BeginAddLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginAddLanguageExtensions(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		armsynapse.LanguageExtensionsList{
			Value: []*armsynapse.LanguageExtension{
				{
					LanguageExtensionName: armsynapse.LanguageExtensionName("PYTHON").ToPtr(),
				},
				{
					LanguageExtensionName: armsynapse.LanguageExtensionName("R").ToPtr(),
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsRemove.json
func ExampleKustoPoolsClient_BeginRemoveLanguageExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRemoveLanguageExtensions(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		armsynapse.LanguageExtensionsList{
			Value: []*armsynapse.LanguageExtension{
				{
					LanguageExtensionName: armsynapse.LanguageExtensionName("PYTHON").ToPtr(),
				},
				{
					LanguageExtensionName: armsynapse.LanguageExtensionName("R").ToPtr(),
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesList.json
func ExampleKustoPoolsClient_ListFollowerDatabases() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	res, err := client.ListFollowerDatabases(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.KustoPoolsClientListFollowerDatabasesResult)
}

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesDetach.json
func ExampleKustoPoolsClient_BeginDetachFollowerDatabases() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDetachFollowerDatabases(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		armsynapse.FollowerDatabaseDefinition{
			AttachedDatabaseConfigurationName: to.StringPtr("<attached-database-configuration-name>"),
			KustoPoolResourceID:               to.StringPtr("<kusto-pool-resource-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
