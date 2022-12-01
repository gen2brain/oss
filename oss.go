// Package oss provides access to OSS (Open Sound System) audio interface.
package oss

import (
	"os"
)

// Audio represents audio output device.
type Audio struct {
	file    *os.File
	playing bool
}

// OpenAudio opens the audio output device.
func OpenAudio(path ...string) (*Audio, error) {
	dev := &Audio{}

	var devPath string
	env := os.Getenv("OSS_AUDIODEV")
	if env != "" {
		devPath = env
	}
	env = os.Getenv("AUDIODEV")
	if env != "" {
		devPath = env
	}
	if len(path) > 0 {
		devPath = path[0]
	}
	if devPath == "" {
		devPath = "/dev/dsp"
	}

	file, err := os.OpenFile(devPath, os.O_WRONLY, 0664)
	if err != nil {
		return nil, err
	}

	dev.file = file

	return dev, nil
}

// Close closes the audio output device.
func (a *Audio) Close() error {
	a.playing = false

	return a.file.Close()
}

// Caps returns the capability mask of an audio device.
func (a *Audio) Caps() (int, error) {
	value, err := ioctl(a.file.Fd(), sndctlDspGetcaps, 0)

	return value, err
}

// Formats returns the supported formats mask.
func (a *Audio) Formats() (int, error) {
	value, err := ioctl(a.file.Fd(), sndctlDspGetfmts, 0)

	return value, err
}

// Playing checks if audio is playing.
func (a *Audio) Playing() bool {
	return a.playing
}

// Nonblock forces non-blocking mode regardless of the O_NONBLOCK file mode flag.
func (a *Audio) Nonblock() (int, error) {
	value, err := ioctl(a.file.Fd(), sndctlDspNonblock, 0)

	return value, err
}

// Format sets the audio format.
func (a *Audio) Format(format int) (int, error) {
	value, err := ioctl(a.file.Fd(), sndctlDspSetfmt, format)

	return value, err
}

// Samplerate sets the audio samplerate (speed).
//
// In some cases it's not possible to support the samplerate requested.
// In that case the nearest possible sampling rate will be used.
func (a *Audio) Samplerate(samplerate int) (int, error) {
	return ioctl(a.file.Fd(), sndctlDspSpeed, samplerate)
}

// Channels sets the number of audio channels.
//
// The returned value may be different from the one requested.
func (a *Audio) Channels(channels int) (int, error) {
	return ioctl(a.file.Fd(), sndctlDspChannels, channels)
}

// SetBufferSize sets the buffer size hint.
// This call just sets the size hint, there is no guarantee that the requested size gets used.
//
// fragments is the maximum number of fragments to be allocated. It can be between 2 and 0x7fff (unlimited).
// The actual number of fragments may be less than requested if there is not enough memory available.
//
// fragsize is the fragment size selector, i.e. for selector N the size is 1<<N.
func (a *Audio) SetBufferSize(fragments, fragsize int) error {
	arg := (fragments << 16) | fragsize
	_, err := ioctl(a.file.Fd(), sndctlDspSetfragment, arg)

	return err
}

// BufferSize returns the buffer size.
func (a *Audio) BufferSize() (int, error) {
	info, err := ioctlI(a.file.Fd(), sndctlDspGetospace)
	if err != nil {
		return 0, err
	}

	return int(info.fragstotal * info.fragsize), err
}

// UnplayedBufferSize returns the size of buffer that is not yet played.
func (a *Audio) UnplayedBufferSize() (int, error) {
	info, err := ioctlI(a.file.Fd(), sndctlDspGetospace)
	if err != nil {
		return 0, err
	}

	return int(info.fragstotal*info.fragsize - info.bytes), err
}

// Delay returns the playback buffering delay.
//
// This call tells how long it's going to take before the next sample to be written gets played by the hardware.
// The delay is returned in bytes, it can be converted to seconds by dividing it by:
//
//	samplerate*channels*sampleSizeInBytes.
func (a *Audio) Delay() (int, error) {
	return ioctl(a.file.Fd(), sndctlDspGetodelay, 0)
}

// Start starts audio playback.
//
// Note that the Write call will automatically start the actual playback operation.
func (a *Audio) Start() error {
	_, err := ioctl(a.file.Fd(), sndctlDspSettrigger, pcmEnableOutput)
	if err == nil {
		a.playing = true
	}

	return err
}

// Write queues one audio buffer to the hardware.
func (a *Audio) Write(b []byte) (int, error) {
	n, err := a.file.Write(b)
	if err != nil {
		a.playing = false
	}

	return n, err
}

// Reset aborts audio playback operation.
//
// There is no need to use this call after opening the device.
// The Open operation automatically initializes the device.
func (a *Audio) Reset() error {
	a.playing = false

	_, err := ioctl(a.file.Fd(), sndctlDspReset, 0)
	if err != nil {
		return err
	}

	_, err = ioctl(a.file.Fd(), sndctlDspSettrigger, 0)

	return err
}

// Sync can be used to wait until all samples written to the audio device have been played.
//
// The sync operation will always cause a more or less noticeable pause or click in the output.
// It is recommended to output few samples of silence before and after making this call.
//
// Closing the device will perform the sync operation automatically so using this ioctl call
// is unnecessary before calling Close.
func (a *Audio) Sync() error {
	_, err := ioctl(a.file.Fd(), sndctlDspSync, 0)

	return err
}

// Mixer represents mixer device.
type Mixer struct {
	file *os.File
}

// OpenMixer opens the mixer device.
func OpenMixer(path ...string) (*Mixer, error) {
	dev := &Mixer{}

	var devPath string
	env := os.Getenv("OSS_MIXERDEV")
	if env != "" {
		devPath = env
	}
	env = os.Getenv("MIXERDEV")
	if env != "" {
		devPath = env
	}
	if len(path) > 0 {
		devPath = path[0]
	}
	if devPath == "" {
		devPath = "/dev/mixer"
	}

	file, err := os.OpenFile(devPath, os.O_RDWR, 0664)
	if err != nil {
		return nil, err
	}

	dev.file = file

	return dev, nil
}

// Close closes the mixer device.
func (m *Mixer) Close() error {
	return m.file.Close()
}

// Controls returns a bitmask specifying all available mixer controls.
//
// To determine if, for example, mixer supports a PCM control:
//
//	if controls & (1 << oss.SoundMixerPcm) != 0 { // true
func (m *Mixer) Controls() (int, error) {
	value, err := ioctl(m.file.Fd(), soundMixerReadDevmask, 0)

	return value, err
}

// SetVolume sets the volume for a given mixer control.
func (m *Mixer) SetVolume(control, leftVol, rightVol int) (int, int, error) {
	volume := (rightVol << 8) | leftVol

	value, err := ioctl(m.file.Fd(), soundMixerWrite(control), volume)

	left := value & 0xff
	right := (value & 0xff00) >> 8

	return left, right, err
}

// Volume returns the volume of a given mixer control.
func (m *Mixer) Volume(control int) (int, int, error) {
	value, err := ioctl(m.file.Fd(), soundMixerRead(control), 0)

	left := value & 0xff
	right := (value & 0xff00) >> 8

	return left, right, err
}
