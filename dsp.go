package oss

const (
	// AfmtQuery - Returns current fmt.
	AfmtQuery = 0x00000000
	// AfmtU8 - U8.
	AfmtU8 = 0x00000008
	// AfmtS16Le - Little endian S16.
	AfmtS16Le = 0x00000010
	// AfmtS16Be - Big endian S16.
	AfmtS16Be = 0x00000020
	// AfmtS8 - S8.
	AfmtS8 = 0x00000040
	// AfmtU16Le - Little endian U16.
	AfmtU16Le = 0x00000080
	// AfmtU16Be - Big endian U16.
	AfmtU16Be = 0x00000100
	// AfmtS32Le - Little endian S32.
	AfmtS32Le = 0x00001000
	// AfmtS32Be - Big endian S32.
	AfmtS32Be = 0x00002000
	// AfmtFloat - Float.
	AfmtFloat = 0x00004000

	// DspCapRevision - Bits for revision level (0 to 255).
	DspCapRevision = 0x000000ff
	// DspCapDuplex - Full duplex record/playback.
	DspCapDuplex = 0x00000100
	// DspCapRealtime - Real time capability.
	DspCapRealtime = 0x00000200
	// DspCapBatch - Audio has some kind of internal buffers which may cause some delays.
	DspCapBatch = 0x00000400
	// DspCapCoproc - Has a coprocessor. Sometimes it's a DSP but usually not.
	DspCapCoproc = 0x00000800
	// DspCapTrigger - Supports SETTRIGGER.
	DspCapTrigger = 0x00001000
	// DspCapMmap - Supports mmap().
	DspCapMmap = 0x00002000
	// DspCapMulti - Supports multiple open.
	DspCapMulti = 0x00004000
	// DspCapBind - Channel binding to front/rear/center/lfe.
	DspCapBind = 0x00008000

	pcmEnableOutput = 0x00000002
)

type audioBufInfo struct {
	// Number of available fragments (partially used ones not counted)
	fragments int32
	// Total number of fragments allocated
	fragstotal int32
	// Size of a fragment in bytes
	fragsize int32

	// Available space in bytes (includes partially used fragments)
	// Note! 'bytes' could be more than fragments*fragsize
	bytes int32
}

var (
	sndctlDspReset       = _io(80, 0)
	sndctlDspSync        = _io(80, 1)
	sndctlDspSpeed       = _iowr(80, 2)
	sndctlDspSetfmt      = _iowr(80, 5)
	sndctlDspChannels    = _iowr(80, 6)
	sndctlDspSetfragment = _iowr(80, 10)
	sndctlDspGetfmts     = _ior(80, 11)
	sndctlDspGetospace   = _iori(80, 12)
	sndctlDspNonblock    = _io(80, 14)
	sndctlDspGetcaps     = _ior(80, 15)
	sndctlDspSettrigger  = _ior(80, 16)
	sndctlDspGetodelay   = _ior(80, 23)
)
