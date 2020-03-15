package universe

const (
	// SystemRunPath is the path to write temporary runtime system related files
	// and directories.
	SystemRunPath = "/run/system"

	// MachineSocketPath is the path to file socket of machine API.
	MachineSocketPath = SystemRunPath + "/machined/machine.sock"
)
