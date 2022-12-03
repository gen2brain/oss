package oss

// Format is audio format.
type Format int

const (
	// AfmtQuery - Returns current fmt.
	AfmtQuery Format = 0x00000000
	// AfmtU8 - U8.
	AfmtU8 Format = 0x00000008
	// AfmtS16Le - Little endian S16.
	AfmtS16Le Format = 0x00000010
	// AfmtS16Be - Big endian S16.
	AfmtS16Be Format = 0x00000020
	// AfmtS8 - S8.
	AfmtS8 Format = 0x00000040
	// AfmtU16Le - Little endian U16.
	AfmtU16Le Format = 0x00000080
	// AfmtU16Be - Big endian U16.
	AfmtU16Be Format = 0x00000100
	// AfmtS32Le - Little endian S32.
	AfmtS32Le Format = 0x00001000
	// AfmtS32Be - Big endian S32.
	AfmtS32Be Format = 0x00002000
	// AfmtFloat - Float.
	AfmtFloat Format = 0x00004000
)

func (f Format) String() string {
	var format string
	switch f {
	case AfmtQuery:
		format = "Query"
	case AfmtU8:
		format = "U8"
	case AfmtS16Le:
		format = "S16Le"
	case AfmtS16Be:
		format = "S16Be"
	case AfmtS8:
		format = "S8"
	case AfmtU16Le:
		format = "U16Le"
	case AfmtU16Be:
		format = "U16Be"
	case AfmtS32Le:
		format = "S32Le"
	case AfmtS32Be:
		format = "S32Be"
	case AfmtFloat:
		format = "Float"
	}

	return format
}

// Cap represents capability.
type Cap int

const (
	// DspCapRevision - Bits for revision level (0 to 255).
	DspCapRevision Cap = 0x000000ff
	// DspCapDuplex - Full duplex record/playback.
	DspCapDuplex Cap = 0x00000100
	// DspCapRealtime - Real time capability.
	DspCapRealtime Cap = 0x00000200
	// DspCapBatch - Audio has some kind of internal buffers which may cause some delays.
	DspCapBatch Cap = 0x00000400
	// DspCapCoproc - Has a coprocessor. Sometimes it's a DSP but usually not.
	DspCapCoproc Cap = 0x00000800
	// DspCapTrigger - Supports SETTRIGGER.
	DspCapTrigger Cap = 0x00001000
	// DspCapMmap - Supports mmap().
	DspCapMmap Cap = 0x00002000
	// DspCapMulti - Supports multiple open.
	DspCapMulti Cap = 0x00004000
	// DspCapBind - Channel binding to front/rear/center/lfe.
	DspCapBind Cap = 0x00008000
)

func (c Cap) String() string {
	var cp string
	switch c {
	case DspCapRevision:
		cp = "Revision"
	case DspCapDuplex:
		cp = "Duplex"
	case DspCapRealtime:
		cp = "Realtime"
	case DspCapBatch:
		cp = "Batch"
	case DspCapCoproc:
		cp = "Coproc"
	case DspCapTrigger:
		cp = "Trigger"
	case DspCapMmap:
		cp = "Mmap"
	case DspCapMulti:
		cp = "Multi"
	case DspCapBind:
		cp = "Bind"
	}

	return cp
}

const (
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
