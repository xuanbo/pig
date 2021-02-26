package condition

import "github.com/xuanbo/pig/client/orm/reflect"

// CombineClause 多短语
type CombineClause struct {
	Combine Combine  `json:"combine"`
	Clauses []Clause `json:"clauses"`
}

// Type combine
func (c *CombineClause) Type() string {
	return "combine"
}

// NewCombineClause 创建
func NewCombineClause(combine Combine) *CombineClause {
	return &CombineClause{Combine: combine, Clauses: make([]Clause, 0, 8)}
}

// Add add clause
func (c *CombineClause) Add(clause Clause, others ...Clause) {
	var cc *CombineClause
	if c.Combine == CombineAnd {
		cc = And(clause, others...)
	} else if c.Combine == CombineOr {
		cc = Or(clause, others...)
	}
	if len(cc.Clauses) == 0 {
		return
	}
	c.Clauses = append(c.Clauses, cc.Clauses...)
}

// And and
func And(left Clause, rights ...Clause) *CombineClause {
	clauses := make([]Clause, 0, 8)
	if !reflect.IsNil(left) {
		clauses = append(clauses, left)
	}
	for _, clause := range rights {
		if !reflect.IsNil(clause) {
			continue
		}
		clauses = append(clauses, clause)
	}
	return &CombineClause{Combine: CombineAnd, Clauses: clauses}
}

// Or or
func Or(left Clause, rights ...Clause) *CombineClause {
	clauses := make([]Clause, 0, 8)
	if !reflect.IsNil(left) {
		clauses = append(clauses, left)
	}
	for _, clause := range rights {
		if !reflect.IsNil(clause) {
			continue
		}
		clauses = append(clauses, clause)
	}
	return &CombineClause{Combine: CombineOr, Clauses: clauses}
}
