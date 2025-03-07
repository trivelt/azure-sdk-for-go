//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListOperations.json
func ExampleAppliancesClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppliancesClient().NewListOperationsPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ApplianceOperationsList = armresourceconnector.ApplianceOperationsList{
		// 	Value: []*armresourceconnector.ApplianceOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/operations/read"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Gets list of Available Operations for Appliances"),
		// 				Operation: to.Ptr("List Available Operations for Appliances"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/register/action"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Registers the subscription for Appliance resource provider and enables the creation of Appliance."),
		// 				Operation: to.Ptr("Registers the Appliance Resource Provider"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliances Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/appliances/read"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Gets an Appliance resource"),
		// 				Operation: to.Ptr("Get Appliance"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/appliances/write"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Creates or Updates Appliance resource"),
		// 				Operation: to.Ptr("Create or Update Appliance"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/appliances/delete"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Deletes Appliance resource"),
		// 				Operation: to.Ptr("Delete Appliance"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/appliances/listClusterUserCredential"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Get an appliance cluster user credential"),
		// 				Operation: to.Ptr("List User Cluster Credential"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/locations/operationsstatus/read"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Get result of Appliance operation"),
		// 				Operation: to.Ptr("Get status of Appliance operation"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliance Operation Status"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ResourceConnector/locations/operationresults/read"),
		// 			Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 				Description: to.Ptr("Get result of Appliance operation"),
		// 				Operation: to.Ptr("Get the status of Appliance operation"),
		// 				Provider: to.Ptr("Microsoft ResourceConnector"),
		// 				Resource: to.Ptr("Appliance Operation Result"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListBySubscription.json
func ExampleAppliancesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppliancesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ApplianceListResult = armresourceconnector.ApplianceListResult{
		// 	Value: []*armresourceconnector.Appliance{
		// 		{
		// 			Name: to.Ptr("appliance01"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armresourceconnector.ApplianceProperties{
		// 				Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 				InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 					Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armresourceconnector.StatusConnected),
		// 				Version: to.Ptr("1.0.1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("appliance02"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance02"),
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armresourceconnector.ApplianceProperties{
		// 				Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 				InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 					Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armresourceconnector.StatusConnected),
		// 				Version: to.Ptr("1.0.1"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/TelemetryConfig.json
func ExampleAppliancesClient_GetTelemetryConfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().GetTelemetryConfig(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplianceGetTelemetryConfigResult = armresourceconnector.ApplianceGetTelemetryConfigResult{
	// 	TelemetryInstrumentationKey: to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListByResourceGroup.json
func ExampleAppliancesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppliancesClient().NewListByResourceGroupPager("testresourcegroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ApplianceListResult = armresourceconnector.ApplianceListResult{
		// 	Value: []*armresourceconnector.Appliance{
		// 		{
		// 			Name: to.Ptr("appliance01"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armresourceconnector.ApplianceProperties{
		// 				Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 				InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 					Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armresourceconnector.StatusConnected),
		// 				Version: to.Ptr("1.0.1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("appliance02"),
		// 			Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance02"),
		// 			SystemData: &armresourceconnector.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armresourceconnector.ApplianceProperties{
		// 				Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
		// 				InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
		// 					Provider: to.Ptr(armresourceconnector.ProviderVMWare),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armresourceconnector.StatusConnected),
		// 				Version: to.Ptr("1.0.1"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesGet.json
func ExampleAppliancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().Get(ctx, "testresourcegroup", "appliance01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Appliance = armresourceconnector.Appliance{
	// 	Name: to.Ptr("appliance01"),
	// 	Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	SystemData: &armresourceconnector.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armresourceconnector.Identity{
	// 		Type: to.Ptr(armresourceconnector.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armresourceconnector.ApplianceProperties{
	// 		Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
	// 		InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
	// 			Provider: to.Ptr(armresourceconnector.ProviderVMWare),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armresourceconnector.StatusConnected),
	// 		Version: to.Ptr("1.0.1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesCreate_Update.json
func ExampleAppliancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppliancesClient().BeginCreateOrUpdate(ctx, "testresourcegroup", "appliance01", armresourceconnector.Appliance{
		Location: to.Ptr("West US"),
		Properties: &armresourceconnector.ApplianceProperties{
			Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
			InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
				Provider: to.Ptr(armresourceconnector.ProviderVMWare),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Appliance = armresourceconnector.Appliance{
	// 	Name: to.Ptr("appliance01"),
	// 	Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	SystemData: &armresourceconnector.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armresourceconnector.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armresourceconnector.Identity{
	// 		Type: to.Ptr(armresourceconnector.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armresourceconnector.ApplianceProperties{
	// 		Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
	// 		InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
	// 			Provider: to.Ptr(armresourceconnector.ProviderVMWare),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armresourceconnector.StatusConnected),
	// 		Version: to.Ptr("1.0.1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesDelete.json
func ExampleAppliancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAppliancesClient().BeginDelete(ctx, "testresourcegroup", "appliance01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesPatch.json
func ExampleAppliancesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().Update(ctx, "testresourcegroup", "appliance01", armresourceconnector.PatchableAppliance{
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Appliance = armresourceconnector.Appliance{
	// 	Name: to.Ptr("appliance01"),
	// 	Type: to.Ptr("Microsoft.ResourceConnector/appliances"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Identity: &armresourceconnector.Identity{
	// 		Type: to.Ptr(armresourceconnector.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armresourceconnector.ApplianceProperties{
	// 		Distro: to.Ptr(armresourceconnector.DistroAKSEdge),
	// 		InfrastructureConfig: &armresourceconnector.AppliancePropertiesInfrastructureConfig{
	// 			Provider: to.Ptr(armresourceconnector.ProviderVMWare),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armresourceconnector.StatusValidating),
	// 		Version: to.Ptr("1.0.1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListClusterUserCredential.json
func ExampleAppliancesClient_ListClusterUserCredential() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().ListClusterUserCredential(ctx, "testresourcegroup", "appliance01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplianceListCredentialResults = armresourceconnector.ApplianceListCredentialResults{
	// 	HybridConnectionConfig: &armresourceconnector.HybridConnectionConfig{
	// 		ExpirationTime: to.Ptr[int64](123456789),
	// 		HybridConnectionName: to.Ptr("provider/type/bc36ffcf318d5bedfc05ba8b0628dba"),
	// 		Relay: to.Ptr("relayName"),
	// 		Token: to.Ptr("mockSecretOtherprovider/type/bc36ffcf318d5bedfc05ba91c157ReceiverToken"),
	// 	},
	// 	Kubeconfigs: []*armresourceconnector.ApplianceCredentialKubeconfig{
	// 		{
	// 			Name: to.Ptr(armresourceconnector.AccessProfileType("kubeconfigName1")),
	// 			Value: to.Ptr("xxxxxxxxxxxxx"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListKeys.json
func ExampleAppliancesClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().ListKeys(ctx, "testresourcegroup", "appliance01", &armresourceconnector.AppliancesClientListKeysOptions{ArtifactType: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplianceListKeysResults = armresourceconnector.ApplianceListKeysResults{
	// 	ArtifactProfiles: map[string]*armresourceconnector.ArtifactProfile{
	// 		"LogsArtifactType": &armresourceconnector.ArtifactProfile{
	// 			Endpoint: to.Ptr("https://<storage-account-name>.blob.core.windows.net/<container-name>?<SAS-token>"),
	// 		},
	// 	},
	// 	Kubeconfigs: []*armresourceconnector.ApplianceCredentialKubeconfig{
	// 		{
	// 			Name: to.Ptr(armresourceconnector.AccessProfileType("kubeconfigName1")),
	// 			Value: to.Ptr("xxxxxxxxxxxxx"),
	// 	}},
	// 	SSHKeys: map[string]*armresourceconnector.SSHKey{
	// 		"LogsKey": &armresourceconnector.SSHKey{
	// 			Certificate: to.Ptr("<Generated Certificate>"),
	// 			CreationTimeStamp: to.Ptr[int64](1660946559),
	// 			ExpirationTimeStamp: to.Ptr[int64](1724119358),
	// 			PrivateKey: to.Ptr("<Generated Private Key>"),
	// 		},
	// 		"ManagementCAKey": &armresourceconnector.SSHKey{
	// 			PublicKey: to.Ptr("<Generated Public Key>"),
	// 		},
	// 		"SSHCustomerUser": &armresourceconnector.SSHKey{
	// 			PrivateKey: to.Ptr("xxxxxxxx"),
	// 			PublicKey: to.Ptr("xxxxxxxx"),
	// 		},
	// 		"ScopedAccessKey": &armresourceconnector.SSHKey{
	// 			Certificate: to.Ptr("<Generated Certificate>"),
	// 			CreationTimeStamp: to.Ptr[int64](1660946559),
	// 			ExpirationTimeStamp: to.Ptr[int64](1724119358),
	// 			PrivateKey: to.Ptr("<Generated Private Key>"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5f700acd3d094d8eedca381932f2e7916afd2e55/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/UpgradeGraph.json
func ExampleAppliancesClient_GetUpgradeGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().GetUpgradeGraph(ctx, "testresourcegroup", "appliance01", "stable", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpgradeGraph = armresourceconnector.UpgradeGraph{
	// 	Name: to.Ptr("stable"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ResourceConnector/appliances/appliance01"),
	// 	Properties: &armresourceconnector.UpgradeGraphProperties{
	// 		ApplianceVersion: to.Ptr("1.0.0"),
	// 		SupportedVersions: []*armresourceconnector.SupportedVersion{
	// 			{
	// 				Metadata: &armresourceconnector.SupportedVersionMetadata{
	// 					CatalogVersion: &armresourceconnector.SupportedVersionCatalogVersion{
	// 						Name: to.Ptr("cloudop-product-information"),
	// 						Data: &armresourceconnector.SupportedVersionCatalogVersionData{
	// 							Audience: to.Ptr("stable"),
	// 							Catalog: to.Ptr("arc-appliance-stable-catalogs-ext"),
	// 							Offer: to.Ptr("arcappliance"),
	// 							Version: to.Ptr("0.1.5.11115"),
	// 						},
	// 						Namespace: to.Ptr("cloudop-system"),
	// 					},
	// 				},
	// 				Version: to.Ptr("1.0.1"),
	// 		}},
	// 	},
	// }
}
