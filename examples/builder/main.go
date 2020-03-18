package main

import (
	// Import the builder package
	"github.com/retgits/pulumi-helpers/builder"
)

func main() {
	// Create a new factory
	factory := builder.NewFactory().WithFolder("../sampolicies")

	// Run go build and panic if an error is returned
	factory.MustBuild()

	// Run zip and panic if an error is returned
	factory.MustZip()
}
