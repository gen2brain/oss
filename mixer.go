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

var (
	soundMixerReadDevmask = _ior(77, soundMixerDevmask)

	soundMixerReadVolume  = _ior(77, SoundMixerVolume)
	soundMixerReadBass    = _ior(77, SoundMixerBass)
	soundMixerReadTreble  = _ior(77, SoundMixerTreble)
	soundMixerReadSynth   = _ior(77, SoundMixerSynth)
	soundMixerReadPcm     = _ior(77, SoundMixerPcm)
	soundMixerReadSpeaker = _ior(77, SoundMixerSpeaker)
	soundMixerReadLine    = _ior(77, SoundMixerLine)
	soundMixerReadMic     = _ior(77, SoundMixerMic)
	soundMixerReadCd      = _ior(77, SoundMixerCd)
	soundMixerReadImix    = _ior(77, SoundMixerImix)
	soundMixerReadAltpcm  = _ior(77, SoundMixerAltpcm)
	soundMixerReadReclev  = _ior(77, SoundMixerReclev)
	soundMixerReadIgain   = _ior(77, SoundMixerIgain)
	soundMixerReadOgain   = _ior(77, SoundMixerOgain)
	soundMixerReadLine1   = _ior(77, SoundMixerLine1)
	soundMixerReadLine2   = _ior(77, SoundMixerLine2)
	soundMixerReadLine3   = _ior(77, SoundMixerLine3)

	soundMixerWriteVolume  = _iowr(77, SoundMixerVolume)
	soundMixerWriteBass    = _iowr(77, SoundMixerBass)
	soundMixerWriteTreble  = _iowr(77, SoundMixerTreble)
	soundMixerWriteSynth   = _iowr(77, SoundMixerSynth)
	soundMixerWritePcm     = _iowr(77, SoundMixerPcm)
	soundMixerWriteSpeaker = _iowr(77, SoundMixerSpeaker)
	soundMixerWriteLine    = _iowr(77, SoundMixerLine)
	soundMixerWriteMic     = _iowr(77, SoundMixerMic)
	soundMixerWriteCd      = _iowr(77, SoundMixerCd)
	soundMixerWriteImix    = _iowr(77, SoundMixerImix)
	soundMixerWriteAltpcm  = _iowr(77, SoundMixerAltpcm)
	soundMixerWriteReclev  = _iowr(77, SoundMixerReclev)
	soundMixerWriteIgain   = _iowr(77, SoundMixerIgain)
	soundMixerWriteOgain   = _iowr(77, SoundMixerOgain)
	soundMixerWriteLine1   = _iowr(77, SoundMixerLine1)
	soundMixerWriteLine2   = _iowr(77, SoundMixerLine2)
	soundMixerWriteLine3   = _iowr(77, SoundMixerLine3)
)
