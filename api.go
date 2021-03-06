package nap

import "fmt"

type API struct {
	BaseURL       string
	Resources     map[string]*RestResource
	DefaultRouter *CBRouter
	Client        *Client
}

func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:       baseURL,
		Resources:     make(map[string]*RestResource),
		DefaultRouter: NewRouter(),
		Client:        NewClient(),
	}
}

func (api *API) SetAuth(auth Authentication) {
	api.Client.SetAuth(auth)
}

func (api *API) AddResource(name string, res *RestResource) {
	api.Resources[name] = res
}

func (api *API) Call(name string, params map[string]string, payload interface{}) error {
	res, ok := api.Resources[name]
	if !ok {
		return fmt.Errorf("resource doesn't exist: %s", name)
	}
	if err := api.Client.ProcessRequest(api.BaseURL, res, params, payload); err != nil {
		return err
	}
	return nil
}

func (api *API) ResourceNames() []string {
	res := []string{}
	for k := range api.Resources {
		res = append(res, k)
	}
	return res
}
