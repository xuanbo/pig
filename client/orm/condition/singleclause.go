package condition

import (
	"github.com/xuanbo/pig/client/orm/reflect"
)

// SingleClause 单一短语
type SingleClause struct {
	Name  string      `json:"name"`
	Op    Op          `json:"op"`
	Value interface{} `json:"value"`
}

// Type single
func (c *SingleClause) Type() string {
	return "single"
}

// Eq =
func Eq(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpEq, Value: value}
}

// NotEq !=
func NotEq(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpNotEq, Value: value}
}

// Gt >
func Gt(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpGt, Value: value}
}

// Gte >=
func Gte(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpGte, Value: value}
}

// Lt <
func Lt(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpLt, Value: value}
}

// Lte <=
func Lte(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpGte, Value: value}
}

// Like like
func Like(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpLike, Value: value}
}

// NotLike not like
func NotLike(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) || isEmpty(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpNotLike, Value: value}
}

// In in
func In(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) {
		return nil
	}
	if !reflect.IsSlice(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpIn, Value: value}
}

// NotIn not in
func NotIn(name string, value interface{}) Clause {
	if name == "" || reflect.IsNil(value) {
		return nil
	}
	if !reflect.IsSlice(value) {
		return nil
	}
	return &SingleClause{Name: name, Op: OpNotIn, Value: value}
}

// IsNull is null
func IsNull(name string) Clause {
	if name == "" {
		return nil
	}
	return &SingleClause{Name: name, Op: OpIsNull}
}

// IsNotNull is not in
func IsNotNull(name string) Clause {
	if name == "" {
		return nil
	}
	return &SingleClause{Name: name, Op: OpIsNotNull}
}

func isEmpty(value interface{}) bool {
	if s, ok := value.(string); ok && s == "" {
		return true
	}
	return false
}
