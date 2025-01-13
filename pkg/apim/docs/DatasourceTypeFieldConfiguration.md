# DatasourceTypeFieldConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | Pointer to [**DatasourceSourceConfig**](DatasourceSourceConfig.md) |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**Mapping** | Pointer to [**DatasourceMappingConfiguration**](DatasourceMappingConfiguration.md) |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 

## Methods

### NewDatasourceTypeFieldConfiguration

`func NewDatasourceTypeFieldConfiguration() *DatasourceTypeFieldConfiguration`

NewDatasourceTypeFieldConfiguration instantiates a new DatasourceTypeFieldConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceTypeFieldConfigurationWithDefaults

`func NewDatasourceTypeFieldConfigurationWithDefaults() *DatasourceTypeFieldConfiguration`

NewDatasourceTypeFieldConfigurationWithDefaults instantiates a new DatasourceTypeFieldConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *DatasourceTypeFieldConfiguration) GetDataSource() DatasourceSourceConfig`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DatasourceTypeFieldConfiguration) GetDataSourceOk() (*DatasourceSourceConfig, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DatasourceTypeFieldConfiguration) SetDataSource(v DatasourceSourceConfig)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DatasourceTypeFieldConfiguration) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetFieldName

`func (o *DatasourceTypeFieldConfiguration) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *DatasourceTypeFieldConfiguration) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *DatasourceTypeFieldConfiguration) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *DatasourceTypeFieldConfiguration) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetMapping

`func (o *DatasourceTypeFieldConfiguration) GetMapping() DatasourceMappingConfiguration`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *DatasourceTypeFieldConfiguration) GetMappingOk() (*DatasourceMappingConfiguration, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *DatasourceTypeFieldConfiguration) SetMapping(v DatasourceMappingConfiguration)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *DatasourceTypeFieldConfiguration) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetTypeName

`func (o *DatasourceTypeFieldConfiguration) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *DatasourceTypeFieldConfiguration) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *DatasourceTypeFieldConfiguration) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *DatasourceTypeFieldConfiguration) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


