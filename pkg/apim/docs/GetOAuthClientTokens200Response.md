# GetOAuthClientTokens200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationStatus**](PaginationStatus.md) |  | [optional] 
**Tokens** | Pointer to [**[]OAuthClientToken**](OAuthClientToken.md) |  | [optional] 

## Methods

### NewGetOAuthClientTokens200Response

`func NewGetOAuthClientTokens200Response() *GetOAuthClientTokens200Response`

NewGetOAuthClientTokens200Response instantiates a new GetOAuthClientTokens200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOAuthClientTokens200ResponseWithDefaults

`func NewGetOAuthClientTokens200ResponseWithDefaults() *GetOAuthClientTokens200Response`

NewGetOAuthClientTokens200ResponseWithDefaults instantiates a new GetOAuthClientTokens200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetOAuthClientTokens200Response) GetPagination() PaginationStatus`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetOAuthClientTokens200Response) GetPaginationOk() (*PaginationStatus, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetOAuthClientTokens200Response) SetPagination(v PaginationStatus)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetOAuthClientTokens200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTokens

`func (o *GetOAuthClientTokens200Response) GetTokens() []OAuthClientToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *GetOAuthClientTokens200Response) GetTokensOk() (*[]OAuthClientToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *GetOAuthClientTokens200Response) SetTokens(v []OAuthClientToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *GetOAuthClientTokens200Response) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### SetTokensNil

`func (o *GetOAuthClientTokens200Response) SetTokensNil(b bool)`

 SetTokensNil sets the value for Tokens to be an explicit nil

### UnsetTokens
`func (o *GetOAuthClientTokens200Response) UnsetTokens()`

UnsetTokens ensures that no value is present for Tokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


