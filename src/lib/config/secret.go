package config

type Secret string

func (s Secret) String() string {
	return "<secret>"
}
