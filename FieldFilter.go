package gravityforms

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type FieldFilter struct {
	Key       interface{} `json:"key"`
	Text      string      `json:"text"`
	Operators []string    `json:"operators"`
	Values    *[]struct {
		Text  string `json:"text"`
		Value string `json:"value"`
	} `json:"values"`
	PreventMultiple bool    `json:"preventMultiple"`
	Placeholder     *string `json:"placeholder"`
	CssClass        *string `json:"cssClass"`
}

func (service *Service) GetFieldFilters(formId int64) (*[]FieldFilter, *errortools.Error) {
	var fieldFilters []FieldFilter

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("forms/%v/field-filters", formId)),
		ResponseModel: &fieldFilters,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &fieldFilters, nil
}
