package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/piusalfred/regsvc/pkg/service"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	Err error `json:"err"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		err := s.Register(ctx, req.Username, req.Email, req.Password)
		return RegisterResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.Err
}

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Token string `json:"token"`
	Err   error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := s.Login(ctx, req.Username, req.Password)
		return LoginResponse{
			Err:   err,
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// AddUserRequest collects the request parameters for the AddUser method.
type AddUserRequest struct {
	User service.User `json:"user"`
}

// AddUserResponse collects the response parameters for the AddUser method.
type AddUserResponse struct {
	Err error `json:"err"`
}

// MakeAddUserEndpoint returns an endpoint that invokes AddUser on the service.
func MakeAddUserEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddUserRequest)
		err := s.AddUser(ctx, req.User)
		return AddUserResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddUserResponse) Failed() error {
	return r.Err
}

// GetUserByIdRequest collects the request parameters for the GetUserById method.
type GetUserByIdRequest struct {
	Uuid string `json:"uuid"`
}

// GetUserByIdResponse collects the response parameters for the GetUserById method.
type GetUserByIdResponse struct {
	User service.User `json:"user"`
	Err  error        `json:"err"`
}

// MakeGetUserByIdEndpoint returns an endpoint that invokes GetUserById on the service.
func MakeGetUserByIdEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIdRequest)
		user, err := s.GetUserById(ctx, req.Uuid)
		return GetUserByIdResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserByIdResponse) Failed() error {
	return r.Err
}

// ListUsersRequest collects the request parameters for the ListUsers method.
type ListUsersRequest struct {
	Args map[string]interface{} `json:"args"`
}

// ListUsersResponse collects the response parameters for the ListUsers method.
type ListUsersResponse struct {
	Users []User `json:"users"`
	Err   error  `json:"err"`
}

// MakeListUsersEndpoint returns an endpoint that invokes ListUsers on the service.
func MakeListUsersEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListUsersRequest)
		users, err := s.ListUsers(ctx, req.Args)
		return ListUsersResponse{
			Err:   err,
			Users: users,
		}, nil
	}
}

// Failed implements Failer.
func (r ListUsersResponse) Failed() error {
	return r.Err
}

// DeleteUserRequest collects the request parameters for the DeleteUser method.
type DeleteUserRequest struct {
	Uuid string `json:"uuid"`
}

// DeleteUserResponse collects the response parameters for the DeleteUser method.
type DeleteUserResponse struct {
	Err error `json:"err"`
}

// MakeDeleteUserEndpoint returns an endpoint that invokes DeleteUser on the service.
func MakeDeleteUserEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		err := s.DeleteUser(ctx, req.Uuid)
		return DeleteUserResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserResponse) Failed() error {
	return r.Err
}

// UpdateUserRequest collects the request parameters for the UpdateUser method.
type UpdateUserRequest struct {
	Uuid string       `json:"uuid"`
	User service.User `json:"user"`
}

// UpdateUserResponse collects the response parameters for the UpdateUser method.
type UpdateUserResponse struct {
	Err error `json:"err"`
}

// MakeUpdateUserEndpoint returns an endpoint that invokes UpdateUser on the service.
func MakeUpdateUserEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := s.UpdateUser(ctx, req.Uuid, req.User)
		return UpdateUserResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserResponse) Failed() error {
	return r.Err
}

// AddNodeRequest collects the request parameters for the AddNode method.
type AddNodeRequest struct {
	Node service.Node `json:"node"`
}

// AddNodeResponse collects the response parameters for the AddNode method.
type AddNodeResponse struct {
	Err error `json:"err"`
}

// MakeAddNodeEndpoint returns an endpoint that invokes AddNode on the service.
func MakeAddNodeEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNodeRequest)
		err := s.AddNode(ctx, req.Node)
		return AddNodeResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddNodeResponse) Failed() error {
	return r.Err
}

// GetNodeByIdRequest collects the request parameters for the GetNodeById method.
type GetNodeByIdRequest struct {
	Uuid string `json:"uuid"`
}

// GetNodeByIdResponse collects the response parameters for the GetNodeById method.
type GetNodeByIdResponse struct {
	Node service.Node `json:"node"`
	Err  error        `json:"err"`
}

// MakeGetNodeByIdEndpoint returns an endpoint that invokes GetNodeById on the service.
func MakeGetNodeByIdEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNodeByIdRequest)
		node, err := s.GetNodeById(ctx, req.Uuid)
		return GetNodeByIdResponse{
			Err:  err,
			Node: node,
		}, nil
	}
}

// Failed implements Failer.
func (r GetNodeByIdResponse) Failed() error {
	return r.Err
}

// ListNodesRequest collects the request parameters for the ListNodes method.
type ListNodesRequest struct {
	Args map[string]interface{} `json:"args"`
}

// ListNodesResponse collects the response parameters for the ListNodes method.
type ListNodesResponse struct {
	Nodes []Node `json:"nodes"`
	Err   error  `json:"err"`
}

// MakeListNodesEndpoint returns an endpoint that invokes ListNodes on the service.
func MakeListNodesEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListNodesRequest)
		nodes, err := s.ListNodes(ctx, req.Args)
		return ListNodesResponse{
			Err:   err,
			Nodes: nodes,
		}, nil
	}
}

// Failed implements Failer.
func (r ListNodesResponse) Failed() error {
	return r.Err
}

// DeleteNodeRequest collects the request parameters for the DeleteNode method.
type DeleteNodeRequest struct {
	Uuid string `json:"uuid"`
}

// DeleteNodeResponse collects the response parameters for the DeleteNode method.
type DeleteNodeResponse struct {
	Err error `json:"err"`
}

// MakeDeleteNodeEndpoint returns an endpoint that invokes DeleteNode on the service.
func MakeDeleteNodeEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteNodeRequest)
		err := s.DeleteNode(ctx, req.Uuid)
		return DeleteNodeResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteNodeResponse) Failed() error {
	return r.Err
}

// UpdateNodeRequest collects the request parameters for the UpdateNode method.
type UpdateNodeRequest struct {
	Uuid string       `json:"uuid"`
	Node service.Node `json:"node"`
}

// UpdateNodeResponse collects the response parameters for the UpdateNode method.
type UpdateNodeResponse struct {
	Err error `json:"err"`
}

// MakeUpdateNodeEndpoint returns an endpoint that invokes UpdateNode on the service.
func MakeUpdateNodeEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateNodeRequest)
		err := s.UpdateNode(ctx, req.Uuid, req.Node)
		return UpdateNodeResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r UpdateNodeResponse) Failed() error {
	return r.Err
}

// AddRegionRequest collects the request parameters for the AddRegion method.
type AddRegionRequest struct {
	Region service.Region `json:"region"`
}

// AddRegionResponse collects the response parameters for the AddRegion method.
type AddRegionResponse struct {
	Err error `json:"err"`
}

// MakeAddRegionEndpoint returns an endpoint that invokes AddRegion on the service.
func MakeAddRegionEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRegionRequest)
		err := s.AddRegion(ctx, req.Region)
		return AddRegionResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r AddRegionResponse) Failed() error {
	return r.Err
}

// GetRegionByIdRequest collects the request parameters for the GetRegionById method.
type GetRegionByIdRequest struct {
	Uuid string `json:"uuid"`
}

// GetRegionByIdResponse collects the response parameters for the GetRegionById method.
type GetRegionByIdResponse struct {
	Region service.Region `json:"region"`
	Err    error          `json:"err"`
}

// MakeGetRegionByIdEndpoint returns an endpoint that invokes GetRegionById on the service.
func MakeGetRegionByIdEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRegionByIdRequest)
		region, err := s.GetRegionById(ctx, req.Uuid)
		return GetRegionByIdResponse{
			Err:    err,
			Region: region,
		}, nil
	}
}

// Failed implements Failer.
func (r GetRegionByIdResponse) Failed() error {
	return r.Err
}

// ListRegionsRequest collects the request parameters for the ListRegions method.
type ListRegionsRequest struct {
	Args map[string]interface{} `json:"args"`
}

// ListRegionsResponse collects the response parameters for the ListRegions method.
type ListRegionsResponse struct {
	Regions []Region `json:"regions"`
	Err     error    `json:"err"`
}

// MakeListRegionsEndpoint returns an endpoint that invokes ListRegions on the service.
func MakeListRegionsEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRegionsRequest)
		regions, err := s.ListRegions(ctx, req.Args)
		return ListRegionsResponse{
			Err:     err,
			Regions: regions,
		}, nil
	}
}

// Failed implements Failer.
func (r ListRegionsResponse) Failed() error {
	return r.Err
}

// DeleteRegionRequest collects the request parameters for the DeleteRegion method.
type DeleteRegionRequest struct {
	Uuid string `json:"uuid"`
}

// DeleteRegionResponse collects the response parameters for the DeleteRegion method.
type DeleteRegionResponse struct {
	Err error `json:"err"`
}

// MakeDeleteRegionEndpoint returns an endpoint that invokes DeleteRegion on the service.
func MakeDeleteRegionEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRegionRequest)
		err := s.DeleteRegion(ctx, req.Uuid)
		return DeleteRegionResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteRegionResponse) Failed() error {
	return r.Err
}

// UpdateRegionRequest collects the request parameters for the UpdateRegion method.
type UpdateRegionRequest struct {
	Uuid   string         `json:"uuid"`
	Region service.Region `json:"region"`
}

// UpdateRegionResponse collects the response parameters for the UpdateRegion method.
type UpdateRegionResponse struct {
	Err error `json:"err"`
}

// MakeUpdateRegionEndpoint returns an endpoint that invokes UpdateRegion on the service.
func MakeUpdateRegionEndpoint(s service.RegsvcService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRegionRequest)
		err := s.UpdateRegion(ctx, req.Uuid, req.Region)
		return UpdateRegionResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r UpdateRegionResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, username string, email string, password string) (err error) {
	request := RegisterRequest{
		Email:    email,
		Password: password,
		Username: username,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).Err
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, username string, password string) (token string, err error) {
	request := LoginRequest{
		Password: password,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Token, response.(LoginResponse).Err
}

// AddUser implements Service. Primarily useful in a client.
func (e Endpoints) AddUser(ctx context.Context, user service.User) (err error) {
	request := AddUserRequest{User: user}
	response, err := e.AddUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddUserResponse).Err
}

// GetUserById implements Service. Primarily useful in a client.
func (e Endpoints) GetUserById(ctx context.Context, uuid string) (user service.User, err error) {
	request := GetUserByIdRequest{Uuid: uuid}
	response, err := e.GetUserByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserByIdResponse).User, response.(GetUserByIdResponse).Err
}

// ListUsers implements Service. Primarily useful in a client.
func (e Endpoints) ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error) {
	request := ListUsersRequest{Args: args}
	response, err := e.ListUsersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListUsersResponse).Users, response.(ListUsersResponse).Err
}

// DeleteUser implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUser(ctx context.Context, uuid string) (err error) {
	request := DeleteUserRequest{Uuid: uuid}
	response, err := e.DeleteUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserResponse).Err
}

// UpdateUser implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUser(ctx context.Context, uuid string, user service.User) (err error) {
	request := UpdateUserRequest{
		User: user,
		Uuid: uuid,
	}
	response, err := e.UpdateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserResponse).Err
}

// AddNode implements Service. Primarily useful in a client.
func (e Endpoints) AddNode(ctx context.Context, node service.Node) (err error) {
	request := AddNodeRequest{Node: node}
	response, err := e.AddNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddNodeResponse).Err
}

// GetNodeById implements Service. Primarily useful in a client.
func (e Endpoints) GetNodeById(ctx context.Context, uuid string) (node service.Node, err error) {
	request := GetNodeByIdRequest{Uuid: uuid}
	response, err := e.GetNodeByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetNodeByIdResponse).Node, response.(GetNodeByIdResponse).Err
}

// ListNodes implements Service. Primarily useful in a client.
func (e Endpoints) ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error) {
	request := ListNodesRequest{Args: args}
	response, err := e.ListNodesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListNodesResponse).Nodes, response.(ListNodesResponse).Err
}

// DeleteNode implements Service. Primarily useful in a client.
func (e Endpoints) DeleteNode(ctx context.Context, uuid string) (err error) {
	request := DeleteNodeRequest{Uuid: uuid}
	response, err := e.DeleteNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteNodeResponse).Err
}

// UpdateNode implements Service. Primarily useful in a client.
func (e Endpoints) UpdateNode(ctx context.Context, uuid string, node service.Node) (err error) {
	request := UpdateNodeRequest{
		Node: node,
		Uuid: uuid,
	}
	response, err := e.UpdateNodeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateNodeResponse).Err
}

// AddRegion implements Service. Primarily useful in a client.
func (e Endpoints) AddRegion(ctx context.Context, region service.Region) (err error) {
	request := AddRegionRequest{Region: region}
	response, err := e.AddRegionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddRegionResponse).Err
}

// GetRegionById implements Service. Primarily useful in a client.
func (e Endpoints) GetRegionById(ctx context.Context, uuid string) (region service.Region, err error) {
	request := GetRegionByIdRequest{Uuid: uuid}
	response, err := e.GetRegionByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetRegionByIdResponse).Region, response.(GetRegionByIdResponse).Err
}

// ListRegions implements Service. Primarily useful in a client.
func (e Endpoints) ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error) {
	request := ListRegionsRequest{Args: args}
	response, err := e.ListRegionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ListRegionsResponse).Regions, response.(ListRegionsResponse).Err
}

// DeleteRegion implements Service. Primarily useful in a client.
func (e Endpoints) DeleteRegion(ctx context.Context, uuid string) (err error) {
	request := DeleteRegionRequest{Uuid: uuid}
	response, err := e.DeleteRegionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteRegionResponse).Err
}

// UpdateRegion implements Service. Primarily useful in a client.
func (e Endpoints) UpdateRegion(ctx context.Context, uuid string, region service.Region) (err error) {
	request := UpdateRegionRequest{
		Region: region,
		Uuid:   uuid,
	}
	response, err := e.UpdateRegionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateRegionResponse).Err
}
