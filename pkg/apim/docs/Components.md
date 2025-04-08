# Components

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **map[string]interface{}** |  | [optional] 
**Responses** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**Examples** | Pointer to **map[string]interface{}** |  | [optional] 
**RequestBodies** | Pointer to **map[string]interface{}** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 
**SecuritySchemes** | Pointer to **map[string]interface{}** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Callbacks** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewComponents

`func NewComponents() *Components`

NewComponents instantiates a new Components object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentsWithDefaults

`func NewComponentsWithDefaults() *Components`

NewComponentsWithDefaults instantiates a new Components object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Components) GetSchemas() map[string]interface{}`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Components) GetSchemasOk() (*map[string]interface{}, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Components) SetSchemas(v map[string]interface{})`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *Components) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetResponses

`func (o *Components) GetResponses() map[string]interface{}`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *Components) GetResponsesOk() (*map[string]interface{}, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *Components) SetResponses(v map[string]interface{})`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *Components) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetParameters

`func (o *Components) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Components) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Components) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Components) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetExamples

`func (o *Components) GetExamples() map[string]interface{}`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *Components) GetExamplesOk() (*map[string]interface{}, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *Components) SetExamples(v map[string]interface{})`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *Components) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetRequestBodies

`func (o *Components) GetRequestBodies() map[string]interface{}`

GetRequestBodies returns the RequestBodies field if non-nil, zero value otherwise.

### GetRequestBodiesOk

`func (o *Components) GetRequestBodiesOk() (*map[string]interface{}, bool)`

GetRequestBodiesOk returns a tuple with the RequestBodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodies

`func (o *Components) SetRequestBodies(v map[string]interface{})`

SetRequestBodies sets RequestBodies field to given value.

### HasRequestBodies

`func (o *Components) HasRequestBodies() bool`

HasRequestBodies returns a boolean if a field has been set.

### GetHeaders

`func (o *Components) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Components) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Components) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Components) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetSecuritySchemes

`func (o *Components) GetSecuritySchemes() map[string]interface{}`

GetSecuritySchemes returns the SecuritySchemes field if non-nil, zero value otherwise.

### GetSecuritySchemesOk

`func (o *Components) GetSecuritySchemesOk() (*map[string]interface{}, bool)`

GetSecuritySchemesOk returns a tuple with the SecuritySchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySchemes

`func (o *Components) SetSecuritySchemes(v map[string]interface{})`

SetSecuritySchemes sets SecuritySchemes field to given value.

### HasSecuritySchemes

`func (o *Components) HasSecuritySchemes() bool`

HasSecuritySchemes returns a boolean if a field has been set.

### GetLinks

`func (o *Components) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Components) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Components) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Components) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCallbacks

`func (o *Components) GetCallbacks() map[string]interface{}`

GetCallbacks returns the Callbacks field if non-nil, zero value otherwise.

### GetCallbacksOk

`func (o *Components) GetCallbacksOk() (*map[string]interface{}, bool)`

GetCallbacksOk returns a tuple with the Callbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbacks

`func (o *Components) SetCallbacks(v map[string]interface{})`

SetCallbacks sets Callbacks field to given value.

### HasCallbacks

`func (o *Components) HasCallbacks() bool`

HasCallbacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


