package orm_test

import (
	"context"
	"testing"

	"github.com/xuanbo/pig"
	"github.com/xuanbo/pig/client/orm"
	"github.com/xuanbo/pig/entity"
)

var client *orm.Client

type Email struct {
	entity.Entity
	Name string `json:"name" gorm:"type:string;size:30"`
}

// TableName 表名
func (Email) TableName() string {
	return "test_email"
}

func init() {
	// 初始化
	pig.Initialize()
	var err error
	if client, err = orm.New(); err != nil {
		panic(err)
	}
}

func TestAutoMigrate(t *testing.T) {
	if err := client.AutoMigrate(&Email{}); err != nil {
		t.Error(err)
	}
}

func TestCreate(t *testing.T) {
	if err := client.Create(&Email{Name: "1345545983@qq.com"}).Error; err != nil {
		t.Error(err)
	}
}

func TestFindByID(t *testing.T) {
	var email Email
	if err := client.FindByID(context.TODO(), "1364417871685357568", &email); err != nil {
		t.Error(err)
	}
	t.Logf("email: %v", email)
}
