package v1alpha1

import (
	"github.com/talos-systems/machined/api"
	"github.com/talos-systems/machined/pkg/runtime"
)

// Sequencer is an implementation of a `Sequencer`.
type Sequencer struct{}

func (*Sequencer) Initialize() []runtime.Phase {
	return []runtime.Phase{
		&LevelZero{},
	}
}

func (*Sequencer) Boot() []runtime.Phase {
	return nil
}

func (*Sequencer) Shutdown() []runtime.Phase {
	return nil
}

func (*Sequencer) Upgrade(*api.UpgradeRequest) []runtime.Phase {
	return nil
}

func (*Sequencer) Reset(*api.ResetRequest) []runtime.Phase {
	return nil
}
