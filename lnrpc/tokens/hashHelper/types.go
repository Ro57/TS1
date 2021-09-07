package hashhelper

type address string
type transaction string
type block string
type lock string

func (a *address) validate() bool {
	return true
}

func (t *transaction) validate() bool {
	return true
}

func (b *block) validate() bool {
	return true
}

func (l *lock) validate() bool {
	return true
}
