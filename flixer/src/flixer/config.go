// Package flixer handles Flixer configuration per user (address)
package flixer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const (
	// The live config file.
	configFile = "config.json"
	// Temporary config file to write.
	tmpFile = "config.json.tmp"
)

var mu sync.Mutex

// UserConfig is the configuration for a given user/address.
type UserConfig struct {
	// End address of the user.
	Address string
	// Region selected for the user.
	Region string
}

// Regionconfig is a map from region name ("uk") to local
// ip alias address ("10.2.0.2").
type RegionConfig map[string]string

// AllConfig represents the global configuration.
type AllConfig struct {
	// Config for all users.
	Users []*UserConfig
	// All supported regions.
	Regions RegionConfig
}

// Read reads the config file into this struct.
func (a *AllConfig) Read() error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}
	return nil
}

// Write writes this struct into the config file.
func (a *AllConfig) Write() error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}
	if err := os.Rename(tmpFile, configFile); err != nil {
		return err
	}
	return nil
}

// RegionCfg returns the current RegionConfig (region->address).
func RegionCfg() (RegionConfig, error) {
	mu.Lock()
	defer mu.Unlock()
	a := AllConfig{}
	// Open any existing config file.
	if err := a.Read(); err != nil {
		fmt.Printf("Cant read config file: %v\n", err)
		return nil, err
	}
	return a.Regions, nil
}

// Regions returns the supported regions.
func Regions() ([]string, error) {
	var err error
	var rc RegionConfig
	rc, err = RegionCfg()
	if err != nil {
		fmt.Printf("Cant read config file: %v\n", err)
		return []string{}, err
	}
	regions := make([]string, len(rc))
	i := 0
	for r := range rc {
		regions[i] = r
		i += 1
	}
	return regions, nil
}

// UserRegion gets or sets the region for a user.
func UserRegion(user, region string) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	a := AllConfig{}
	// Open any existing config file.
	if err := a.Read(); err != nil {
		fmt.Printf("Creating new config file.\n")
	}
	var cfg *UserConfig
	// Get any existing matching UserConfig.
	for _, userConfig := range a.Users {
		if userConfig.Address == user {
			cfg = userConfig
		}
	}
	// If no match, create a new one.
	if cfg == nil {
		cfg = &UserConfig{Address: user, Region: region}
		a.Users = append(a.Users, cfg)
	}
	// If region is specified, save it.
	if region != "" {
		cfg.Region = region
		if err := a.Write(); err != nil {
			return "", err
		}
	}
	return cfg.Region, nil
}
