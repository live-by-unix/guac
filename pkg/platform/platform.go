package platform

import (
	"runtime"
)

type Platform string

const (
	Linux   Platform = "linux"
	Darwin  Platform = "darwin"
	Windows Platform = "windows"
)

var currentPlatform Platform

func Detect() {
	switch runtime.GOOS {
	case "linux":
		currentPlatform = Linux
	case "darwin":
		currentPlatform = Darwin
	case "windows":
		currentPlatform = Windows
	default:
		currentPlatform = Platform(runtime.GOOS)
	}
}

func Current() Platform {
	return currentPlatform
}

func IsLinux() bool {
	return currentPlatform == Linux
}

func IsDarwin() bool {
	return currentPlatform == Darwin
}

func IsWindows() bool {
	return currentPlatform == Windows
}

func String() string {
	return string(currentPlatform)
}
