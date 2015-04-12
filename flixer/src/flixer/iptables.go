// Package flixer handles iptables munging for netflix addresses
package flixer

import (
	"log"
	"os"
	"os/exec"
	"text/template"
)

const (
	// Template for iptables script.
	scriptTmpl = "iptables.tmpl"
	// Script that will be written/run.
	scriptFile = "iptables.sh"
)

// The setting for a user.
type ClientSetting struct {
	// Address of the end user.
	ClientIP string
	// Local address to redirect to.
	ProxyIP string
}

type IPTables struct {
}

// Commit commits the IPTables configuration.
func (i *IPTables) Commit() error {
	if err := i.writeConfig(); err != nil {
		return err
	}
	if err := i.restart(); err != nil {
		return err
	}
	return nil
}

// Write the configuration to the iptables script.
func (i *IPTables) writeConfig() error {
	// Load the current user config.
	cfg := AllConfig{}
	if err := cfg.Read(); err != nil {
		return err
	}
	// Load the iptables template.
	t, err := template.ParseFiles(scriptTmpl)
	if err != nil {
		return err
	}
	// Extract region/endpoints.
	var rCfg RegionConfig
	rCfg, err = RegionCfg()
	if err != nil {
		return err
	}
	// Create data structure to pass to the template.
	data := make([]ClientSetting, 0)
	for _, user := range cfg.Users {
		ur, err := UserRegion(user.Address, "")
		if err != nil || len(ur) < 1 {
			continue
		}
		cs := ClientSetting{
			ClientIP: user.Address,
			ProxyIP:  rCfg[ur],
		}
		data = append(data, cs)
	}
	// Open the iptables script.
	f, err := os.Create(scriptFile)
	defer f.Close()
	if err != nil {
		return err
	}
	// Run the template.
	if err := t.Execute(f, data); err != nil {
		return err
	}
	log.Printf("Wrote iptables config file.")
	return nil
}

// Reloads the iptables setup.
func (i *IPTables) restart() error {
	cmd := exec.Command("sudo", "/bin/sh", "/opt/flixer/iptables.sh")
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Printf("Reconfigured iptables.")
	return nil
}
