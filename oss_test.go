package oss_test

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
	"time"

	"github.com/gen2brain/oss"
)

func TestAudio(t *testing.T) {
	dev, err := oss.OpenAudio()
	if err != nil {
		t.Fatal(err)
	}

	err = dev.SetBufferSize(2048, 2)
	if err != nil {
		t.Error(err)
	}

	caps, err := dev.Caps()
	if err != nil {
		t.Error(err)
	}

	if caps&int(oss.DspCapTrigger) == 0 {
		t.Errorf("Controls: no trigger cap")
	}

	formats, err := dev.Formats()
	if err != nil {
		t.Error(err)
	}

	if formats&int(oss.AfmtS16Le) == 0 {
		t.Errorf("Formats: no S16Le format")
	}

	channels, err := dev.Channels(2)
	if err != nil {
		t.Error(err)
	}

	if channels != 2 {
		t.Errorf("Channels: got %d, want %d", channels, 2)
	}

	speed := 44100
	samplerate, err := dev.Samplerate(speed)
	if err != nil {
		t.Error(err)
	}

	if samplerate != speed {
		t.Errorf("Samplerate: got %d, want %d", samplerate, speed)
	}

	reqFormat := oss.AfmtS16Le
	format, err := dev.Format(reqFormat)
	if err != nil {
		t.Error(err)
	}

	if format != reqFormat {
		t.Errorf("Format: got %d, want %d", format, reqFormat)
	}

	err = dev.Start()
	if err != nil {
		t.Error(err)
	}

	// Play 1 second of beep.
	duration := 1 * time.Second
	for x := 0.; x < duration.Seconds(); {
		var buf bytes.Buffer

		for i := 0; i < 2048; i++ {
			v := math.Sin(x * 2 * math.Pi * 440)
			v *= 0.1 // make it a little quieter

			sample := int16(v * math.MaxInt16)

			for c := 0; c < channels; c++ {
				err = binary.Write(&buf, binary.LittleEndian, sample)
				if err != nil {
					t.Error(err)
				}
			}

			x += 1 / float64(samplerate)
		}

		_, err := dev.Write(buf.Bytes())
		if err != nil {
			t.Error(err)
		}
	}

	playing := dev.Playing()
	if !playing {
		t.Errorf("Playing: not playing")
	}

	_, err = dev.BufferSize()
	if err != nil {
		t.Error(err)
	}

	_, err = dev.Delay()
	if err != nil {
		t.Error(err)
	}

	_, err = dev.UnplayedBufferSize()
	if err != nil {
		t.Error(err)
	}

	err = dev.Sync()
	if err != nil {
		t.Error(err)
	}

	err = dev.Reset()
	if err != nil {
		t.Error(err)
	}

	err = dev.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestMixer(t *testing.T) {
	mix, err := oss.OpenMixer()
	if err != nil {
		t.Fatal(err)
	}

	caps, err := mix.Controls()
	if err != nil {
		t.Error(err)
	}

	if caps&(1<<oss.SoundMixerVolume) == 0 {
		t.Errorf("Controls: no Master control")
	}

	reqLeft, reqRight := 50, 50
	curLeft, curRight, err := mix.Volume(oss.SoundMixerVolume)
	if err != nil {
		t.Error(err)
	}

	left, right, err := mix.SetVolume(oss.SoundMixerVolume, reqLeft, reqRight)
	if err != nil {
		t.Error(err)
	}

	if left != reqLeft || right != reqRight {
		t.Errorf("SetVolume: got %d,%d want %d,%d", left, right, reqLeft, reqRight)
	}

	_, _, err = mix.SetVolume(oss.SoundMixerVolume, curLeft, curRight)
	if err != nil {
		t.Error(err)
	}

	err = mix.Close()
	if err != nil {
		t.Error(err)
	}
}
