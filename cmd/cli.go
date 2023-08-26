package cmd

import (
	"flag"
)

const (
	START     int = 0
	MIGRATE   int = 1
	SUPERUSER int = 2
)

func Cli() int {
	migrate := flag.Bool("migrate", false, "Make migrations")
	superuser := flag.Bool("superuser", false, "Create superuser use env SUID=123456")
	flag.Parse()
	if *migrate {
		return MIGRATE
	}
	if *superuser {
		return SUPERUSER
	}
	return START
}
