package iter

// Supplier 产生一个元素
type Supplier[T any] func() T

// Consumer 消费一个元素
type Consumer[T any] func(T)

// UnaryFunction 将一个类型转为另一个类型
type UnaryFunction[T, R any] func(T) R

// Predicate 断言是否满足指定条件
type Predicate[T any] func(T) bool

// UnaryOperator 对输入进行一元运算返回相同类型的结果
type UnaryOperator[T any] func(T) T

// BinaryFunction 将两个类型转为第三个类型
type BinaryFunction[T, R, U any] func(T, R) U

// BinaryOperator 输入两个相同类型的参数，对其做二元运算，返回相同类型的结果
type BinaryOperator[T any] func(T, T) T

// Comparator 比较两个元素.
// 第一个元素大于第二个元素时，返回正数;
// 第一个元素小于第二个元素时，返回负数;
// 否则返回 0.
type Comparator[T any] func(T, T) int

type Less[T any] func(T, T) bool

// SupplierKV 产生一个KV
type SupplierKV[K, V any] func() (K, V)

// UnaryKVFunction 将一个类型转为另一个类型
type UnaryKVFunction[K, V, R any] func(K, V) R
type UnaryKVFunction2[K, V, RK, RV any] func(K, V) (RK, RV)

// Predicate 断言是否满足指定条件
type PredicateKV[K, V any] func(K, V) bool

type UnaryKVOperator[K, V any] func(K, V) (K, V)

type BinaryKVFunction[K, V, R, U any] func(K, V, R) U
type BinaryKVFunction2[K, V, RK, RV, UK, UV any] func(K, V, RK, RV) (UK, UV)

type BinaryKVOperator[K, V any] func(K, V, K, V) (K, V)

// Comparator 比较两个元素.
// 第一个元素大于第二个元素时，返回正数;
// 第一个元素小于第二个元素时，返回负数;
// 否则返回 0.
type ComparatorKV[K, V any] func(K, V, K, V) int
type LessKV[K, V any] func(K, V, K, V) bool

// ConsumerKV 消费一个KV
type ConsumerKV[K, V any] func(K, V)

func PairOf[T1 any, T2 any](f T1, s T2) Pair[T1, T2] {
	return Pair[T1, T2]{f, s}
}

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func (a *Pair[T1, T2]) Val() (T1, T2) {
	return a.First, a.Second
}

func PairPtrOf[T1 any, T2 any](f T1, s T2) *Pair[T1, T2] {
	return &Pair[T1, T2]{f, s}
}

func TupleOf[T1 any, T2 any, T3 any](f T1, s T2, t T3) *Tuple[T1, T2, T3] {
	return &Tuple[T1, T2, T3]{f, s, t}
}

type Tuple[T1 any, T2 any, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

func (a *Tuple[T1, T2, T3]) Val() (T1, T2, T3) {
	return a.First, a.Second, a.Third
}
