package histogram

import "golang.org/x/exp/constraints"

func NewDiscreteHistogram[T constraints.Ordered]() *DiscreteHistogram[T] {
	return &DiscreteHistogram[T]{
		Buckets: map[T]int{},
		Count:   0,
	}
}

type DiscreteHistogram[T constraints.Ordered] struct {
	// Count is the total size of all buckets.
	Count int

	// Buckets over which values are partionned.
	Buckets map[T]int
}

func (dh *DiscreteHistogram[T]) Add(value T) {
	dh.Buckets[value]++
	dh.Count++
}