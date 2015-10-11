// This file is part of the Plan-B-ot package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package bot

import (
	"encoding/json"
	"os"
)

// Config is the configuration used throughout the  project
var Config Configuration

// Configuration holds the overall configuration settings
type Configuration struct {
	Port       string
	Route      string
	Token      string
	WebhookURL string
}

// ReadConfig reads the configuration from a file into the struct
func ReadConfig(filename string) error {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
