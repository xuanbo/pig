package orm

import (
	"context"

	"github.com/xuanbo/pig"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Client 数据库操作客户端
type Client struct {
	*gorm.DB
}

// FindByID 主键查询
// 使用例子：FindByID(ctx, "1", &email)
func (c *Client) FindByID(ctx context.Context, id string, entity interface{}) error {
	return c.WithContext(ctx).Find(entity, id).Error
}

// FindByIDs 主键批量查询
// 使用例子：FindByIDs(ctx, []string{"1", "2"}, &emails)
func (c *Client) FindByIDs(ctx context.Context, ids []string, entity interface{}) error {
	return c.WithContext(ctx).Find(entity, "id IN ?", ids).Error
}

// DeleteByID 主键删除
// 使用例子：DeleteByID(ctx, "1", &Email{})
func (c *Client) DeleteByID(ctx context.Context, id string, entity interface{}) error {
	return c.WithContext(ctx).Delete(entity, id).Error
}

// DeleteByIDs 主键批量查询
// 使用例子：DeleteByIDs(ctx, []string{"1", "2"}, &Email{})
func (c *Client) DeleteByIDs(ctx context.Context, ids []string, entity interface{}) error {
	return c.WithContext(ctx).Delete(entity, "id IN ?", ids).Error
}

// New 创建
func New() (*Client, error) {
	var (
		dns             = pig.Config().GetString("mysql.dns")
		maxIdleConns    = pig.Config().GetInt("mysql.pool.maxIdleConns")
		maxOpenConns    = pig.Config().GetInt("mysql.pool.maxOpenConns")
		connMaxLifetime = pig.Config().GetDuration("mysql.pool.connMaxLifetime")

		slowThreshold = pig.Config().GetDuration("mysql.sql.slowThreshold")
		message       = pig.Config().GetString("mysql.sql.message")
	)
	if message == "" {
		message = "Execute SQL"
	}

	gormDB, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: &Logger{
			SlowThreshold: slowThreshold,
			Message:       message,
		},
	})
	if err != nil {
		return nil, err
	}

	// 设置日志级别
	gormDB.Logger.LogMode(logger.Info)

	// 设置连接池
	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	// Ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Client{gormDB}, nil
}
