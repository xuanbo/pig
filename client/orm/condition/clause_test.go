package condition_test

import (
	"encoding/json"
	"testing"

	"github.com/xuanbo/pig/client/orm/condition"
)

func TestClause(t *testing.T) {
	clause := condition.NewCombineClause(condition.CombineAnd)
	clause.Add(condition.Eq("name", nil))
	clause.Add(condition.In("type", []interface{}{1}))

	b, err := json.Marshal(clause)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("json: %s", string(b))
}
