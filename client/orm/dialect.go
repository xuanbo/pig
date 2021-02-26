package orm

import (
	"fmt"
	"strings"

	"github.com/xuanbo/pig/client/orm/condition"
	"github.com/xuanbo/pig/client/orm/reflect"
)

const (
	// MySQL MySQL驱动
	MySQL Dialect = iota
	// Postgres pg驱动
	Postgres
)

// Dialect 数据库方言类型
type Dialect uint8

// ParseClause 解析短语
func (dialect Dialect) ParseClause(clause condition.Clause) (string, interface{}) {
	if reflect.IsNil(clause) {
		return "", nil
	}
	if sc, ok := clause.(*condition.SingleClause); ok {
		return dialect.parseSingleClause(sc)
	}
	if cc, ok := clause.(*condition.CombineClause); ok {
		return dialect.parseCombineClause(cc)
	}
	return "", nil
}

func (dialect Dialect) parseSingleClause(clause *condition.SingleClause) (string, interface{}) {
	var name string
	if dialect == MySQL {
		name = "`" + clause.Name + "`"
	}
	switch clause.Op {
	case condition.OpEq:
		return name + " = ?", clause.Value
	case condition.OpNotEq:
		return name + " <> ?", clause.Value
	case condition.OpGt:
		return name + " > ?", clause.Value
	case condition.OpGte:
		return name + " >= ?", clause.Value
	case condition.OpLt:
		return name + " < ?", clause.Value
	case condition.OpLte:
		return name + " <= ?", clause.Value
	case condition.OpLike:
		return name + " LIKE ?", "%" + fmt.Sprintf("%v", clause.Value) + "%"
	case condition.OpNotLike:
		return name + " NOT LIKE ?", "%" + fmt.Sprintf("%v", clause.Value) + "%"
	case condition.OpIn:
		return name + " IN ?", clause.Value
	case condition.OpNotIn:
		return name + " NOT IN ?", clause.Value
	case condition.OpIsNull:
		return name + " IS NULL", nil
	case condition.OpIsNotNull:
		return name + " IS NOT NULL", nil
	default:
		return "", nil
	}
}

func (dialect Dialect) parseCombineClause(clause *condition.CombineClause) (string, []interface{}) {
	sl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	for _, c := range clause.Clauses {
		s, v := dialect.ParseClause(c)
		if s == "" || v == nil {
			continue
		}
		sl = append(sl, "("+s+")")
		vl = append(vl, v)
	}
	if len(vl) == 0 {
		return "", nil
	}
	if clause.Combine == condition.CombineAnd {
		return strings.Join(sl, " AND "), vl
	}
	return strings.Join(sl, " OR "), vl
}
