# PolicyUpdateObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPolicies** | Pointer to **[]string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyUpdateObj

`func NewPolicyUpdateObj() *PolicyUpdateObj`

NewPolicyUpdateObj instantiates a new PolicyUpdateObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateObjWithDefaults

`func NewPolicyUpdateObjWithDefaults() *PolicyUpdateObj`

NewPolicyUpdateObjWithDefaults instantiates a new PolicyUpdateObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPolicies

`func (o *PolicyUpdateObj) GetApplyPolicies() []string`

GetApplyPolicies returns the ApplyPolicies field if non-nil, zero value otherwise.

### GetApplyPoliciesOk

`func (o *PolicyUpdateObj) GetApplyPoliciesOk() (*[]string, bool)`

GetApplyPoliciesOk returns a tuple with the ApplyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPolicies

`func (o *PolicyUpdateObj) SetApplyPolicies(v []string)`

SetApplyPolicies sets ApplyPolicies field to given value.

### HasApplyPolicies

`func (o *PolicyUpdateObj) HasApplyPolicies() bool`

HasApplyPolicies returns a boolean if a field has been set.

### SetApplyPoliciesNil

`func (o *PolicyUpdateObj) SetApplyPoliciesNil(b bool)`

 SetApplyPoliciesNil sets the value for ApplyPolicies to be an explicit nil

### UnsetApplyPolicies
`func (o *PolicyUpdateObj) UnsetApplyPolicies()`

UnsetApplyPolicies ensures that no value is present for ApplyPolicies, not even an explicit nil
### GetPolicy

`func (o *PolicyUpdateObj) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyUpdateObj) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyUpdateObj) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyUpdateObj) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


