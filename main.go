package fuzzygo

import (
	"fuzzy-go/membership"
	"sync"
)

type (
	Set struct {
		Fuzzified []membership.Value
		Values    []membership.Value
		Fn        membership.Function
	}
)

func (s *Set) Fuzzify() {
	if s.Fuzzified == nil {
		s.Fuzzified = make([]membership.Value, 0)
	}
	for i := 0; i < len(s.Values); i++ {
		s.Fuzzified = append(s.Fuzzified, s.Fn(s.Values[i]))
	}
}

func fuzzify(a, b *Set) {
	wg := sync.WaitGroup{}
	if a.Fuzzified == nil || len(a.Fuzzified) < 1 {
		wg.Add(1)
		go func(wgs *sync.WaitGroup) {
			defer wgs.Done()
			a.Fuzzify()
		}(&wg)
	}
	if b.Fuzzified == nil || len(b.Fuzzified) < 1 {
		wg.Add(1)
		go func(wgs *sync.WaitGroup) {
			defer wgs.Done()
			b.Fuzzify()
		}(&wg)
	}
	wg.Wait()
}

func Union(a, b *Set) (res []membership.Value) {
	fuzzify(a, b)
	lna, lnb := len(a.Fuzzified), len(b.Fuzzified)
	ln := max(lna, lnb)
	res = make([]membership.Value, ln)
	for i := 0; i < ln; i++ {
		if i < lna && i < lnb {
			res[i] = max(a.Fuzzified[i], b.Fuzzified[i])
		} else if i < lna {
			res[i] = a.Fuzzified[i]
		} else if i < lnb {
			res[i] = b.Fuzzified[i]
		}
	}
	return
}

func Intersect(a, b *Set) (res []membership.Value) {
	fuzzify(a, b)
	lna, lnb := len(a.Fuzzified), len(b.Fuzzified)
	ln := min(lna, lnb)
	res = make([]membership.Value, ln)
	for i := 0; i < ln; i++ {
		if i < lna && i < lnb {
			res[i] = min(a.Fuzzified[i], b.Fuzzified[i])
		}
	}
	return
}
