# PaginationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNum** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageTotal** | Pointer to **int32** |  | [optional] 

## Methods

### NewPaginationStatus

`func NewPaginationStatus() *PaginationStatus`

NewPaginationStatus instantiates a new PaginationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationStatusWithDefaults

`func NewPaginationStatusWithDefaults() *PaginationStatus`

NewPaginationStatusWithDefaults instantiates a new PaginationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNum

`func (o *PaginationStatus) GetPageNum() int32`

GetPageNum returns the PageNum field if non-nil, zero value otherwise.

### GetPageNumOk

`func (o *PaginationStatus) GetPageNumOk() (*int32, bool)`

GetPageNumOk returns a tuple with the PageNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNum

`func (o *PaginationStatus) SetPageNum(v int32)`

SetPageNum sets PageNum field to given value.

### HasPageNum

`func (o *PaginationStatus) HasPageNum() bool`

HasPageNum returns a boolean if a field has been set.

### GetPageSize

`func (o *PaginationStatus) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginationStatus) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginationStatus) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PaginationStatus) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageTotal

`func (o *PaginationStatus) GetPageTotal() int32`

GetPageTotal returns the PageTotal field if non-nil, zero value otherwise.

### GetPageTotalOk

`func (o *PaginationStatus) GetPageTotalOk() (*int32, bool)`

GetPageTotalOk returns a tuple with the PageTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTotal

`func (o *PaginationStatus) SetPageTotal(v int32)`

SetPageTotal sets PageTotal field to given value.

### HasPageTotal

`func (o *PaginationStatus) HasPageTotal() bool`

HasPageTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


