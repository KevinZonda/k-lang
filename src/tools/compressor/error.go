package compressor

type BaseCompressorError struct{ msg string }

type CompressorError *BaseCompressorError

func (e *BaseCompressorError) Error() string {
	return e.msg
}

func newError(msg string) CompressorError {
	return &BaseCompressorError{msg: msg}
}

var ErrNotCompressed = newError("not a compressed file")

var ErrDecompressFailed = newError("decompress failed")

var ErrDeserialiseFailed = newError("deserialise failed")

var ErrSerialiseFailed = newError("serialise failed")
