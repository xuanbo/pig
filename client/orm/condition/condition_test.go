package condition_test

import (
	"encoding/json"
	"testing"

	"github.com/xuanbo/pig/client/orm/condition"
)

type Condition struct {
	Name   string   `json:"name" condition:"name:name;op:like"`
	Age    *int     `json:"age" condition:"name:age;op:gt"`
	Status *bool    `json:"status" condition:"name:status;op:eq"`
	List   []string `json:"list" condition:"name:list;op:in"`
}

func TestParseCondition(t *testing.T) {
	var c Condition
	err := json.Unmarshal([]byte(`{"name": "zhangsan", "list": ["1", "2"]}`), &c)
	if err != nil {
		t.Error(err)
	}

	clause := condition.ParseCondition(c)

	b, err := json.Marshal(clause)
	if err != nil {
		t.Error(err)
	}
	t.Logf("clause: %s", string(b))
}
