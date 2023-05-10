package prey

type StubPrey struct {
	GetSpeedFunc func() float64
}

func NewStubPrey() *StubPrey {
	return &StubPrey{}
}

func (stub *StubPrey) GetSpeed() float64 {
	return stub.GetSpeedFunc()
}