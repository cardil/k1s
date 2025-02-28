package stack

type Stack uint

const (
	DevelopmentStack Stack = iota
	TestStack
	StagingStack
	ProductionStack
)

func Mapping() map[Stack][]string {
	return map[Stack][]string{
		DevelopmentStack: {"dev", "development"},
		TestStack:        {"test", "testing"},
		StagingStack:     {"stage", "staging"},
		ProductionStack:  {"prod", "production"},
	}
}

func (s Stack) String() string {
	return Mapping()[s][0]
}
