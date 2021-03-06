// build linux

package sysinit

import (
	"fmt"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
	"github.com/milosgajdos83/servpeek/utils/service"
)

const (
	systemctl = "systemctl"
)

///////////////////////////////////
// TODO: re-implement using dbus //
///////////////////////////////////

// SystemdInit provides systemd init commands
type SystemdInit struct{}

// Start starts systemd service. It returns error if the service fails to be started
func (sm *SystemdInit) Start(svcName string) error {
	startCmd := command.NewCommand(systemctl, []string{"start", svcName + ".service"}...)
	_, err := startCmd.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Stop stops systemd service. It returns error if the service fails to be stopped
func (sm *SystemdInit) Stop(svcName string) error {
	stopCmd := command.NewCommand(systemctl, []string{"stop", svcName + ".service"}...)
	_, err := stopCmd.RunCombined()
	if err != nil {
		return err
	}
	return nil
}

// Status queries status of systemd service and returns it.
// It returns error if the status of the queried service could not be determined
func (sm *SystemdInit) Status(svcName string) (service.Status, error) {
	statusCmd := command.NewCommand(systemctl, []string{"status", svcName + ".service"}...)
	status, err := statusCmd.RunCombined()
	if err != nil {
		return service.Unknown, err
	}
	switch {
	case strings.Contains(status, "active (running)"):
		return service.Running, nil
	case strings.Contains(status, "inactive (stopped)"):
		return service.Stopped, nil
	}
	return service.Unknown, fmt.Errorf("Unable to determine %s status", svcName)
}
