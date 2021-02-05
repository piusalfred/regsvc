package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpoint1 "github.com/piusalfred/regsvc/pkg/endpoint"
	http2 "github.com/piusalfred/regsvc/pkg/http"
	service "github.com/piusalfred/regsvc/pkg/service"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.RegsvcService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var registerEndpoint endpoint.Endpoint
	{
		registerEndpoint = http.NewClient("POST", copyURL(u, "/register"), encodeHTTPGenericRequest, decodeRegisterResponse, options["Register"]...).Endpoint()
	}

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = http.NewClient("POST", copyURL(u, "/login"), encodeHTTPGenericRequest, decodeLoginResponse, options["Login"]...).Endpoint()
	}

	var addUserEndpoint endpoint.Endpoint
	{
		addUserEndpoint = http.NewClient("POST", copyURL(u, "/add-user"), encodeHTTPGenericRequest, decodeAddUserResponse, options["AddUser"]...).Endpoint()
	}

	var getUserByIdEndpoint endpoint.Endpoint
	{
		getUserByIdEndpoint = http.NewClient("POST", copyURL(u, "/get-user-by-id"), encodeHTTPGenericRequest, decodeGetUserByIdResponse, options["GetUserById"]...).Endpoint()
	}

	var listUsersEndpoint endpoint.Endpoint
	{
		listUsersEndpoint = http.NewClient("POST", copyURL(u, "/list-users"), encodeHTTPGenericRequest, decodeListUsersResponse, options["ListUsers"]...).Endpoint()
	}

	var deleteUserEndpoint endpoint.Endpoint
	{
		deleteUserEndpoint = http.NewClient("POST", copyURL(u, "/delete-user"), encodeHTTPGenericRequest, decodeDeleteUserResponse, options["DeleteUser"]...).Endpoint()
	}

	var updateUserEndpoint endpoint.Endpoint
	{
		updateUserEndpoint = http.NewClient("POST", copyURL(u, "/update-user"), encodeHTTPGenericRequest, decodeUpdateUserResponse, options["UpdateUser"]...).Endpoint()
	}

	var addNodeEndpoint endpoint.Endpoint
	{
		addNodeEndpoint = http.NewClient("POST", copyURL(u, "/add-node"), encodeHTTPGenericRequest, decodeAddNodeResponse, options["AddNode"]...).Endpoint()
	}

	var getNodeByIdEndpoint endpoint.Endpoint
	{
		getNodeByIdEndpoint = http.NewClient("POST", copyURL(u, "/get-node-by-id"), encodeHTTPGenericRequest, decodeGetNodeByIdResponse, options["GetNodeById"]...).Endpoint()
	}

	var listNodesEndpoint endpoint.Endpoint
	{
		listNodesEndpoint = http.NewClient("POST", copyURL(u, "/list-nodes"), encodeHTTPGenericRequest, decodeListNodesResponse, options["ListNodes"]...).Endpoint()
	}

	var deleteNodeEndpoint endpoint.Endpoint
	{
		deleteNodeEndpoint = http.NewClient("POST", copyURL(u, "/delete-node"), encodeHTTPGenericRequest, decodeDeleteNodeResponse, options["DeleteNode"]...).Endpoint()
	}

	var updateNodeEndpoint endpoint.Endpoint
	{
		updateNodeEndpoint = http.NewClient("POST", copyURL(u, "/update-node"), encodeHTTPGenericRequest, decodeUpdateNodeResponse, options["UpdateNode"]...).Endpoint()
	}

	var addRegionEndpoint endpoint.Endpoint
	{
		addRegionEndpoint = http.NewClient("POST", copyURL(u, "/add-region"), encodeHTTPGenericRequest, decodeAddRegionResponse, options["AddRegion"]...).Endpoint()
	}

	var getRegionByIdEndpoint endpoint.Endpoint
	{
		getRegionByIdEndpoint = http.NewClient("POST", copyURL(u, "/get-region-by-id"), encodeHTTPGenericRequest, decodeGetRegionByIdResponse, options["GetRegionById"]...).Endpoint()
	}

	var listRegionsEndpoint endpoint.Endpoint
	{
		listRegionsEndpoint = http.NewClient("POST", copyURL(u, "/list-regions"), encodeHTTPGenericRequest, decodeListRegionsResponse, options["ListRegions"]...).Endpoint()
	}

	var deleteRegionEndpoint endpoint.Endpoint
	{
		deleteRegionEndpoint = http.NewClient("POST", copyURL(u, "/delete-region"), encodeHTTPGenericRequest, decodeDeleteRegionResponse, options["DeleteRegion"]...).Endpoint()
	}

	var updateRegionEndpoint endpoint.Endpoint
	{
		updateRegionEndpoint = http.NewClient("POST", copyURL(u, "/update-region"), encodeHTTPGenericRequest, decodeUpdateRegionResponse, options["UpdateRegion"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		AddNodeEndpoint:       addNodeEndpoint,
		AddRegionEndpoint:     addRegionEndpoint,
		AddUserEndpoint:       addUserEndpoint,
		DeleteNodeEndpoint:    deleteNodeEndpoint,
		DeleteRegionEndpoint:  deleteRegionEndpoint,
		DeleteUserEndpoint:    deleteUserEndpoint,
		GetNodeByIdEndpoint:   getNodeByIdEndpoint,
		GetRegionByIdEndpoint: getRegionByIdEndpoint,
		GetUserByIdEndpoint:   getUserByIdEndpoint,
		ListNodesEndpoint:     listNodesEndpoint,
		ListRegionsEndpoint:   listRegionsEndpoint,
		ListUsersEndpoint:     listUsersEndpoint,
		LoginEndpoint:         loginEndpoint,
		RegisterEndpoint:      registerEndpoint,
		UpdateNodeEndpoint:    updateNodeEndpoint,
		UpdateRegionEndpoint:  updateRegionEndpoint,
		UpdateUserEndpoint:    updateUserEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeRegisterResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeRegisterResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.RegisterResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeLoginResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeLoginResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.LoginResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListUsersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListUsersResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListUsersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetNodeByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetNodeByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetNodeByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListNodesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListNodesResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListNodesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateNodeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateNodeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateNodeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddRegionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddRegionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddRegionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetRegionByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetRegionByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetRegionByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeListRegionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeListRegionsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ListRegionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteRegionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteRegionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteRegionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateRegionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateRegionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateRegionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
