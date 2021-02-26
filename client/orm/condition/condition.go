package condition

import "github.com/xuanbo/pig/client/orm/reflect"

// ParseCondition 解析条件
func ParseCondition(condition interface{}) Clause {
	fields := reflect.Fields(condition)
	if len(fields) == 0 {
		return nil
	}
	cc := NewCombineClause(CombineAnd)
	for _, field := range fields {
		switch field.Op {
		case "eq":
			cc.Add(Eq(field.Name, field.Value))
		case "not_eq":
			cc.Add(NotEq(field.Name, field.Value))
		case "gt":
			cc.Add(Gt(field.Name, field.Value))
		case "gte":
			cc.Add(Gte(field.Name, field.Value))
		case "lt":
			cc.Add(Lt(field.Name, field.Value))
		case "lte":
			cc.Add(Lte(field.Name, field.Value))
		case "like":
			cc.Add(Like(field.Name, field.Value))
		case "not_like":
			cc.Add(NotLike(field.Name, field.Value))
		case "in":
			cc.Add(In(field.Name, field.Value))
		case "not_in":
			cc.Add(NotIn(field.Name, field.Value))
		case "is_null":
			cc.Add(IsNull(field.Name))
		case "is_not_null":
			cc.Add(IsNotNull(field.Name))
		}
	}
	return cc
}
