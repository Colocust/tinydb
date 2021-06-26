package structure

type (
	ZRangeSpec struct {
		min, max     float32
		minex, maxex bool //与最小值/最大值是否互斥（暂时理解为是否允许与最小值/最大值相同）
	}
)

func NewZRangeSpec(min, max float32, minex, maxex bool) *ZRangeSpec {
	return &ZRangeSpec{
		min:   min,
		max:   max,
		minex: minex,
		maxex: maxex,
	}
}

func (zrs *ZRangeSpec) isValueGteMin(value float32) bool {
	if zrs.minex {
		if value > zrs.min {
			return true
		}
		return false
	} else {
		if value >= zrs.min {
			return true
		}
		return false
	}
}

func (zrs *ZRangeSpec) isValueLteMax(value float32) bool {
	if zrs.maxex {
		if value < zrs.max {
			return true
		}
		return false
	} else {
		if value <= zrs.max {
			return true
		}
		return false
	}
}

func (zrs *ZRangeSpec) isValueLtMin(value float32) bool {
	if zrs.minex {
		if value <= zrs.min {
			return true
		}
		return false
	} else {
		if value < zrs.min {
			return true
		}
		return false
	}
}

func (zrs *ZRangeSpec) isValueGtMax(value float32) bool {
	if zrs.maxex {
		if value >= zrs.max {
			return true
		}
		return false
	} else {
		if value > zrs.max {
			return true
		}
		return false
	}
}
