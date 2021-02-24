package service

import (
	"context"

	"github.com/xuanbo/pig/client/orm"
	"github.com/xuanbo/pig/entity"
)

// User 用户服务
type User struct {
	orm *orm.Client
}

// Login 登录
func (User) Login(ctx context.Context, user *entity.User) (map[string]interface{}, error) {
	return nil, nil
}
