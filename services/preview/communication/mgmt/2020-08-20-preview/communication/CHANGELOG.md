Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewOperationListPage` parameter(s) have been changed from `(func(context.Context, OperationList) (OperationList, error))` to `(OperationList, func(context.Context, OperationList) (OperationList, error))`
- Function `NewServiceResourceListPage` parameter(s) have been changed from `(func(context.Context, ServiceResourceList) (ServiceResourceList, error))` to `(ServiceResourceList, func(context.Context, ServiceResourceList) (ServiceResourceList, error))`

## New Content

- New function `ServiceClient.CheckNameAvailabilitySender(*http.Request) (*http.Response, error)`
- New function `ServiceClient.CheckNameAvailabilityPreparer(context.Context, *NameAvailabilityParameters) (*http.Request, error)`
- New function `ServiceClient.CheckNameAvailability(context.Context, *NameAvailabilityParameters) (NameAvailability, error)`
- New function `ServiceClient.CheckNameAvailabilityResponder(*http.Response) (NameAvailability, error)`
- New struct `NameAvailability`
- New struct `NameAvailabilityParameters`