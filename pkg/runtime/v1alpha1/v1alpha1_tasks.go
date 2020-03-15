package v1alpha1

import (
	"github.com/talos-systems/machined/internal/kspp"
	"github.com/talos-systems/machined/pkg/runtime"
)

// KSPP represents the KSPP task.
type KSPP struct{}

// Func returns the runtime function.
func (task *KSPP) Func(mode runtime.Mode) runtime.TaskFunc {
	switch mode {
	case runtime.Container:
		return nil
	default:
		return task.standard
	}
}

func (task *KSPP) standard(r runtime.Runtime) (err error) {
	if err = kspp.CheckKSPPKernelParameters(); err != nil {
		return err
	}

	if err = kspp.WriteKSPPSysctls(); err != nil {
		return err
	}

	return nil
}
