package storagecache

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AscOperationsClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from
// either NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage
// Caches.
type AscOperationsClient struct {
	BaseClient
}

// NewAscOperationsClient creates an instance of the AscOperationsClient client.
func NewAscOperationsClient(subscriptionID string) AscOperationsClient {
	return NewAscOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAscOperationsClientWithBaseURI creates an instance of the AscOperationsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAscOperationsClientWithBaseURI(baseURI string, subscriptionID string) AscOperationsClient {
	return AscOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the status of an asynchronous operation for the Azure HPC Cache
// Parameters:
// location - the name of the region used to look up the operation.
// operationID - the operation id which uniquely identifies the asynchronous operation.
func (client AscOperationsClient) Get(ctx context.Context, location string, operationID string) (result AscOperation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AscOperationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.AscOperationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.AscOperationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.AscOperationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AscOperationsClient) GetPreparer(ctx context.Context, location string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/locations/{location}/ascOperations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AscOperationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AscOperationsClient) GetResponder(resp *http.Response) (result AscOperation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}