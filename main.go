package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	verbose *bool
	rootDir = "template"
)

func verboseLog(format string, message ...interface{}) {
	if *verbose {
		log.Printf(format, message...)
	}
}

func copyFile(srcFile, dstPath, dstFile string) error {
	// Get the full path of destination file
	fullDstPath := filepath.Join(dstPath, dstFile)

	// Read source file
	srcData, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}

	// Get source FileMode
	srcFileInfo, err := os.Stat(srcFile)

	// Create the destination subdirectories
	dir := filepath.Dir(fullDstPath)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fullDstPath, srcData, srcFileInfo.Mode())
	if err != nil {
		return err
	}

	verboseLog("%s copied in %s\n", srcFile, fullDstPath)

	return nil
}

func main() {
	defaultInstallPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/remogatto/gorgasm-template/")
	installPath := flag.String("install-path", defaultInstallPath, "Package installation directory")
	help := flag.Bool("help", false, "Show usage")
	verbose = flag.Bool("verbose", false, "Be verbose")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "gorgasm-template - Create a template for a basic Gorgasm application\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "\tgorgasm-template [options] dirname\n\n")
		fmt.Fprintf(os.Stderr, "Options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help == true {
		flag.Usage()
		return
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
	}

	dstPath := flag.Arg(0)

	if _, err := os.Stat(dstPath); err == nil {
		panic(fmt.Errorf("Directory %s already exists\n", dstPath))
	}

	templatePath := filepath.Join(*installPath, rootDir)
	err := filepath.Walk(
		templatePath,
		func(src string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				splits := strings.Split(src, "/")
				dstParts := make([]string, len(splits))
				for i := 0; i < len(splits); i++ {
					index := len(splits) - i - 1
					if splits[index] != rootDir {
						continue
					}
					copy(dstParts, splits[index+1:len(splits)])
					break
				}
				err := copyFile(src, dstPath, filepath.Join(dstParts...))
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("A new Gorgasm template was successful created in %s\n", dstPath)
	fmt.Printf("Now:\n\n\tcd %s\n\tgotask init\n\tgotask run xorg # or\n\tgotask run android\n\n", dstPath)
}
