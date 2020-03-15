package runtime

import (
	"fmt"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/talos-systems/machined/api"
	"google.golang.org/grpc"
)

// Runtime defines the runtime parameters.
type Runtime interface {
	Platform() Platform
	Config() Configurator
}

// Controller represents the controller responsible for managing the execution
// of sequences.
type Controller struct {
	*grpc.Server

	Sequencer Sequencer
	Runtime   Runtime

	semaphore int32
}

// Run executes all phases known to a `Controller` in serial. `Controller`
// aborts immediately if any phase fails.
func (c *Controller) Run(seq Sequence, data interface{}) (err error) {
	if c.TryLock() {
		return ErrLocked
	}

	defer c.Unlock()

	var phases []Phase

	switch seq {
	case Boot:
		phases = c.Sequencer.Boot()
	case Initialize:
		phases = c.Sequencer.Initialize()
	case Shutdown:
		phases = c.Sequencer.Shutdown()
	case Upgrade:
		var (
			req *api.UpgradeRequest
			ok  bool
		)

		if req, ok = data.(*api.UpgradeRequest); !ok {
			return ErrInvalidSequenceData
		}

		phases = c.Sequencer.Upgrade(req)
	case Reset:
		var (
			req *api.ResetRequest
			ok  bool
		)

		if req, ok = data.(*api.ResetRequest); !ok {
			return ErrInvalidSequenceData
		}

		phases = c.Sequencer.Reset(req)
	}

	// We must ensure that the runtime is configured since all sequences,
	// excluding initialize, depend on the runtime.

	if seq != Initialize {
		if c.Runtime == nil {
			return ErrUndefinedRuntime
		}
	}

	log.Printf("[sequence]: %q", seq.String())

	start := time.Now()

	for i, phase := range phases {
		log.Printf("[phase]: %s", fmt.Sprintf("%s %d/%d", strings.Title(seq.String()), i+1, len(phases)+1))

		if err = c.runPhase(phase); err != nil {
			return fmt.Errorf("error running phase %q: %w", name, err)
		}
	}

	log.Printf("[sequence]: %q, done: %s", seq.String(), time.Since(start))

	return nil
}

// Phase represents a collection of tasks to be performed concurrently.
type Phase interface {
	Tasks() []Task
}

func (c *Controller) runPhase(phase Phase) error {
	errCh := make(chan error)

	start := time.Now()

	for i, task := range phase.Tasks() {
		log.Printf("[task]: %s", fmt.Sprintf("%s %d/%d", name, i+1, len(tasks)+1))
		go c.runTask(task, errCh)
	}

	var result *multierror.Error

	for range phase.Tasks() {
		err := <-errCh
		if err != nil {
			log.Printf("[phase]: %s error running task: %s", name, err)
		}

		result = multierror.Append(result, err)
	}

	log.Printf("[phase]: %s done, %s", name, time.Since(start))

	return result.ErrorOrNil()
}

// TaskFunc defines the function that a task will execute for a specific runtime
// mode.
type TaskFunc func(Runtime) error

// Task represents a task within a `Phase`.
type Task interface {
	Func(Mode) TaskFunc
}

func (c *Controller) runTask(t Task, errCh chan<- error) {
	var err error

	defer func() {
		errCh <- err
	}()

	if f := t.Func(c.Runtime.Platform().Mode()); f != nil {
		start := time.Now()

		err = f(c.Runtime)

		log.Printf("[task]: %s", time.Since(start))
	}

	return
}

// TryLock attempts to set a lock that prevents multiple sequences from running
// at once. If currently locked, a value of true will be returned. If not
// currently locked, a value of false will be returned.
func (c *Controller) TryLock() bool {
	return !atomic.CompareAndSwapInt32(&c.semaphore, 0, 1)
}

// Unlock removes the lock set by `TryLock`.
func (c *Controller) Unlock() bool {
	return atomic.CompareAndSwapInt32(&c.semaphore, 1, 0)
}
