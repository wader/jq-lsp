package profile

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func MaybeProfile() func() {
	return Profile(os.Getenv("CPUPROFILE"), os.Getenv("MEMPROFILE"))
}

func Profile(cpuProfilePath string, memProfilePath string) func() {
	var deferFns []func()

	if cpuProfilePath != "" {
		f, err := os.Create(cpuProfilePath)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		} else {
			deferFns = append(deferFns, func() {
				pprof.StopCPUProfile()
				if err := f.Close(); err != nil {
					log.Fatal("could not close CPU profile: ", err)
				}
			})
		}
	}

	return func() {
		for _, fn := range deferFns {
			fn()
		}

		if memProfilePath != "" {
			f, err := os.Create(memProfilePath)
			if err != nil {
				log.Fatal("could not create memory profile: ", err)
			}
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("could not write memory profile: ", err)
			} else {
				runtime.GC() // get up-to-date statistics
				if err := f.Close(); err != nil {
					log.Fatal("could not close memory profile: ", err)
				}
			}
		}
	}
}
