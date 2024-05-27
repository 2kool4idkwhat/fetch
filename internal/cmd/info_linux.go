//go:build linux

package cmd

import (
	"fmt"
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

func kernelLine() string {
	if OwO {
		return fmt.Sprintf("%s: %s", green("Kewnew"), Kernel())
	}

	return fmt.Sprintf("%s: %s", green("Kernel"), Kernel())
}

// returns the cpu arch
func Arch() string {

	arch, err := os.ReadFile("/proc/sys/kernel/arch")
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(arch))
}

func archLine() string {
	if OwO {
		return fmt.Sprintf("%s: %s", green("CPUwU Awch"), Arch())
	}

	return fmt.Sprintf("%s: %s", green("CPU Arch"), Arch())
}

// returns de/wm/compositor
func Desktop() string {
	// TODO: find a better way to do this

	// see https://www.freedesktop.org/software/systemd/man/latest/pam_systemd.html#desktop=
	desktop, exists := os.LookupEnv("XDG_SESSION_DESKTOP")
	if exists {
		return desktop
	}

	// if $XDG_SESSION_DESKTOP doesn't exist
	fallback, exists := os.LookupEnv("DESKTOP_SESSION")
	if exists {
		return fallback
	}

	return "unknown"
}

func desktopLine() string {

	return fmt.Sprintf("%s: %s", green("DE/WM"), Desktop())
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

	if OwO {
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

func distroLine() string {
	if OwO {
		return fmt.Sprintf("%s: %s", green("OwOS"), Distro())
	}

	return fmt.Sprintf("%s: %s", green("OS"), Distro())
}

// returns the user's shell
func Shell() string {
	shellenv, exists := os.LookupEnv("SHELL")
	if !exists {
		return "unknown"
	}

	// $SHELL will look like /usr/bin/bash so we need to
	// remove the last slash and the stuff before it

	// this splits $SHELL into a slice like ["usr", "bin", "bash"]
	shell := strings.Split(shellenv, "/")

	// and now we just return the last item from the slice
	return shell[len(shell)-1]
}

func shellLine() string {
	if OwO {
		return fmt.Sprintf("%s: %s", green("Sheww"), Shell())
	}

	return fmt.Sprintf("%s: %s", green("Shell"), Shell())
}
