package user

type ID uint64

func NewID(id uint64) ID {
	return ID(id)
}

func (id ID) Uint64() uint64 {
	return uint64(id)
}

func (id ID) Uint() uint {
	return uint(id)
}
