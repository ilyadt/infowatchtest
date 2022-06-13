package histogram

import (
	"fmt"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

type kv[T constraints.Ordered] struct {
	Key   T
	Value int64
}

func fmtBucketFunc[T constraints.Ordered]() func(T) string {
	var typ T

	switch any(typ).(type) {
	case rune:
		return func(a T) string { return strconv.QuoteRune((any(a).(rune))) }
	case byte:
		return func(a T) string { return strconv.QuoteRune(rune((any(a).(byte)))) }
	default:
		return func(a T) string { return fmt.Sprintf("%+v", a) }
	}
}


func (dh *DiscreteHistogram[T]) PrintSorted() {
	fmtFunc := fmtBucketFunc[T]()

	ss := make([]kv[T], 0, len(dh.Buckets))
    for k, v := range dh.Buckets {
        ss = append(ss, kv[T]{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

	fmt.Printf("#####################\n")

	for _, v := range ss {
		fmt.Printf("%8s -- %5.2f%%\n", fmtFunc(v.Key), float64(v.Value*100)/float64(dh.Count))
	}

	fmt.Printf("#####################\n")
}
