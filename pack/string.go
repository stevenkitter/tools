package pack

type String string

// Equal
type Compare interface {
	ExactId() string
}

func (s String) ExactId() string {
	return string(s)
}

// Contain 包含
func Contain(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}
