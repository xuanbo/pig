package condition

const (
	// OpEq =
	OpEq Op = iota
	// OpNotEq !=
	OpNotEq
	// OpGt >
	OpGt
	// OpGte >=
	OpGte
	// OpLt <
	OpLt
	// OpLte <=
	OpLte
	// OpLike like
	OpLike
	// OpNotLike not like
	OpNotLike
	// OpIn in
	OpIn
	// OpNotIn not in
	OpNotIn
	// OpIsNull is null
	OpIsNull
	// OpIsNotNull is not null
	OpIsNotNull
)

const (
	// CombineAnd and
	CombineAnd Combine = iota
	// CombineOr or
	CombineOr
)

type (
	// Op 操作
	Op uint8
	// Combine 组合
	Combine uint8
)

// Clause 短语
type Clause interface {
	Type() string
}
