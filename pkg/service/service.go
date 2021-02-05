package service

import "context"

type User struct {
	UUID     string `json:"id,omitempty"`       //id or user token | uuid
	Name     string `json:"name"`               //fullname
	Email    string `json:"email"`              //email
	Password string `json:"password,omitempty"` //password of user
	Group    int    `json:"group,omitempty"`    //user group
	Region   string `json:"region,omitempty"`   //operating region in case of multi cloud
	Created  string `json:"created,omitempty"`  //when was this user added
}

//Node represent the edge node deployed in the igrid network
type Node struct {
	UUID    string `json:"uuid"`
	Addr    string `json:"addr"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Region  string `json:"region"`
	Latd    string `json:"latitude"`
	Long    string `json:"longitude"`
	Created string `json:"created"`
	Master  string `json:"master,omitempty"`
}

type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type UserService interface {
	//Register an run for the first time a user registers to use the service
	Register(ctx context.Context, username, email, password string) (err error)
	//Login is to obtain access token
	Login(ctx context.Context, username, password string) (token string, err error)
	AddUser(ctx context.Context, user User) (err error)
	GetUserById(ctx context.Context, uuid string) (user User, err error)
	ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error)
	DeleteUser(ctx context.Context, uuid string) (err error)
	UpdateUser(ctx context.Context, uuid string, user User) (err error)
}

type NodeService interface {
	AddNode(ctx context.Context, node Node) (err error)
	GetNodeById(ctx context.Context, uuid string) (node Node, err error)
	ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error)
	DeleteNode(ctx context.Context, uuid string) (err error)
	UpdateNode(ctx context.Context, uuid string, node Node) (err error)
}

type RegionService interface {
	AddRegion(ctx context.Context, region Region) (err error)
	GetRegionById(ctx context.Context, uuid string) (region Region, err error)
	ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error)
	DeleteRegion(ctx context.Context, uuid string) (err error)
	UpdateRegion(ctx context.Context, uuid string, region Region) (err error)
}

// RegsvcService describes the service.
type RegsvcService interface {

	UserService

	NodeService

	RegionService
}

type basicRegsvcService struct{}

func (b *basicRegsvcService) Register(ctx context.Context, username string, email string, password string) (err error) {
	// TODO implement the business logic of Register
	return err
}
func (b *basicRegsvcService) Login(ctx context.Context, username string, password string) (token string, err error) {
	// TODO implement the business logic of Login
	return token, err
}
func (b *basicRegsvcService) AddUser(ctx context.Context, user User) (err error) {
	// TODO implement the business logic of AddUser
	return err
}
func (b *basicRegsvcService) GetUserById(ctx context.Context, uuid string) (user User, err error) {
	// TODO implement the business logic of GetUserById
	return user, err
}
func (b *basicRegsvcService) ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error) {
	// TODO implement the business logic of ListUsers
	return users, err
}
func (b *basicRegsvcService) DeleteUser(ctx context.Context, uuid string) (err error) {
	// TODO implement the business logic of DeleteUser
	return err
}
func (b *basicRegsvcService) UpdateUser(ctx context.Context, uuid string, user User) (err error) {
	// TODO implement the business logic of UpdateUser
	return err
}
func (b *basicRegsvcService) AddNode(ctx context.Context, node Node) (err error) {
	// TODO implement the business logic of AddNode
	return err
}
func (b *basicRegsvcService) GetNodeById(ctx context.Context, uuid string) (node Node, err error) {
	// TODO implement the business logic of GetNodeById
	return node, err
}
func (b *basicRegsvcService) ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error) {
	// TODO implement the business logic of ListNodes
	return nodes, err
}
func (b *basicRegsvcService) DeleteNode(ctx context.Context, uuid string) (err error) {
	// TODO implement the business logic of DeleteNode
	return err
}
func (b *basicRegsvcService) UpdateNode(ctx context.Context, uuid string, node Node) (err error) {
	// TODO implement the business logic of UpdateNode
	return err
}
func (b *basicRegsvcService) AddRegion(ctx context.Context, region Region) (err error) {
	// TODO implement the business logic of AddRegion
	return err
}
func (b *basicRegsvcService) GetRegionById(ctx context.Context, uuid string) (region Region, err error) {
	// TODO implement the business logic of GetRegionById
	return region, err
}
func (b *basicRegsvcService) ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error) {
	// TODO implement the business logic of ListRegions
	return regions, err
}
func (b *basicRegsvcService) DeleteRegion(ctx context.Context, uuid string) (err error) {
	// TODO implement the business logic of DeleteRegion
	return err
}
func (b *basicRegsvcService) UpdateRegion(ctx context.Context, uuid string, region Region) (err error) {
	// TODO implement the business logic of UpdateRegion
	return err
}

// NewBasicRegsvcService returns a naive, stateless implementation of RegsvcService.
func NewBasicRegsvcService() RegsvcService {
	return &basicRegsvcService{}
}

// New returns a RegsvcService with all of the expected middleware wired in.
func New(middleware []Middleware) RegsvcService {
	var svc RegsvcService = NewBasicRegsvcService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

//todo
//INFO[0000] Do not forget to append your service middleware to your service middlewares
//INFO[0000] Add it to cmd/service/service.go#getServiceMiddleware()