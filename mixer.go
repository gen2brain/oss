package oss

// Control represents mixer control.
type Control int

const (
	// SoundMixerVolume - Master volume.
	SoundMixerVolume Control = 0
	// SoundMixerBass - Bass.
	SoundMixerBass Control = 1
	// SoundMixerTreble - Treble.
	SoundMixerTreble Control = 2
	// SoundMixerSynth - Synth.
	SoundMixerSynth Control = 3
	// SoundMixerPcm - PCM.
	SoundMixerPcm Control = 4
	// SoundMixerSpeaker - Speaker.
	SoundMixerSpeaker Control = 5
	// SoundMixerLine - Line.
	SoundMixerLine Control = 6
	// SoundMixerMic - Mic.
	SoundMixerMic Control = 7
	// SoundMixerCd - CD.
	SoundMixerCd Control = 8
	// SoundMixerImix - Recording monitor.
	SoundMixerImix Control = 9
	// SoundMixerAltpcm - Alt PCM.
	SoundMixerAltpcm Control = 10
	// SoundMixerReclev - Recording level.
	SoundMixerReclev Control = 11
	// SoundMixerIgain - Input gain.
	SoundMixerIgain Control = 12
	// SoundMixerOgain - Output gain.
	SoundMixerOgain Control = 13
	// SoundMixerLine1 - Input source 1  (aux1).
	SoundMixerLine1 Control = 14
	// SoundMixerLine2 - Input source 2  (aux2).
	SoundMixerLine2 Control = 15
	// SoundMixerLine3 - Input source 3  (line).
	SoundMixerLine3 Control = 16
)

func (c Control) String() string {
	var control string
	switch c {
	case SoundMixerVolume:
		control = "Volume"
	case SoundMixerBass:
		control = "Bass"
	case SoundMixerTreble:
		control = "Treble"
	case SoundMixerSynth:
		control = "Synth"
	case SoundMixerPcm:
		control = "Pcm"
	case SoundMixerSpeaker:
		control = "Speaker"
	case SoundMixerLine:
		control = "Line"
	case SoundMixerMic:
		control = "Mic"
	case SoundMixerCd:
		control = "Cd"
	case SoundMixerImix:
		control = "Imix"
	case SoundMixerAltpcm:
		control = "Altpcm"
	case SoundMixerReclev:
		control = "Reclev"
	case SoundMixerIgain:
		control = "Igain"
	case SoundMixerOgain:
		control = "Ogain"
	case SoundMixerLine1:
		control = "Line1"
	case SoundMixerLine2:
		control = "Line2"
	case SoundMixerLine3:
		control = "Line3"
	}

	return control
}

const (
	soundMixerDevmask = 0xfe
)

func soundMixerRead(control Control) int {
	return _ior(77, int(control))
}

func soundMixerWrite(control Control) int {
	return _iowr(77, int(control))
}

var (
	soundMixerReadDevmask = _ior(77, soundMixerDevmask)
)
