package orm_test

import (
	"context"
	"testing"

	"github.com/xuanbo/pig"
	"github.com/xuanbo/pig/client/orm"
	"github.com/xuanbo/pig/client/orm/condition"
	"github.com/xuanbo/pig/entity"
	"github.com/xuanbo/pig/model"
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
	if err := client.API.AutoMigrate(context.TODO(), &Email{}); err != nil {
		t.Error(err)
	}
}

func TestInsert(t *testing.T) {
	email := &Email{Name: "1345545983@qq.com"}
	if err := client.API.Insert(context.TODO(), email); err != nil {
		t.Error(err)
	}
	t.Logf("email: %v", email)

}

func TestFindByID(t *testing.T) {
	var email Email
	if err := client.API.FindByID(context.TODO(), "1364417871685357568", &email); err != nil {
		t.Error(err)
	}
	t.Logf("email: %v", email)
}

func TestFindByIDs(t *testing.T) {
	var emails []Email
	if err := client.API.FindByIDs(context.TODO(), []string{"1364417871685357568"}, &emails); err != nil {
		t.Error(err)
	}
	t.Logf("emails: %v", emails)
}

type EmailCondition struct {
	ID   string `json:"id" gorm:"type:string;size:30" condition:"name:id;op:eq"`
	Name string `json:"name" gorm:"type:string;size:30" condition:"name:name;op:like"`
}

func TestQuery(t *testing.T) {
	var emails []Email
	if err := client.API.Query(context.TODO(), &EmailCondition{ID: "1364417871685357568", Name: "1345545983"}, &emails); err != nil {
		t.Error(err)
	}
	t.Logf("emails: %v", emails)

	if err := client.API.Query(context.TODO(), nil, &emails); err != nil {
		t.Error(err)
	}
	t.Logf("emails: %v", emails)

	if err := client.API.Query(context.TODO(), condition.Like("name", "1345545983"), &emails); err != nil {
		t.Error(err)
	}
	t.Logf("emails: %v", emails)
}

func TestCount(t *testing.T) {
	total, err := client.API.Count(context.TODO(), &EmailCondition{ID: "1364417871685357568", Name: "1345545983"}, &Email{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("total: %v", total)

	total, err = client.API.Count(context.TODO(), nil, &Email{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("total: %v", total)

	total, err = client.API.Count(context.TODO(), condition.Like("name", "1345545983"), &Email{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("total: %v", total)
}

func TestPage(t *testing.T) {
	var (
		emails     []*Email
		pagination = model.NewPagination(1, 10)
	)
	err := client.API.Page(context.TODO(), &EmailCondition{ID: "1364417871685357568", Name: "1345545983"}, pagination, &emails)
	if err != nil {
		t.Error(err)
	}
	t.Logf("pagination: %v", pagination)

	err = client.API.Page(context.TODO(), nil, pagination, &emails)
	if err != nil {
		t.Error(err)
	}
	t.Logf("pagination: %v", pagination)

	err = client.API.Page(context.TODO(), condition.Like("name", "1345545983"), pagination, &emails)
	if err != nil {
		t.Error(err)
	}
	t.Logf("pagination: %v", pagination)
}
