package keyvaluestore

import "io"

// codex: encode/decode record for append-only file

// limits
/*
maxKeyLen = 256
maxValLen = 64*1024 = 64KB
*/

// errors
/*
ErrInvalidLength
ErrChecksum
ErrEOF
*/

// Encode

func Encode(key, value []byte) (record []byte, err error) {
	// validate len
	// allocate buffer
	// write key_len/val_len (little-endian)
	// write key/value
	// compute  CRC32
	// write chksum
	// return record

}

// Decode

func Decode(r io.Reader) (key []byte, value []byte, bytesRead int, err error) {
	// read key_len
	// read val_len
	// validate len
	// read key
	// read val
	// read chksum
	// compute CRC32 & compare
	// return key/val/bytesRead
}
