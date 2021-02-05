// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/piusalfred/regsvc/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	RegisterEndpoint      endpoint.Endpoint
	LoginEndpoint         endpoint.Endpoint
	AddUserEndpoint       endpoint.Endpoint
	GetUserByIdEndpoint   endpoint.Endpoint
	ListUsersEndpoint     endpoint.Endpoint
	DeleteUserEndpoint    endpoint.Endpoint
	UpdateUserEndpoint    endpoint.Endpoint
	AddNodeEndpoint       endpoint.Endpoint
	GetNodeByIdEndpoint   endpoint.Endpoint
	ListNodesEndpoint     endpoint.Endpoint
	DeleteNodeEndpoint    endpoint.Endpoint
	UpdateNodeEndpoint    endpoint.Endpoint
	AddRegionEndpoint     endpoint.Endpoint
	GetRegionByIdEndpoint endpoint.Endpoint
	ListRegionsEndpoint   endpoint.Endpoint
	DeleteRegionEndpoint  endpoint.Endpoint
	UpdateRegionEndpoint  endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.RegsvcService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddNodeEndpoint:       MakeAddNodeEndpoint(s),
		AddRegionEndpoint:     MakeAddRegionEndpoint(s),
		AddUserEndpoint:       MakeAddUserEndpoint(s),
		DeleteNodeEndpoint:    MakeDeleteNodeEndpoint(s),
		DeleteRegionEndpoint:  MakeDeleteRegionEndpoint(s),
		DeleteUserEndpoint:    MakeDeleteUserEndpoint(s),
		GetNodeByIdEndpoint:   MakeGetNodeByIdEndpoint(s),
		GetRegionByIdEndpoint: MakeGetRegionByIdEndpoint(s),
		GetUserByIdEndpoint:   MakeGetUserByIdEndpoint(s),
		ListNodesEndpoint:     MakeListNodesEndpoint(s),
		ListRegionsEndpoint:   MakeListRegionsEndpoint(s),
		ListUsersEndpoint:     MakeListUsersEndpoint(s),
		LoginEndpoint:         MakeLoginEndpoint(s),
		RegisterEndpoint:      MakeRegisterEndpoint(s),
		UpdateNodeEndpoint:    MakeUpdateNodeEndpoint(s),
		UpdateRegionEndpoint:  MakeUpdateRegionEndpoint(s),
		UpdateUserEndpoint:    MakeUpdateUserEndpoint(s),
	}
	for _, m := range mdw["Register"] {
		eps.RegisterEndpoint = m(eps.RegisterEndpoint)
	}
	for _, m := range mdw["Login"] {
		eps.LoginEndpoint = m(eps.LoginEndpoint)
	}
	for _, m := range mdw["AddUser"] {
		eps.AddUserEndpoint = m(eps.AddUserEndpoint)
	}
	for _, m := range mdw["GetUserById"] {
		eps.GetUserByIdEndpoint = m(eps.GetUserByIdEndpoint)
	}
	for _, m := range mdw["ListUsers"] {
		eps.ListUsersEndpoint = m(eps.ListUsersEndpoint)
	}
	for _, m := range mdw["DeleteUser"] {
		eps.DeleteUserEndpoint = m(eps.DeleteUserEndpoint)
	}
	for _, m := range mdw["UpdateUser"] {
		eps.UpdateUserEndpoint = m(eps.UpdateUserEndpoint)
	}
	for _, m := range mdw["AddNode"] {
		eps.AddNodeEndpoint = m(eps.AddNodeEndpoint)
	}
	for _, m := range mdw["GetNodeById"] {
		eps.GetNodeByIdEndpoint = m(eps.GetNodeByIdEndpoint)
	}
	for _, m := range mdw["ListNodes"] {
		eps.ListNodesEndpoint = m(eps.ListNodesEndpoint)
	}
	for _, m := range mdw["DeleteNode"] {
		eps.DeleteNodeEndpoint = m(eps.DeleteNodeEndpoint)
	}
	for _, m := range mdw["UpdateNode"] {
		eps.UpdateNodeEndpoint = m(eps.UpdateNodeEndpoint)
	}
	for _, m := range mdw["AddRegion"] {
		eps.AddRegionEndpoint = m(eps.AddRegionEndpoint)
	}
	for _, m := range mdw["GetRegionById"] {
		eps.GetRegionByIdEndpoint = m(eps.GetRegionByIdEndpoint)
	}
	for _, m := range mdw["ListRegions"] {
		eps.ListRegionsEndpoint = m(eps.ListRegionsEndpoint)
	}
	for _, m := range mdw["DeleteRegion"] {
		eps.DeleteRegionEndpoint = m(eps.DeleteRegionEndpoint)
	}
	for _, m := range mdw["UpdateRegion"] {
		eps.UpdateRegionEndpoint = m(eps.UpdateRegionEndpoint)
	}
	return eps
}