package user

type Name string

func NewName(n string) Name {
	return Name(n)
}

func (n Name) String() string {
	return string(n)
}
