// Copyright 2024 The Zimtohrli Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// score handles listening test datasets.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/google/zimtohrli/go/data"
	"github.com/google/zimtohrli/go/goohrli"
	"github.com/google/zimtohrli/go/progress"
	"github.com/google/zimtohrli/go/worker"
)

const (
	sampleRate = 48000
)

func main() {
	details := flag.String("details", "", "Path to database directory with a study to show the details from.")
	calculate := flag.String("calculate", "", "Path to a database directory with a study to calculate metrics for.")
	calculateZimtohrli := flag.Bool("calculate_zimtohrli", true, "Whether to calculate Zimtohrli scores.")
	zimtohrliFrequencyResolution := flag.Float64("zimtohrli_frequency_resolution", 8.0, "Smallest bandwidth of of the filterbank.")
	correlate := flag.String("correlate", "", "Path to a database directory with a study to correlate scores for.")
	workers := flag.Int("workers", runtime.NumCPU(), "Number of concurrent workers for tasks.")
	flag.Parse()

	if *details == "" && *calculate == "" && *correlate == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *details != "" {
		study, err := data.OpenStudy(*details)
		if err != nil {
			log.Fatal(err)
		}
		if err := study.ViewEachReference(func(ref *data.Reference) error {
			b, err := json.MarshalIndent(ref, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", b)
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}

	if *calculate != "" {
		study, err := data.OpenStudy(*calculate)
		if err != nil {
			log.Fatal(err)
		}
		bar := progress.New("Calculating")
		pool := &worker.Pool[any]{
			Workers:  *workers,
			OnChange: bar.Update,
		}
		measurements := map[data.ScoreType]data.Measurement{}
		if *calculateZimtohrli {
			z := goohrli.New(sampleRate, *zimtohrliFrequencyResolution)
			z.PerceptualSampleRate = 100.0
			measurements[data.Zimtohrli] = z.NormalizedAudioDistance
		}
		if err := study.Calculate(measurements, pool); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

	if *correlate != "" {
		study, err := data.OpenStudy(*correlate)
		if err != nil {
			log.Fatal(err)
		}
		corrTable, err := study.Correlate()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(corrTable)
	}
}
