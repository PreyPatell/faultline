// Package osdetect identifies the host OS so chaos handlers can choose
// between native tc/cgroups (Linux) and container-level fallbacks (macOS/Windows).
package osdetect

import "runtime"

// Platform is the detected operating system.
type Platform string

const (
	Linux   Platform = "linux"
	MacOS   Platform = "darwin"
	Windows Platform = "windows"
	Unknown Platform = "unknown"
)

// Detect returns the current host OS.
func Detect() Platform {
	switch runtime.GOOS {
	case "linux":
		return Linux
	case "darwin":
		return MacOS
	case "windows":
		return Windows
	default:
		return Unknown
	}
}

// SupportsTC returns true when Linux tc (traffic control) is available.
func SupportsTC() bool { return Detect() == Linux }

// SupportsCgroups returns true when cgroups v2 is available (Linux only).
func SupportsCgroups() bool { return Detect() == Linux }
