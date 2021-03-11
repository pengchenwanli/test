package services

import (
	"context"
	"test/model"
)

const contextKey = "context key"

type Context struct {
	Token *model.Token
	Admin *model.Admin
}

func WithContext(parent context.Context, ctx *Context) context.Context {
	i, ok := parent.(interface {
		Set(key string, value interface{})
	})
	if ok {
		i.Set(contextKey, ctx)
		return parent
	}
	return context.WithValue(parent, contextKey, ctx)
	//context.Background可以理解为知道上下文要干什么，context.TOdo理解为清楚要使用哪个上下文、或者还没有可用的上下文
}
func GetContext(ctx context.Context) *Context {
	return ctx.Value(contextKey).(*Context)
}
