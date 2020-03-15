// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package runtime

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/talos-systems/machined/api"
	"github.com/talos-systems/machined/pkg/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	Controller runtime.Controller
}

// Register implements the factory.Registrator interface.
func (s *Server) Register(s *grpc.Server) {
	api.RegisterMachineServiceServer(s, c)
}

// Reboot implements the api.MachineServer interface.
func (s *Server) Reboot(ctx context.Context, in *empty.Empty) (reply *api.RebootResponse, err error) {
	go func() {
		if err := s.Controller.Run(Reboot, nil); err != nil {
			log.Println(err)
		}
	}()

	reply = &api.RebootResponse{}

	return nil, nil
}

// Shutdown implements the api.MachineServer interface.
func (s *Server) Shutdown(ctx context.Context, in *empty.Empty) (reply *api.ShutdownResponse, err error) {
	go func() {
		if err := s.Controller.Run(Shutdown, nil); err != nil {
			log.Println(err)
		}
	}()

	reply = &api.ShutdownResponse{}

	return reply, nil
}

// Upgrade initiates an upgrade.
func (s *Server) Upgrade(ctx context.Context, in *api.UpgradeRequest) (reply *api.UpgradeResponse, err error) {
	go func() {
		if err := s.Controller.Run(Upgrade, nil); err != nil {
			log.Println(err)
		}
	}()

	reply = &api.UpgradeResponse{}

	return reply, nil
}

// Reset resets the node.
func (s *Server) Reset(ctx context.Context, in *api.ResetRequest) (reply *api.ResetResponse, err error) {
	go func() {
		if err := s.Controller.Run(Reset, nil); err != nil {
			log.Println(err)
		}
	}()

	reply = &api.ResetResponse{}

	return reply, nil
}

// ServiceList returns list of the registered services and their status
func (s *Server) ServiceList(ctx context.Context, in *empty.Empty) (reply *api.ServiceListResponse, err error) {
	return nil, nil
}

// ServiceStart implements the api.MachineServer interface and starts a
// service running on Talos.
func (s *Server) ServiceStart(ctx context.Context, in *api.ServiceStartRequest) (reply *api.ServiceStartResponse, err error) {
	return nil, nil
}

// ServiceStop implements the api.MachineServer interface and stops a
// service running on Talos.
func (s *Server) ServiceStop(ctx context.Context, in *api.ServiceStopRequest) (reply *api.ServiceStopResponse, err error) {
	return nil, nil
}

// ServiceRestart implements the api.MachineServer interface and stops a
// service running on Talos.
func (s *Server) ServiceRestart(ctx context.Context, in *api.ServiceRestartRequest) (reply *api.ServiceRestartResponse, err error) {
	return nil, nil
}

// Copy implements the api.MachineServer interface and copies data out of Talos node
func (s *Server) Copy(req *api.CopyRequest, s api.MachineService_CopyServer) (err error) {
	return nil
}

// List implements the api.MachineServer interface.
func (s *Server) List(req *api.ListRequest, s api.MachineService_ListServer) (err error) {
	return nil
}

// Mounts implements the api.OSDServer interface.
func (s *Server) Mounts(ctx context.Context, in *empty.Empty) (reply *api.MountsResponse, err error) {
	return nil, nil
}

// Version implements the api.MachineServer interface.
func (s *Server) Version(ctx context.Context, in *empty.Empty) (reply *api.VersionResponse, err error) {
	return nil, nil
}

// Kubeconfig implements the osapi.OSDServer interface.
func (s *Server) Kubeconfig(empty *empty.Empty, s api.MachineService_KubeconfigServer) (err error) {
	return nil
}

// Logs provides a service or container logs can be requested and the contents of the
// log file are streamed in chunks.
// nolint: gocyclo
func (s *Server) Logs(req *api.LogsRequest, l api.MachineService_LogsServer) (err error) {
	return nil
}

// Read implements the read API.
func (s *Server) Read(in *api.ReadRequest, srv api.MachineService_ReadServer) (err error) {
	return nil
}
