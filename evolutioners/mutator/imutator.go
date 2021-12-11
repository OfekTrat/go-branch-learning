package mutator

type IMutator interface {
	Mutate(int) interface{}
}
