package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/industria/unused-file-remover/filesystem"
)

func main() {
	var cacheLocation = flag.String("cachePath", "", "Path where binaries are cached on disk")
	var maxAgeArg = flag.String("maxage", "168h", "Maximum age, as Go duration format, file access time must have before it is deleted (default: 168h 7 days)")
	var verbose = flag.Bool("verbose", false, "Verbose output")

	flag.Usage = func() {
		log.Printf("Usage: %s arguments\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *cacheLocation == "" {
		log.Fatalf("cachePath argument is required")
	}

	maxAge, err := time.ParseDuration(*maxAgeArg)
	if err != nil {
		log.Fatalf("illegal format of the age argument %s", *maxAgeArg)
	}

	if *verbose {
		log.Printf("scanning cacheLocation: %s with maximum age %s\n", *cacheLocation, maxAge)
	}

	filepath.WalkDir(*cacheLocation, createWalk(maxAge, *verbose))
}

func createWalk(maxAge time.Duration, verbose bool) fs.WalkDirFunc {
	now := time.Now().UTC()
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if d.Type().IsRegular() {
			fi, err := d.Info()
			if err != nil {
				log.Printf("failed to stat file %s : %v", path, err)
				return nil // just move on
			}

			if stat, ok := fi.Sys().(*syscall.Stat_t); ok {
				atim := filesystem.Atim(stat)
				age := now.Sub(atim)
				if age > maxAge {
					if err := os.Remove(path); err != nil {
						log.Printf("failed to remove file %s : %v", path, err)
					} else {
						if verbose {
							log.Printf("removed file %s age %v atim: %s", path, age, atim.Format(time.RFC3339))
						}
					}
				}
			}
		}
		return nil
	}
}
