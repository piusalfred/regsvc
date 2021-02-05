package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(RegsvcService) RegsvcService

type loggingMiddleware struct {
	logger log.Logger
	next   RegsvcService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a RegsvcService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next RegsvcService) RegsvcService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Register(ctx context.Context, username string, email string, password string) (err error) {
	defer func() {
		l.logger.Log("method", "Register", "username", username, "email", email, "password", password, "err", err)
	}()
	return l.next.Register(ctx, username, email, password)
}
func (l loggingMiddleware) Login(ctx context.Context, username string, password string) (token string, err error) {
	defer func() {
		l.logger.Log("method", "Login", "username", username, "password", password, "token", token, "err", err)
	}()
	return l.next.Login(ctx, username, password)
}
func (l loggingMiddleware) AddUser(ctx context.Context, user User) (err error) {
	defer func() {
		l.logger.Log("method", "AddUser", "user", user, "err", err)
	}()
	return l.next.AddUser(ctx, user)
}
func (l loggingMiddleware) GetUserById(ctx context.Context, uuid string) (user User, err error) {
	defer func() {
		l.logger.Log("method", "GetUserById", "uuid", uuid, "user", user, "err", err)
	}()
	return l.next.GetUserById(ctx, uuid)
}
func (l loggingMiddleware) ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error) {
	defer func() {
		l.logger.Log("method", "ListUsers", "args", args, "users", users, "err", err)
	}()
	return l.next.ListUsers(ctx, args)
}
func (l loggingMiddleware) DeleteUser(ctx context.Context, uuid string) (err error) {
	defer func() {
		l.logger.Log("method", "DeleteUser", "uuid", uuid, "err", err)
	}()
	return l.next.DeleteUser(ctx, uuid)
}
func (l loggingMiddleware) UpdateUser(ctx context.Context, uuid string, user User) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateUser", "uuid", uuid, "user", user, "err", err)
	}()
	return l.next.UpdateUser(ctx, uuid, user)
}
func (l loggingMiddleware) AddNode(ctx context.Context, node Node) (err error) {
	defer func() {
		l.logger.Log("method", "AddNode", "node", node, "err", err)
	}()
	return l.next.AddNode(ctx, node)
}
func (l loggingMiddleware) GetNodeById(ctx context.Context, uuid string) (node Node, err error) {
	defer func() {
		l.logger.Log("method", "GetNodeById", "uuid", uuid, "node", node, "err", err)
	}()
	return l.next.GetNodeById(ctx, uuid)
}
func (l loggingMiddleware) ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error) {
	defer func() {
		l.logger.Log("method", "ListNodes", "args", args, "nodes", nodes, "err", err)
	}()
	return l.next.ListNodes(ctx, args)
}
func (l loggingMiddleware) DeleteNode(ctx context.Context, uuid string) (err error) {
	defer func() {
		l.logger.Log("method", "DeleteNode", "uuid", uuid, "err", err)
	}()
	return l.next.DeleteNode(ctx, uuid)
}
func (l loggingMiddleware) UpdateNode(ctx context.Context, uuid string, node Node) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateNode", "uuid", uuid, "node", node, "err", err)
	}()
	return l.next.UpdateNode(ctx, uuid, node)
}
func (l loggingMiddleware) AddRegion(ctx context.Context, region Region) (err error) {
	defer func() {
		l.logger.Log("method", "AddRegion", "region", region, "err", err)
	}()
	return l.next.AddRegion(ctx, region)
}
func (l loggingMiddleware) GetRegionById(ctx context.Context, uuid string) (region Region, err error) {
	defer func() {
		l.logger.Log("method", "GetRegionById", "uuid", uuid, "region", region, "err", err)
	}()
	return l.next.GetRegionById(ctx, uuid)
}
func (l loggingMiddleware) ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error) {
	defer func() {
		l.logger.Log("method", "ListRegions", "args", args, "regions", regions, "err", err)
	}()
	return l.next.ListRegions(ctx, args)
}
func (l loggingMiddleware) DeleteRegion(ctx context.Context, uuid string) (err error) {
	defer func() {
		l.logger.Log("method", "DeleteRegion", "uuid", uuid, "err", err)
	}()
	return l.next.DeleteRegion(ctx, uuid)
}
func (l loggingMiddleware) UpdateRegion(ctx context.Context, uuid string, region Region) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateRegion", "uuid", uuid, "region", region, "err", err)
	}()
	return l.next.UpdateRegion(ctx, uuid, region)
}

type authnMiddleware struct {
	next RegsvcService
}

// AuthnMiddleware returns a RegsvcService Middleware.
func AuthnMiddleware() Middleware {
	return func(next RegsvcService) RegsvcService {
		return &authnMiddleware{next}
	}

}
func (a authnMiddleware) Register(ctx context.Context, username string, email string, password string) (err error) {
	// Implement your middleware logic here

	return a.next.Register(ctx, username, email, password)
}
func (a authnMiddleware) Login(ctx context.Context, username string, password string) (token string, err error) {
	// Implement your middleware logic here

	return a.next.Login(ctx, username, password)
}
func (a authnMiddleware) AddUser(ctx context.Context, user User) (err error) {
	// Implement your middleware logic here

	return a.next.AddUser(ctx, user)
}
func (a authnMiddleware) GetUserById(ctx context.Context, uuid string) (user User, err error) {
	// Implement your middleware logic here

	return a.next.GetUserById(ctx, uuid)
}
func (a authnMiddleware) ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error) {
	// Implement your middleware logic here

	return a.next.ListUsers(ctx, args)
}
func (a authnMiddleware) DeleteUser(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return a.next.DeleteUser(ctx, uuid)
}
func (a authnMiddleware) UpdateUser(ctx context.Context, uuid string, user User) (err error) {
	// Implement your middleware logic here

	return a.next.UpdateUser(ctx, uuid, user)
}
func (a authnMiddleware) AddNode(ctx context.Context, node Node) (err error) {
	// Implement your middleware logic here

	return a.next.AddNode(ctx, node)
}
func (a authnMiddleware) GetNodeById(ctx context.Context, uuid string) (node Node, err error) {
	// Implement your middleware logic here

	return a.next.GetNodeById(ctx, uuid)
}
func (a authnMiddleware) ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error) {
	// Implement your middleware logic here

	return a.next.ListNodes(ctx, args)
}
func (a authnMiddleware) DeleteNode(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return a.next.DeleteNode(ctx, uuid)
}
func (a authnMiddleware) UpdateNode(ctx context.Context, uuid string, node Node) (err error) {
	// Implement your middleware logic here

	return a.next.UpdateNode(ctx, uuid, node)
}
func (a authnMiddleware) AddRegion(ctx context.Context, region Region) (err error) {
	// Implement your middleware logic here

	return a.next.AddRegion(ctx, region)
}
func (a authnMiddleware) GetRegionById(ctx context.Context, uuid string) (region Region, err error) {
	// Implement your middleware logic here

	return a.next.GetRegionById(ctx, uuid)
}
func (a authnMiddleware) ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error) {
	// Implement your middleware logic here

	return a.next.ListRegions(ctx, args)
}
func (a authnMiddleware) DeleteRegion(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return a.next.DeleteRegion(ctx, uuid)
}
func (a authnMiddleware) UpdateRegion(ctx context.Context, uuid string, region Region) (err error) {
	// Implement your middleware logic here

	return a.next.UpdateRegion(ctx, uuid, region)
}

type authzMiddleware struct {
	next RegsvcService
}

// AuthzMiddleware returns a RegsvcService Middleware.
func AuthzMiddleware() Middleware {
	return func(next RegsvcService) RegsvcService {
		return &authzMiddleware{next}
	}

}
func (a authzMiddleware) Register(ctx context.Context, username string, email string, password string) (err error) {
	// Implement your middleware logic here

	return a.next.Register(ctx, username, email, password)
}
func (a authzMiddleware) Login(ctx context.Context, username string, password string) (token string, err error) {
	// Implement your middleware logic here

	return a.next.Login(ctx, username, password)
}
func (a authzMiddleware) AddUser(ctx context.Context, user User) (err error) {
	// Implement your middleware logic here

	return a.next.AddUser(ctx, user)
}
func (a authzMiddleware) GetUserById(ctx context.Context, uuid string) (user User, err error) {
	// Implement your middleware logic here

	return a.next.GetUserById(ctx, uuid)
}
func (a authzMiddleware) ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error) {
	// Implement your middleware logic here

	return a.next.ListUsers(ctx, args)
}
func (a authzMiddleware) DeleteUser(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return a.next.DeleteUser(ctx, uuid)
}
func (a authzMiddleware) UpdateUser(ctx context.Context, uuid string, user User) (err error) {
	// Implement your middleware logic here

	return a.next.UpdateUser(ctx, uuid, user)
}
func (a authzMiddleware) AddNode(ctx context.Context, node Node) (err error) {
	// Implement your middleware logic here

	return a.next.AddNode(ctx, node)
}
func (a authzMiddleware) GetNodeById(ctx context.Context, uuid string) (node Node, err error) {
	// Implement your middleware logic here

	return a.next.GetNodeById(ctx, uuid)
}
func (a authzMiddleware) ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error) {
	// Implement your middleware logic here

	return a.next.ListNodes(ctx, args)
}
func (a authzMiddleware) DeleteNode(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return a.next.DeleteNode(ctx, uuid)
}
func (a authzMiddleware) UpdateNode(ctx context.Context, uuid string, node Node) (err error) {
	// Implement your middleware logic here

	return a.next.UpdateNode(ctx, uuid, node)
}
func (a authzMiddleware) AddRegion(ctx context.Context, region Region) (err error) {
	// Implement your middleware logic here

	return a.next.AddRegion(ctx, region)
}
func (a authzMiddleware) GetRegionById(ctx context.Context, uuid string) (region Region, err error) {
	// Implement your middleware logic here

	return a.next.GetRegionById(ctx, uuid)
}
func (a authzMiddleware) ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error) {
	// Implement your middleware logic here

	return a.next.ListRegions(ctx, args)
}
func (a authzMiddleware) DeleteRegion(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return a.next.DeleteRegion(ctx, uuid)
}
func (a authzMiddleware) UpdateRegion(ctx context.Context, uuid string, region Region) (err error) {
	// Implement your middleware logic here

	return a.next.UpdateRegion(ctx, uuid, region)
}

type eventsMiddleware struct {
	next RegsvcService
}

// EventsMiddleware returns a RegsvcService Middleware.
func EventsMiddleware() Middleware {
	return func(next RegsvcService) RegsvcService {
		return &eventsMiddleware{next}
	}

}
func (e eventsMiddleware) Register(ctx context.Context, username string, email string, password string) (err error) {
	// Implement your middleware logic here

	return e.next.Register(ctx, username, email, password)
}
func (e eventsMiddleware) Login(ctx context.Context, username string, password string) (token string, err error) {
	// Implement your middleware logic here

	return e.next.Login(ctx, username, password)
}
func (e eventsMiddleware) AddUser(ctx context.Context, user User) (err error) {
	// Implement your middleware logic here

	return e.next.AddUser(ctx, user)
}
func (e eventsMiddleware) GetUserById(ctx context.Context, uuid string) (user User, err error) {
	// Implement your middleware logic here

	return e.next.GetUserById(ctx, uuid)
}
func (e eventsMiddleware) ListUsers(ctx context.Context, args map[string]interface{}) (users []User, err error) {
	// Implement your middleware logic here

	return e.next.ListUsers(ctx, args)
}
func (e eventsMiddleware) DeleteUser(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return e.next.DeleteUser(ctx, uuid)
}
func (e eventsMiddleware) UpdateUser(ctx context.Context, uuid string, user User) (err error) {
	// Implement your middleware logic here

	return e.next.UpdateUser(ctx, uuid, user)
}
func (e eventsMiddleware) AddNode(ctx context.Context, node Node) (err error) {
	// Implement your middleware logic here

	return e.next.AddNode(ctx, node)
}
func (e eventsMiddleware) GetNodeById(ctx context.Context, uuid string) (node Node, err error) {
	// Implement your middleware logic here

	return e.next.GetNodeById(ctx, uuid)
}
func (e eventsMiddleware) ListNodes(ctx context.Context, args map[string]interface{}) (nodes []Node, err error) {
	// Implement your middleware logic here

	return e.next.ListNodes(ctx, args)
}
func (e eventsMiddleware) DeleteNode(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return e.next.DeleteNode(ctx, uuid)
}
func (e eventsMiddleware) UpdateNode(ctx context.Context, uuid string, node Node) (err error) {
	// Implement your middleware logic here

	return e.next.UpdateNode(ctx, uuid, node)
}
func (e eventsMiddleware) AddRegion(ctx context.Context, region Region) (err error) {
	// Implement your middleware logic here

	return e.next.AddRegion(ctx, region)
}
func (e eventsMiddleware) GetRegionById(ctx context.Context, uuid string) (region Region, err error) {
	// Implement your middleware logic here

	return e.next.GetRegionById(ctx, uuid)
}
func (e eventsMiddleware) ListRegions(ctx context.Context, args map[string]interface{}) (regions []Region, err error) {
	// Implement your middleware logic here

	return e.next.ListRegions(ctx, args)
}
func (e eventsMiddleware) DeleteRegion(ctx context.Context, uuid string) (err error) {
	// Implement your middleware logic here

	return e.next.DeleteRegion(ctx, uuid)
}
func (e eventsMiddleware) UpdateRegion(ctx context.Context, uuid string, region Region) (err error) {
	// Implement your middleware logic here

	return e.next.UpdateRegion(ctx, uuid, region)
}
