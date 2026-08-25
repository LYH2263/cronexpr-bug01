package field

import "fmt"

// Bitmap stores allowed values for one cron field.
type Bitmap struct {
	kind Kind
	bits []bool
	lo   int
	hi   int
}

func NewBitmap(kind Kind) Bitmap {
	lo, hi := bounds(kind)
	bits := make([]bool, hi-lo+1)
	return Bitmap{kind: kind, bits: bits, lo: lo, hi: hi}
}

func bounds(kind Kind) (int, int) {
	switch kind {
	case Second, Minute:
		return 0, 59
	case Hour:
		return 0, 23
	case DayOfMonth:
		return 1, 31
	case Month:
		return 1, 12
	case DayOfWeek:
		return 0, 6
	default:
		return 0, 59
	}
}

func (b Bitmap) Bounds() (int, int) { return b.lo, b.hi }

func (b *Bitmap) SetAll() {
	for i := range b.bits {
		b.bits[i] = true
	}
}

func (b *Bitmap) Set(v int) {
	if v < b.lo || v > b.hi {
		return
	}
	b.bits[v-b.lo] = true
}

func (b Bitmap) Has(v int) bool {
	if v < b.lo || v > b.hi {
		return false
	}
	return b.bits[v-b.lo]
}

func (b Bitmap) Clone() Bitmap {
	cp := make([]bool, len(b.bits))
	copy(cp, b.bits)
	return Bitmap{kind: b.kind, bits: cp, lo: b.lo, hi: b.hi}
}

func (b Bitmap) Values() []int {
	out := make([]int, 0, 8)
	for i, ok := range b.bits {
		if ok {
			out = append(out, i+b.lo)
		}
	}
	return out
}

// BitSlice exposes raw bits for diagnostics (returns copy).
func (b Bitmap) BitSlice() []bool {
	cp := make([]bool, len(b.bits))
	copy(cp, b.bits)
	return cp
}

func (b Bitmap) String() string {
	vals := b.Values()
	if len(vals) == 0 {
		return "empty"
	}
	return fmt.Sprintf("%v", vals)
}
