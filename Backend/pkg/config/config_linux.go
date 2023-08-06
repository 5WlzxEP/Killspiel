package config

import "os"

var globalPaths = [...]string{"/etc", os.Getenv("HOME") + "/.config"}
