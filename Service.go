package gravityforms

import (
	"encoding/base64"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

const (
	apiName string = "GravityForms"
	baseUrl string = "https://%s/wp-json/gf/v2"
)

type Service struct {
	domain      string
	username    string
	password    string
	httpService *go_http.Service
}

type ServiceConfig struct {
	Domain   string
	Username string
	Password string
}

func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		domain:      serviceConfig.Domain,
		username:    serviceConfig.Username,
		password:    serviceConfig.Password,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authorization token to header
	if requestConfig.NonDefaultHeaders == nil {
		requestConfig.NonDefaultHeaders = &http.Header{}
	}
	requestConfig.NonDefaultHeaders.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", service.username, service.password)))))

	return service.httpService.HttpRequest(requestConfig)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.username
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}
