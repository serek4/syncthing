// Copyright (C) 2025 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/coreos/go-semver/semver"
)

func main() {

	// Get the latest "v1.22.3" or "v1.22.3-1" style tag.
	latestTag, err := cmd("git", "describe", "--abbrev=0", "--match", "v[0-9].*")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	latest, err := semver.NewVersion(latestTag[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	if latest.PreRelease != "" {

		// We are already on a prerelease.
		// Bump the prerelease counter.
		parts := latest.PreRelease.Slice()
		for i, p := range parts {
			if v, err := strconv.Atoi(p); err == nil {
				parts[i] = strconv.Itoa(v + 1)
				latest.PreRelease = semver.PreRelease(strings.Join(parts, "."))
				fmt.Println("v" + latest.String())
				return
			}
		}
	}
	latest.PreRelease = "1"

	fmt.Println("v" + latest.String())
}

func cmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bs)), nil
}
