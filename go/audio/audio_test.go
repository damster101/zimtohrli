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

package audio

import (
	"bytes"
	"math"
	"testing"
)

type test_1khz_48khz_wav struct {
	name     string
	data     []byte
	channels int
}

var (
	wavs = []test_1khz_48khz_wav{
		{
			name: "1khz_sine_48khz_rate_1ch_s16le_wav",
			data: []byte{
				0x52, 0x49, 0x46, 0x46, 0xa6, 0x00, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45,
				0x66, 0x6d, 0x74, 0x20, 0x10, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
				0x80, 0xbb, 0x00, 0x00, 0x00, 0x77, 0x01, 0x00, 0x02, 0x00, 0x10, 0x00,
				0x4c, 0x49, 0x53, 0x54, 0x1a, 0x00, 0x00, 0x00, 0x49, 0x4e, 0x46, 0x4f,
				0x49, 0x53, 0x46, 0x54, 0x0e, 0x00, 0x00, 0x00, 0x4c, 0x61, 0x76, 0x66,
				0x36, 0x30, 0x2e, 0x31, 0x36, 0x2e, 0x31, 0x30, 0x30, 0x00, 0x64, 0x61,
				0x74, 0x61, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x02, 0x24, 0x04,
				0x1e, 0x06, 0xff, 0x07, 0xbd, 0x09, 0x4f, 0x0b, 0xb1, 0x0c, 0xda, 0x0d,
				0xc7, 0x0e, 0x73, 0x0f, 0xdc, 0x0f, 0xff, 0x0f, 0xdc, 0x0f, 0x74, 0x0f,
				0xc8, 0x0e, 0xdb, 0x0d, 0xb1, 0x0c, 0x50, 0x0b, 0xbd, 0x09, 0x00, 0x08,
				0x20, 0x06, 0x24, 0x04, 0x17, 0x02, 0x01, 0x00, 0xea, 0xfd, 0xdc, 0xfb,
				0xe2, 0xf9, 0x01, 0xf8, 0x43, 0xf6, 0xb1, 0xf4, 0x4f, 0xf3, 0x26, 0xf2,
				0x39, 0xf1, 0x8d, 0xf0, 0x24, 0xf0, 0x01, 0xf0, 0x24, 0xf0, 0x8c, 0xf0,
				0x38, 0xf1, 0x25, 0xf2, 0x4f, 0xf3, 0xb0, 0xf4, 0x43, 0xf6, 0x00, 0xf8,
				0xe0, 0xf9, 0xdc, 0xfb, 0xe9, 0xfd,
			},
			channels: 1,
		},
		{
			name: "1khz_sine_48khz_rate_2ch_s16le_wav",
			data: []byte{
				0x52, 0x49, 0x46, 0x46, 0x06, 0x01, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45,
				0x66, 0x6d, 0x74, 0x20, 0x10, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x00,
				0x80, 0xbb, 0x00, 0x00, 0x00, 0xee, 0x02, 0x00, 0x04, 0x00, 0x10, 0x00,
				0x4c, 0x49, 0x53, 0x54, 0x1a, 0x00, 0x00, 0x00, 0x49, 0x4e, 0x46, 0x4f,
				0x49, 0x53, 0x46, 0x54, 0x0e, 0x00, 0x00, 0x00, 0x4c, 0x61, 0x76, 0x66,
				0x36, 0x30, 0x2e, 0x31, 0x36, 0x2e, 0x31, 0x30, 0x30, 0x00, 0x64, 0x61,
				0x74, 0x61, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16, 0x02,
				0x16, 0x02, 0x24, 0x04, 0x24, 0x04, 0x1e, 0x06, 0x1e, 0x06, 0xff, 0x07,
				0xff, 0x07, 0xbd, 0x09, 0xbd, 0x09, 0x4f, 0x0b, 0x4f, 0x0b, 0xb1, 0x0c,
				0xb1, 0x0c, 0xda, 0x0d, 0xda, 0x0d, 0xc7, 0x0e, 0xc7, 0x0e, 0x73, 0x0f,
				0x73, 0x0f, 0xdc, 0x0f, 0xdc, 0x0f, 0xff, 0x0f, 0xff, 0x0f, 0xdc, 0x0f,
				0xdc, 0x0f, 0x74, 0x0f, 0x74, 0x0f, 0xc8, 0x0e, 0xc8, 0x0e, 0xdb, 0x0d,
				0xdb, 0x0d, 0xb1, 0x0c, 0xb1, 0x0c, 0x50, 0x0b, 0x50, 0x0b, 0xbd, 0x09,
				0xbd, 0x09, 0x00, 0x08, 0x00, 0x08, 0x20, 0x06, 0x20, 0x06, 0x24, 0x04,
				0x24, 0x04, 0x17, 0x02, 0x17, 0x02, 0x01, 0x00, 0x01, 0x00, 0xea, 0xfd,
				0xea, 0xfd, 0xdc, 0xfb, 0xdc, 0xfb, 0xe2, 0xf9, 0xe2, 0xf9, 0x01, 0xf8,
				0x01, 0xf8, 0x43, 0xf6, 0x43, 0xf6, 0xb1, 0xf4, 0xb1, 0xf4, 0x4f, 0xf3,
				0x4f, 0xf3, 0x26, 0xf2, 0x26, 0xf2, 0x39, 0xf1, 0x39, 0xf1, 0x8d, 0xf0,
				0x8d, 0xf0, 0x24, 0xf0, 0x24, 0xf0, 0x01, 0xf0, 0x01, 0xf0, 0x24, 0xf0,
				0x24, 0xf0, 0x8c, 0xf0, 0x8c, 0xf0, 0x38, 0xf1, 0x38, 0xf1, 0x25, 0xf2,
				0x25, 0xf2, 0x4f, 0xf3, 0x4f, 0xf3, 0xb0, 0xf4, 0xb0, 0xf4, 0x43, 0xf6,
				0x43, 0xf6, 0x00, 0xf8, 0x00, 0xf8, 0xe0, 0xf9, 0xe0, 0xf9, 0xdc, 0xfb,
				0xdc, 0xfb, 0xe9, 0xfd, 0xe9, 0xfd,
			},
			channels: 2,
		},
	}
)

func readWAVTest(data []byte, wantChannels int) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		w, err := ReadWAV(bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
		if w.FormatChunk.NumChannels != int16(wantChannels) {
			t.Errorf("got %v channels, want %v", w.FormatChunk.NumChannels, wantChannels)
		}
		if w.FormatChunk.SampleRate != 48000 {
			t.Errorf("got sample rate %v, want %v", w.FormatChunk.SampleRate, 48000)
		}
		sampleRateReciporcal := 1.0 / float64(w.FormatChunk.SampleRate)
		audio, err := w.Audio()
		if err != nil {
			t.Fatal(err)
		}
		audio.Amplify(1.0 / audio.MaxAbsAmplitude)
		for _, channel := range audio.Samples {
			for sampleIndex, sample := range channel {
				wantSample := math.Sin(2 * math.Pi * 1000 * sampleRateReciporcal * float64(sampleIndex))
				if math.Abs(wantSample-float64(sample)) > 1e-3 {
					t.Errorf("got sample %v %v, want %v", sampleIndex, sample, wantSample)
				}
			}
		}

	}
}

func TestReadWAV(t *testing.T) {
	for _, w := range wavs {
		t.Run(w.name, readWAVTest(w.data, w.channels))
	}
}
