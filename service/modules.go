package service

import (
	"context"
	"errors"
)

//nolint:lll
//go:generate mockgen -source modules.go -package mock -destination mock/modules.go Modules
type Modules interface {
	Set(ctx context.Context, name string, module Module) error
	Get(ctx context.Context, name string) (Module, error)
	HealthCheck() HealthCheck
}

var (
	ErrWritingModuleFailed = errors.New("module write failed")
	ErrReadingModule       = errors.New("module read failed")
)
