package v1alpha1

import "github.com/talos-systems/machined/pkg/runtime"

type LevelZero struct{}

func (*LevelZero) Tasks() []runtime.Task {
	return []runtime.Task{
		&KSPP{},
	}
}
