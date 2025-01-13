# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**ClientCertificates** | Pointer to [**ClientCertificates**](ClientCertificates.md) |  | [optional] 
**CustomDomain** | Pointer to [**Domain**](Domain.md) |  | [optional] 
**DetailedActivityLogs** | Pointer to [**DetailedActivityLogs**](DetailedActivityLogs.md) |  | [optional] 
**DetailedTracing** | Pointer to [**DetailedTracing**](DetailedTracing.md) |  | [optional] 
**EventHandlers** | Pointer to [**[]EventHandler**](EventHandler.md) |  | [optional] 
**GatewayTags** | Pointer to [**GatewayTags**](GatewayTags.md) |  | [optional] 
**ListenPath** | Pointer to [**ListenPath**](ListenPath.md) |  | [optional] 

## Methods

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *Server) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *Server) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *Server) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *Server) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetClientCertificates

`func (o *Server) GetClientCertificates() ClientCertificates`

GetClientCertificates returns the ClientCertificates field if non-nil, zero value otherwise.

### GetClientCertificatesOk

`func (o *Server) GetClientCertificatesOk() (*ClientCertificates, bool)`

GetClientCertificatesOk returns a tuple with the ClientCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificates

`func (o *Server) SetClientCertificates(v ClientCertificates)`

SetClientCertificates sets ClientCertificates field to given value.

### HasClientCertificates

`func (o *Server) HasClientCertificates() bool`

HasClientCertificates returns a boolean if a field has been set.

### GetCustomDomain

`func (o *Server) GetCustomDomain() Domain`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *Server) GetCustomDomainOk() (*Domain, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *Server) SetCustomDomain(v Domain)`

SetCustomDomain sets CustomDomain field to given value.

### HasCustomDomain

`func (o *Server) HasCustomDomain() bool`

HasCustomDomain returns a boolean if a field has been set.

### GetDetailedActivityLogs

`func (o *Server) GetDetailedActivityLogs() DetailedActivityLogs`

GetDetailedActivityLogs returns the DetailedActivityLogs field if non-nil, zero value otherwise.

### GetDetailedActivityLogsOk

`func (o *Server) GetDetailedActivityLogsOk() (*DetailedActivityLogs, bool)`

GetDetailedActivityLogsOk returns a tuple with the DetailedActivityLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedActivityLogs

`func (o *Server) SetDetailedActivityLogs(v DetailedActivityLogs)`

SetDetailedActivityLogs sets DetailedActivityLogs field to given value.

### HasDetailedActivityLogs

`func (o *Server) HasDetailedActivityLogs() bool`

HasDetailedActivityLogs returns a boolean if a field has been set.

### GetDetailedTracing

`func (o *Server) GetDetailedTracing() DetailedTracing`

GetDetailedTracing returns the DetailedTracing field if non-nil, zero value otherwise.

### GetDetailedTracingOk

`func (o *Server) GetDetailedTracingOk() (*DetailedTracing, bool)`

GetDetailedTracingOk returns a tuple with the DetailedTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedTracing

`func (o *Server) SetDetailedTracing(v DetailedTracing)`

SetDetailedTracing sets DetailedTracing field to given value.

### HasDetailedTracing

`func (o *Server) HasDetailedTracing() bool`

HasDetailedTracing returns a boolean if a field has been set.

### GetEventHandlers

`func (o *Server) GetEventHandlers() []EventHandler`

GetEventHandlers returns the EventHandlers field if non-nil, zero value otherwise.

### GetEventHandlersOk

`func (o *Server) GetEventHandlersOk() (*[]EventHandler, bool)`

GetEventHandlersOk returns a tuple with the EventHandlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlers

`func (o *Server) SetEventHandlers(v []EventHandler)`

SetEventHandlers sets EventHandlers field to given value.

### HasEventHandlers

`func (o *Server) HasEventHandlers() bool`

HasEventHandlers returns a boolean if a field has been set.

### GetGatewayTags

`func (o *Server) GetGatewayTags() GatewayTags`

GetGatewayTags returns the GatewayTags field if non-nil, zero value otherwise.

### GetGatewayTagsOk

`func (o *Server) GetGatewayTagsOk() (*GatewayTags, bool)`

GetGatewayTagsOk returns a tuple with the GatewayTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTags

`func (o *Server) SetGatewayTags(v GatewayTags)`

SetGatewayTags sets GatewayTags field to given value.

### HasGatewayTags

`func (o *Server) HasGatewayTags() bool`

HasGatewayTags returns a boolean if a field has been set.

### GetListenPath

`func (o *Server) GetListenPath() ListenPath`

GetListenPath returns the ListenPath field if non-nil, zero value otherwise.

### GetListenPathOk

`func (o *Server) GetListenPathOk() (*ListenPath, bool)`

GetListenPathOk returns a tuple with the ListenPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPath

`func (o *Server) SetListenPath(v ListenPath)`

SetListenPath sets ListenPath field to given value.

### HasListenPath

`func (o *Server) HasListenPath() bool`

HasListenPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


