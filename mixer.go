package oss

const (
	// SoundMixerVolume - Master volume.
	SoundMixerVolume = 0
	// SoundMixerBass - Bass.
	SoundMixerBass = 1
	// SoundMixerTreble - Treble.
	SoundMixerTreble = 2
	// SoundMixerSynth - Synth.
	SoundMixerSynth = 3
	// SoundMixerPcm - PCM.
	SoundMixerPcm = 4
	// SoundMixerSpeaker - Speaker.
	SoundMixerSpeaker = 5
	// SoundMixerLine - Line.
	SoundMixerLine = 6
	// SoundMixerMic - Mic.
	SoundMixerMic = 7
	// SoundMixerCd - CD.
	SoundMixerCd = 8
	// SoundMixerImix - Recording monitor.
	SoundMixerImix = 9
	// SoundMixerAltpcm - Alt PCM.
	SoundMixerAltpcm = 10
	// SoundMixerReclev - Recording level.
	SoundMixerReclev = 11
	// SoundMixerIgain - Input gain.
	SoundMixerIgain = 12
	// SoundMixerOgain - Output gain.
	SoundMixerOgain = 13
	// SoundMixerLine1 - Input source 1  (aux1).
	SoundMixerLine1 = 14
	// SoundMixerLine2 - Input source 2  (aux2).
	SoundMixerLine2 = 15
	// SoundMixerLine3 - Input source 3  (line).
	SoundMixerLine3 = 16

	soundMixerDevmask = 0xfe
)

func soundMixerRead(control int) int {
	return _ior(77, control)
}

func soundMixerWrite(control int) int {
	return _iowr(77, control)
}

var (
	soundMixerReadDevmask = _ior(77, soundMixerDevmask)
)
