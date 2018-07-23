package dadata

type SuggestRequestParamsLocation struct {
	CityFiasID    string `json:"city_fias_id,omitempty"` // search only in this area
	City          string `json:"city,omitempty"`
	CityKladrID   string `json:"city_kladr_id,omitempty"`
	FiasID        string `json:"fias_id,omitempty"`
	KladrID       string `json:"kladr_id,omitempty"`
	Region        string `json:"region,omitempty"`
	RegionFiasID  string `json:"region_fias_id,omitempty"`
	RegionKladrID string `json:"region_kladr_id,omitempty"`
}

// SuggestRequestParams Request struct
type SuggestRequestParams struct {
	Query         string                         `json:"query"` // user input for suggestion
	Count         int                            `json:"count"` // ligmit for results
	Locations     []SuggestRequestParamsLocation `json:"locations"`
	RestrictValue bool                           `json:"restrict_value"` // don't show restricts (region) on results
}

// SuggestAddressResponse result slice for address suggestions
type SuggestAddressResponse struct {
	Suggestions []ResponseAddress `"json":"suggestions"`
}

// SuggestNameResponse result slice for name suggestions
type SuggestNameResponse struct {
	Suggestions []ResponseName `"json":"suggestions"`
}

// SuggestBankResponse result slice for bank suggestions
type SuggestBankResponse struct {
	Suggestions []ResponseBank `"json":"suggestions"`
}

// SuggestPartyResponse result slice for party suggestions
type SuggestPartyResponse struct {
	Suggestions []ResponseParty `"json":"suggestions"`
}

// SuggestEmailResponse result slice for email suggestions
type SuggestEmailResponse struct {
	Suggestions []ResponseEmail `"json":"suggestions"`
}

func (daData *DaData) sendSuggestRequest(lastURLPart string, requestParams SuggestRequestParams, result interface{}) error {
	return daData.sendRequest("suggest/"+lastURLPart, requestParams, result)
}

// SuggestAddresses try to return suggest addresses by requestParams
func (daData *DaData) SuggestAddresses(requestParams SuggestRequestParams) ([]ResponseAddress, error) {

	result := SuggestAddressResponse{}
	if err := daData.sendSuggestRequest("address", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestNames try to return suggest names by requestParams
func (daData *DaData) SuggestNames(requestParams SuggestRequestParams) ([]ResponseName, error) {

	result := SuggestNameResponse{}
	if err := daData.sendSuggestRequest("fio", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestBanks try to return suggest banks by requestParams
func (daData *DaData) SuggestBanks(requestParams SuggestRequestParams) ([]ResponseBank, error) {

	result := SuggestBankResponse{}
	if err := daData.sendSuggestRequest("bank", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestParties try to return suggest parties by requestParams
func (daData *DaData) SuggestParties(requestParams SuggestRequestParams) ([]ResponseParty, error) {

	result := SuggestPartyResponse{}
	if err := daData.sendSuggestRequest("party", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

// SuggestEmails try to return suggest emails by requestParams
func (daData *DaData) SuggestEmails(requestParams SuggestRequestParams) ([]ResponseEmail, error) {

	result := SuggestEmailResponse{}
	if err := daData.sendSuggestRequest("email", requestParams, &result); err != nil {
		return nil, err
	}

	return result.Suggestions, nil
}

/*
{
"city":"",
"city_fias_id":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5",
"city_kladr_id":"7700000000000",
"city_type":"г",
"city_type_full":"город",
"city_with_type":"г Москва",
"country":"Россия",
"fias_actuality_state":"0",
"fias_code":"77000000000000000000000",
"fias_id":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5",
"fias_level":"1",
"geo_lat":"55.7540471",
"geo_lon":"37.620405",
"kladr_id":"7700000000000",
"okato":"45000000000",
"oktmo":"45000000",
"postal_code":"101000",
"qc_geo":"4",
"region":"Москва",
"region_fias_id":"0c5b2444-70a0-4932-980c-b4dc0d3f02b5",
"region_kladr_id":"7700000000000",
"region_type":"г",
"region_type_full":"город",
"region_with_type":"г Москва",
"source":"Москва",
"tax_office":"7700",
"tax_office_legal":"7700",
"timezone":"UTC+3",
}
*/
