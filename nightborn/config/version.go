/* This file is auto-generated, manual edits in this file will be overwritten! */
package config

import (
	"os/exec"
	"strings"
)

// Get version from environment variable or from git
func getVersion(version string) *string {
	if version == "" {
		return getVersionFromGit()
	}
	return &version
}

func getVersionFromGit() *string {
	// Derive a version string from Git. Example outputs:
	// 	v1.0.1-0-g9de4
	// 	v2.0-8-g77df-dirty
	// 	4f72d7
	cmd := exec.Command("git", "describe", "--long", "--always", "--dirty")
	b, err := cmd.Output()
	if err != nil {
		// Either Git is not available or the current directory is not a Git repository.
		return nil
	}
	version := strings.TrimSpace(string(b))
	return &version
}
