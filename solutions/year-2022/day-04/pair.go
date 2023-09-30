package main

type Pair struct {
	start, end int
}

func (p Pair) Contains(o Pair) bool {
	return p.start <= o.start && o.end <= p.end
}
