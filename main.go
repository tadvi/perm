package perm

// Credits: based on https://stackoverflow.com/a/30230552/14660 by Paul Hankin

type Perm struct {
	orig    []int
	perm    []int
	started bool
}

func New(orig []int) *Perm {
	return &Perm{orig: orig, perm: make([]int, len(orig))}
}

func (p *Perm) Next() bool {
	if p.started {
		p.next()
	}
	p.started = true

	if p.perm[0] >= len(p.perm) {
		return false
	}
	return true
}

func (p *Perm) Get() []int {
	a := append(p.orig[:0:0], p.orig...)
	for i, v := range p.perm {
		a[i], a[i+v] = a[i+v], a[i]
	}
	return a
}

func (p *Perm) next() {
	last := len(p.perm) - 1
	for i := last; i >= 0; i-- {
		if i == 0 || last-i > p.perm[i] {
			p.perm[i]++
			return
		}
		p.perm[i] = 0
	}
}
