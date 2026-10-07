package wav

type RIFFChunk struct {
	FileTypeBlocID uint32 // Identifier « RIFF »  (0x52, 0x49, 0x46, 0x46)
	FileSize       uint32 // Overall file size minus 8 bytes
	FileFormatID   uint32 // Format = « WAVE »  (0x57, 0x41, 0x56, 0x45)
}

func NewRIFFChunk() RIFFChunk {
	return RIFFChunk{
		FileTypeBlocID: 0x52494646,
		FileSize:       0, // To be filled in later
		FileFormatID:   0x57415645,
	}
}

type DataFormatChunk struct {
	FormatBlocID  uint32 // Identifier « fmt␣ »  (0x66, 0x6D, 0x74, 0x20)
	BlocSize      uint32 // Chunk size minus 8 bytes, which is 16 bytes here  (0x10)
	AudioFormat   uint16 // Audio format (1: PCM integer, 3: IEEE 754 float)
	NbrChannels   uint16 // Number of channels
	Frequency     uint32 // Sample rate (in hertz)
	BitsPerSample uint16 // Number of bits per sample
	BytePerBloc   uint16 // Number of bytes per block (NbrChannels * BitsPerSample / 8).
	BytePerSec    uint32 // Number of bytes to read per second (Frequency * BytePerBloc).
}

const BitsPerSample uint16 = 16

func NewDataFormatChunk(channels uint16, sampleRate uint32) DataFormatChunk {
	bytePerBlock := channels * BitsPerSample / 8
	return DataFormatChunk{
		FormatBlocID:  0x666D7420,
		BlocSize:      0, // To be set by caller later
		AudioFormat:   1,
		NbrChannels:   channels,
		BitsPerSample: BitsPerSample,
		Frequency:     sampleRate,
		BytePerBloc:   bytePerBlock,
		BytePerSec:    sampleRate * uint32(bytePerBlock),
	}
}

func NewMonoDataFormatChunk() DataFormatChunk {
	return NewDataFormatChunk(1, 44100)
}

type SampleDataChunk struct {
	DataBlocID uint32 // Identifier « data »  (0x64, 0x61, 0x74, 0x61)
	DataSize   uint32 // SampledData size
	// SampledData...
}

func NewSampleDataChunk(dataSize uint32) SampleDataChunk {
	return SampleDataChunk{
		DataBlocID: 0x64617461,
		DataSize:   dataSize,
	}
}
