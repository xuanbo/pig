package server

import "context"

// Server 服务
type Server interface {
	Serve() error
	Stop(ctx context.Context) error
}
