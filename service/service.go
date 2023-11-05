package service

import (
	"context"
	"sync"
)

type Service struct {
	ctx context.Context
}

var once sync.Once
var svc Service

// 单例模式应用

func New(ctx context.Context) Service {
	once.Do(func() {
		svc = Service{ctx: ctx}
	})
	return svc
}
