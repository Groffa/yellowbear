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
	BytePerSec    uint32 // Number of bytes to read per second (Frequency * BytePerBloc).
	BytePerBloc   uint16 // Number of bytes per block (NbrChannels * BitsPerSample / 8).
	BitsPerSample uint16 // Number of bits per sample
}
