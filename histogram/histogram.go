package histogram

import "golang.org/x/exp/constraints"

func NewDiscreteHistogram[T constraints.Ordered](preAllocBuc int) *DiscreteHistogram[T] {
	return &DiscreteHistogram[T]{
		Buckets: make(map[T]int64, preAllocBuc),
		Count:   0,
	}
}

type DiscreteHistogram[T constraints.Ordered] struct {
	// Count is the total size of all buckets.
	Count int64

	// Buckets over which values are partionned.
	Buckets map[T]int64
}

func (dh *DiscreteHistogram[T]) Add(value T) {
	dh.Buckets[value]++
	dh.Count++
}
