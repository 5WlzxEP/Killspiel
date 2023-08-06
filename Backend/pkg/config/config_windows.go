package config

import "os"

var globalPaths = [...]string{os.Getenv("APPDATA"), os.Getenv("HOMEPATH") + "/.config"}
