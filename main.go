package main

import (
	"log"
	"time"

	"golang.org/x/sys/unix"

	"github.com/talos-systems/machined/internal/grpc/factory"
	"github.com/talos-systems/machined/pkg/runtime"
	v1alpha1runtime "github.com/talos-systems/machined/pkg/runtime/v1alpha1"
	v1alpha1server "github.com/talos-systems/machined/pkg/server/v1alpha1"
	"github.com/talos-systems/talos/pkg/universe"
)

func reboot() {
	for i := 10; i >= 0; i-- {
		log.Printf("rebooting in %d seconds\n", i)
		time.Sleep(1 * time.Second)
	}

	unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
}

func main() {
	controller := runtime.Controller{
		Sequencer: &v1alpha1runtime.Sequencer{},
	}

	if err := controller.Run(runtime.Initialize, nil); err != nil {
		log.Println(err)
		goto reboot
	}

	// TODO: Update the controller with the platform, and configurator.
	// Investigate how we might return, or pass the configurator and platform to
	// the subsequent sequences.

	if err := controller.Run(runtime.Boot, nil); err != nil {
		log.Println(err)
		goto reboot
	}

	server := &v1alpha1server.Server{
		Controller: controller,
	}

	err := factory.ListenAndServe(server, factory.Network("unix"), factory.SocketPath(universe.MachineSocketPath))
	if err != nil {
		log.Println(err)
		goto reboot
	}

reboot:
	reboot()
}
