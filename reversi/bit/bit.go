package bit

type Bit struct {
	Value uint64
}

func NewBit() *Bit {
	return &Bit{}
}

func (b *Bit) Get() uint64 {
	return b.Value
}

func (b *Bit) Or(value uint64) {
	b.Value = b.Value | value
}

func (b *Bit) And(value uint64) {
	b.Value = b.Value & value
}

func (b *Bit) Xor(value uint64) {
	b.Value = b.Value ^ value
}

func (b *Bit) OverWrite(value uint64) {
	b.Value = value
}

func (b *Bit) Reset() {
	b.Value = 0
}
