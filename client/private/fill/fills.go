package fill

type Fills []Fill

func (f Fills) Pri() float32 {
	var siz float32
	{
		siz = f.Siz()
	}

	var agg float32
	for _, v := range f {
		per := v.Siz() * 100 / siz
		agg += per * v.Pri()
	}

	var pri float32
	{
		pri = agg / 100
	}

	return pri
}

func (f Fills) Siz() float32 {
	var s float32

	for _, v := range f {
		s += v.Siz()
	}

	return s
}
