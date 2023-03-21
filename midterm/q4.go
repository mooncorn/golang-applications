package main

type FooStruct struct {
	s1 *string
	s2 *string
	s3 *string
	i1 *int
	i2 *int
	f1 *float64
}

func return2strings1float(s *FooStruct) (*string, *string, *float64) {
	return s.s1, s.s2, s.f1
}

func return1string2integers(s *FooStruct) (*string, *int, *int) {
	return s.s3, s.i1, s.i2
}

func return1float(s *FooStruct) *float64 {
	return s.f1
}

func main() {
	s1 := "fsdfasdf"
	s2 := "fsdaffsad"
	s3 := "fsadfsfdsafs"
	i1 := 554
	i2 := 43
	f1 := 23.2

	myStruct := FooStruct{
		s1: &s1,
		s2: &s2,
		s3: &s3,
		i1: &i1,
		i2: &i2,
		f1: &f1,
	}
}
