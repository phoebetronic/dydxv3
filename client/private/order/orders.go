package order

type Orders []Order

func (o Orders) Add(a Orders) Orders {
	o = append(o, a...)
	return o
}

func (o Orders) Con(t Order) bool {
	for _, v := range o {
		if v.ID == t.ID {
			return true
		}
	}

	return false
}

func (o Orders) Pri() float32 {
	var siz float32
	{
		siz = o.Siz()
	}

	var agg float32
	for _, v := range o {
		per := v.Siz() * 100 / siz
		agg += per * v.Pri()
	}

	var pri float32
	{
		pri = agg / 100
	}

	return pri
}

func (o Orders) Rem(r Orders) Orders {
	var n Orders

	for _, v := range o {
		if !r.Con(v) {
			n = append(n, v)
		}
	}

	return n
}

func (o Orders) Siz() float32 {
	var s float32

	for _, v := range o {
		s += v.Siz()
	}

	return s
}

func (o Orders) Sta(s string) Orders {
	var n Orders

	for _, v := range o {
		if v.Status == s {
			n = append(n, v)
		}
	}

	return n
}
