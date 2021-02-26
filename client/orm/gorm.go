package orm

import (
	"errors"

	"github.com/xuanbo/pig"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Client 数据库操作客户端
type Client struct {
	DB      *gorm.DB
	Dialect Dialect
	API     *API
}

// New 创建
func New() (*Client, error) {
	var (
		dialectStr      = pig.Config().GetString("datasource.dialect")
		dns             = pig.Config().GetString("datasource.dns")
		maxIdleConns    = pig.Config().GetInt("datasource.pool.maxIdleConns")
		maxOpenConns    = pig.Config().GetInt("datasource.pool.maxOpenConns")
		connMaxLifetime = pig.Config().GetDuration("datasource.pool.connMaxLifetime")

		slowThreshold = pig.Config().GetDuration("datasource.sql.slowThreshold")
		message       = pig.Config().GetString("datasource.sql.message")

		gormDB    *gorm.DB
		dialector gorm.Dialector
		dialect   Dialect
		err       error
	)
	switch dialectStr {
	case "mysql":
		dialector = mysql.Open(dns)
		dialect = MySQL
	case "postgres":
		dialector = postgres.Open(dns)
		dialect = Postgres
	default:
		return nil, errors.New("配置 [datasource.dialect] 不合法，可选值: [mysql、postgres]")
	}

	if message == "" {
		message = "Execute SQL"
	}

	gormDB, err = gorm.Open(dialector, &gorm.Config{
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

	return &Client{
		gormDB,
		dialect,
		&API{gormDB, dialect},
	}, nil
}
