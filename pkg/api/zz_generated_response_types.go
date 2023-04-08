//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package api

// DeploymentManagementClientCreatEventSubscriptionResponse contains the response from method DeploymentManagementClient.CreatEventSubscription.
type DeploymentManagementClientCreatEventSubscriptionResponse struct {
	CreateEventSubscriptionResponse
}

// DeploymentManagementClientCreateDeploymentResponse contains the response from method DeploymentManagementClient.CreateDeployment.
type DeploymentManagementClientCreateDeploymentResponse struct {
	Deployment
}

// DeploymentManagementClientDeleteEventSubscriptionResponse contains the response from method DeploymentManagementClient.DeleteEventSubscription.
type DeploymentManagementClientDeleteEventSubscriptionResponse struct {
	// placeholder for future response values
}

// DeploymentManagementClientGetDeploymentResponse contains the response from method DeploymentManagementClient.GetDeployment.
type DeploymentManagementClientGetDeploymentResponse struct {
	Deployment
}

// DeploymentManagementClientGetEventSubscriptionResponse contains the response from method DeploymentManagementClient.GetEventSubscription.
type DeploymentManagementClientGetEventSubscriptionResponse struct {
	EventSubscriptionResponse
}

// DeploymentManagementClientGetEventTypesResponse contains the response from method DeploymentManagementClient.GetEventTypes.
type DeploymentManagementClientGetEventTypesResponse struct {
	// Array of EventType
	EventTypeArray []*EventType
}

// DeploymentManagementClientGetInvokedDeploymentOperationResponse contains the response from method DeploymentManagementClient.GetInvokedDeploymentOperation.
type DeploymentManagementClientGetInvokedDeploymentOperationResponse struct {
	InvokedDeploymentOperationResponse
}

// DeploymentManagementClientInvokeDeploymentOperationResponse contains the response from method DeploymentManagementClient.InvokeDeploymentOperation.
type DeploymentManagementClientInvokeDeploymentOperationResponse struct {
	InvokedDeploymentOperationResponse
}

// DeploymentManagementClientListDeploymentsResponse contains the response from method DeploymentManagementClient.ListDeployments.
type DeploymentManagementClientListDeploymentsResponse struct {
	// Array of Deployment
	DeploymentArray []*Deployment
}

// DeploymentManagementClientListEventSubscriptionsResponse contains the response from method DeploymentManagementClient.ListEventSubscriptions.
type DeploymentManagementClientListEventSubscriptionsResponse struct {
	// Array of EventSubscriptionResponse
	EventSubscriptionResponseArray []*EventSubscriptionResponse
}

// DeploymentManagementClientListOperationsResponse contains the response from method DeploymentManagementClient.ListOperations.
type DeploymentManagementClientListOperationsResponse struct {
	// Array of Operation
	OperationArray []*Operation
}

// DeploymentManagementClientUpdateDeploymentResponse contains the response from method DeploymentManagementClient.UpdateDeployment.
type DeploymentManagementClientUpdateDeploymentResponse struct {
	// placeholder for future response values
}
