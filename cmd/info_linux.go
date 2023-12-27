//go:build linux

package cmd

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// returns kernel version
func Kernel() string {

	// see https://www.kernel.org/doc/html/latest/admin-guide/sysctl/kernel.html?highlight=boot_id#osrelease-ostype-version
	version, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(version))
}

// returns the cpu arch
func Arch() string {

	arch, err := os.ReadFile("/proc/sys/kernel/arch")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(arch))
}

// returns de/wm/compositor
func Desktop() string {
	// TODO: find a better way to do this

	// see https://www.freedesktop.org/software/systemd/man/latest/pam_systemd.html#desktop=
	desktop, exists := os.LookupEnv("XDG_SESSION_DESKTOP")
	if exists == true {
		return desktop
	}

	// if $XDG_SESSION_DESKTOP doesn't exist
	fallback, exists := os.LookupEnv("DESKTOP_SESSION")
	if exists != true {
		return "unknown"
	}

	return fallback
}

func Distro() string {
	// see $ man os-release

	// I'm too lazy to properly parse this file, so I'm just
	// using godotenv :P
	osRelease, err := godotenv.Read("/etc/os-release")

	// if there's no NAME key in /etc/os-release then apparently
	// we can use "Linux" as a placeholder
	if err != nil {
		return "Linux"
	}

	distro := osRelease["NAME"]
	if distro == "" {
		return "Linux"
	}

	if OwO == true {
		switch distro {
		case "NixOS":
			return "NyixOwOS"
		case "Ubuntu":
			return "UwUntu"
		case "Debian GNU/Linux":
			return "Debinyan GNUwU/Linuwux"
		case "openSUSE Tumbleweed":
			return "owopenSUS"
		case "Arch Linux":
			return "Nyarch (btw)"
		}
	}

	// this simplifies distro names
	switch distro {
	case "Arch Linux":
		return "Arch (btw)"
	case "Alpine Linux":
		return "Alpine"
	case "Debian GNU/Linux":
		return "Debian"
	case "Fedora Linux":
		return "Fedora"
	case "openSUSE Tumbleweed":
		return "openSUSE"
	}

	return distro
}

// returns the user's shell
func Shell() string {
	shellenv, exists := os.LookupEnv("SHELL")
	if exists != true {
		return "unknown"
	}

	// $SHELL will look like /usr/bin/bash so we need to
	// remove the last slash and the stuff before it

	// this splits $SHELL into a slice like ["usr", "bin", "bash"]
	shell := strings.Split(shellenv, "/")

	// and now we just return the last item from the slice
	return shell[len(shell)-1]
}
