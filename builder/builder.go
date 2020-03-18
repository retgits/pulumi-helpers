// Package builder helps with generating zip files for AWS Lambda functions.
// The builder package assumes that running "go build" will suffice to build
// the executable and will create a zipfile with the same name as the parent
// folder.
package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

const (
	// Shells are the executables ran for each OS
	windowsShell = "cmd"
	linuxShell   = "sh"
	darwinShell  = "sh"

	// ShellFlags are the flags needed to allow reading from string
	windowsFlag = "/C"
	linuxFlag   = "-c"
	darwinFlag  = "-c"

	// Args are the args needed to build or zip the AWS Lambda function
	buildArgs = "GOOS=linux GOARCH=amd64 go build"
	zipArgs   = "zip ./%s.zip ./%s"

	// UnknownRuntimeErr is the error returned when the runtime is unknown
	UnknownRuntimeErr = "unknown runtime %s"
)

// Factory is the main struct to create new zip files.
type Factory struct {
	// Folder is the root folder where the go files for the function exist
	folder string
}

// NewFactory returns a new Factory pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewFactory() *Factory {
	return &Factory{}
}

// WithFolder sets the root folder to use and returns a pointer to the
// existing resource to allow chaining.
func (f *Factory) WithFolder(folder string) *Factory {
	f.folder = folder
	return f
}

// Build runs go build in the folder set by the Factory
func (f *Factory) Build() error {
	return f.run(buildArgs)
}

// MustBuild is like Build but panics if an error is returned
func (f *Factory) MustBuild() {
	err := f.run(buildArgs)
	if err != nil {
		panic(err)
	}
}

// Zip runs the zip command in the folder set by the Factory
func (f *Factory) Zip() error {
	_, file := path.Split(f.folder)
	return f.run(fmt.Sprintf(zipArgs, file, file))
}

// MustZip is like Zip but panics if an error is returned
func (f *Factory) MustZip() {
	_, file := path.Split(f.folder)
	err := f.run(fmt.Sprintf(zipArgs, file, file))
	if err != nil {
		panic(err)
	}
}

// run creates a Cmd struct to execute the named program with the given arguments.
// After that, it starts the specified command and waits for it to complete.
func (f *Factory) run(args string) error {
	var shell, shellFlag string

	switch runtime.GOOS {
	case "windows":
		shell = windowsShell
		shellFlag = windowsFlag
	case "darwin":
		shell = darwinShell
		shellFlag = darwinFlag
	case "linux":
		shell = linuxShell
		shellFlag = linuxFlag
	default:
		return fmt.Errorf(UnknownRuntimeErr, runtime.GOOS)
	}

	cmd := exec.Command(shell, shellFlag, args)
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = f.folder
	fmt.Println(cmd.Dir)
	return cmd.Run()
}
