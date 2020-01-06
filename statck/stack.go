type Stack struct{
	data[] int
}


func New() *Stack{
	return &Stack{
		data: []int{},
	}
}

func (s *Statck) IsEmpty bool