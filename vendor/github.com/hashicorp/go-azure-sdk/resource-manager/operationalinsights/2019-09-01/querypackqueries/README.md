
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/querypackqueries` Documentation

The `querypackqueries` SDK allows for interaction with the Azure Resource Manager Service `operationalinsights` (API Version `2019-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/querypackqueries"
```


### Client Initialization

```go
client := querypackqueries.NewQueryPackQueriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `QueryPackQueriesClient.QueriesDelete`

```go
ctx := context.TODO()
id := querypackqueries.NewQueriesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackValue", "idValue")

read, err := client.QueriesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QueryPackQueriesClient.QueriesGet`

```go
ctx := context.TODO()
id := querypackqueries.NewQueriesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackValue", "idValue")

read, err := client.QueriesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QueryPackQueriesClient.QueriesList`

```go
ctx := context.TODO()
id := querypackqueries.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackValue")

// alternatively `client.QueriesList(ctx, id, querypackqueries.DefaultQueriesListOperationOptions())` can be used to do batched pagination
items, err := client.QueriesListComplete(ctx, id, querypackqueries.DefaultQueriesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `QueryPackQueriesClient.QueriesPut`

```go
ctx := context.TODO()
id := querypackqueries.NewQueriesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackValue", "idValue")

payload := querypackqueries.LogAnalyticsQueryPackQuery{
	// ...
}


read, err := client.QueriesPut(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QueryPackQueriesClient.QueriesSearch`

```go
ctx := context.TODO()
id := querypackqueries.NewQueryPackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackValue")

payload := querypackqueries.LogAnalyticsQueryPackQuerySearchProperties{
	// ...
}


// alternatively `client.QueriesSearch(ctx, id, payload, querypackqueries.DefaultQueriesSearchOperationOptions())` can be used to do batched pagination
items, err := client.QueriesSearchComplete(ctx, id, payload, querypackqueries.DefaultQueriesSearchOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `QueryPackQueriesClient.QueriesUpdate`

```go
ctx := context.TODO()
id := querypackqueries.NewQueriesID("12345678-1234-9876-4563-123456789012", "example-resource-group", "queryPackValue", "idValue")

payload := querypackqueries.LogAnalyticsQueryPackQuery{
	// ...
}


read, err := client.QueriesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
