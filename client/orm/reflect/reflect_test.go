package reflect_test

import (
	"encoding/json"
	"testing"

	"github.com/xuanbo/pig/client/orm/reflect"
)

type Email struct {
	Name string `json:"name" gorm:"type:string;size:30" condition:"name:name;op:like;"`
}

func TestGetStructFields(t *testing.T) {
	email := Email{
		Name: "1345545983@qq.com",
	}
	fields := reflect.Fields(email)
	b, err := json.Marshal(fields)
	if err != nil {
		t.Error(err)
	}
	t.Logf("fields: %s", string(b))
}
