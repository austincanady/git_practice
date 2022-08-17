package testmod

type TestMod struct{
  value string
}

type TestModI interface {
	SetValue(string)
	Value() string
}

func NewTestMod(v string) TestModI {
	b := new(TestMod)
	b.value = v
	return b
}

// SetValue sets the bool value
func (b *TestMod) SetValue(v string) {
	b.value = v
}

// Value gets the bool flag
func (b TestMod) Value() string {
	return b.value
}
