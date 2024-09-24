# SourcesResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]SourceRecord**](SourceRecord.md) |  | [optional] 

## Methods

### NewSourcesResp

`func NewSourcesResp() *SourcesResp`

NewSourcesResp instantiates a new SourcesResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesRespWithDefaults

`func NewSourcesRespWithDefaults() *SourcesResp`

NewSourcesRespWithDefaults instantiates a new SourcesResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *SourcesResp) GetSources() []SourceRecord`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SourcesResp) GetSourcesOk() (*[]SourceRecord, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SourcesResp) SetSources(v []SourceRecord)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SourcesResp) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


