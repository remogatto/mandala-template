// +build darwin

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	"github.com/tideland/goas/v2/loop"
	glfw "github.com/go-gl/glfw3"
	"github.com/remogatto/mandala"
)

func main() {

	runtime.LockOSThread()

	verbose := flag.Bool("verbose", false, "produce verbose output")
	debug := flag.Bool("debug", false, "produce debug output")
	size := flag.String("size", "320x480", "set the size of the window")

	flag.Parse()

	if *verbose {
		mandala.Verbose = true
	}

	if *debug {
		mandala.Debug = true
	}

	dims := strings.Split(strings.ToLower(*size), "x")
	width, err := strconv.Atoi(dims[0])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(dims[1])
	if err != nil {
		panic(err)
	}

	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(width, height, "{{.AppName}}", nil, nil)
	if err != nil {
		panic(err)
	}

	mandala.Init(window)

	// Create a rendering loop control struct containing a set of
	// channels that control rendering.
	renderLoopControl := newRenderLoopControl()

	// Start the rendering loop
	loop.GoRecoverable(
		renderLoopFunc(renderLoopControl),
		func(rs loop.Recoverings) (loop.Recoverings, error) {
			for _, r := range rs {
				log.Printf("%s\n%s", r.Reason, mandala.Stacktrace())
			}
			return rs, fmt.Errorf("Unrecoverable loop\n")
		},
	)
	// Start the event loop
	loop.GoRecoverable(
		eventLoopFunc(renderLoopControl),
		func(rs loop.Recoverings) (loop.Recoverings, error) {
			for _, r := range rs {
				log.Printf("%s\n%s", r.Reason, mandala.Stacktrace())
			}
			return rs, fmt.Errorf("Unrecoverable loop\n")
		},
	)

	for !window.ShouldClose() {
		glfw.WaitEvents()
	}

}
