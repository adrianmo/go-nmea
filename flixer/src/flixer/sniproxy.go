// Package flixer is the service for proxying Netflix regions.
package flixer

import (
	"log"
	"os/exec"
)

// SNIProxy represents the SNIProxy service restarter.
type SNIProxy struct {
}

// Commit commits and updates SNIProxy.
func (s *SNIProxy) Commit() error {
	if err := s.restart(); err != nil {
		return err
	}
	return nil
}

func (s *SNIProxy) restart() error {
	cmd := exec.Command("sudo", "/usr/sbin/service", "sniproxy", "restart")
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Printf("Restarted sniproxy daemon.")
	return nil
}
