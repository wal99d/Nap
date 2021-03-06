package nap

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

type RestResource struct {
	Endpoint string // like "/get"
	Method   string // like "GET"
	Router   *CBRouter
}

func NewResource(endpoint, method string, router *CBRouter) *RestResource {
	return &RestResource{
		Endpoint: endpoint,
		Method:   method,
		Router:   router,
	}
}

func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}
	t, err := template.New("resource").Parse(r.Endpoint)
	if err != nil {
		log.Fatalln("Unable to parse endpoint")
	}
	buffer := &bytes.Buffer{}
	t.Execute(buffer, params)
	endpoint, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalln("Unable to read endpoint")
	}
	return string(endpoint)
}
