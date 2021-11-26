package main

import (
	"log"

	"github.com/danstis/go-apt-token/internal/version"
)

// Main entry point for the app.
func main() {
	log.Printf("Version %q", version.Version)
}
