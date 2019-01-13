package bullhorn

import (
	"encoding/binary"
	"errors"
	"io"
)

// ErrChecksumLength is returned when a buffer isn't long enough to contain a
// checksum.
var ErrChecksumLength = errors.New(
	"buffer wasn't long enough to contain a checksum")

// ErrInvalidChecksum is reutrned when a buffer's sum doesn't match the
// checksum.
var ErrInvalidChecksum = errors.New("checksum didn't match the buffer")

// ErrLength is returned if the buffer is too long.
var ErrLength = errors.New("buffer is too long")

// CompareSum checks the byte array for a 4-byte checksum in the beginning
// of the array and compares it to the sum of the rest of the bytes.
//
// Returns the bytes without the checksum if the sum is good.
//
// Returns an error if the array doesn't have the right length or of the sum
// doesn't match.
func CompareSum(bs []byte) ([]byte, error) {
	if len(bs) < 4 {
		return nil, ErrChecksumLength
	}
	if BytesToUint32(bs[:4]) != Sum(bs[4:]) {
		return nil, ErrInvalidChecksum
	}
	return bs[4:], nil
}

// Sum of the bytes into a 4 byte checksum.
func Sum(bs []byte) uint32 {
	sum := uint32(0)
	for _, b := range bs {
		sum += uint32(b)
	}
	return sum
}

// Uint32ToBytes converts the uint32 to a big-endian byte-array.
func Uint32ToBytes(x uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, x)
	return bs
}

// BytesToUint32 converts the big-endian byte-array to a uint32.
func BytesToUint32(bs []byte) uint32 {
	return binary.BigEndian.Uint32(bs)
}

// Uint16ToBytes converts the uint16 to a big-endian byte-array.
func Uint16ToBytes(x uint16) []byte {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, x)
	return bs
}

// BytesToUint16 converts the big-endian byte-array to a uint16.
func BytesToUint16(bs []byte) uint16 {
	return binary.BigEndian.Uint16(bs)
}

// AddLength adds the length as an unsigned-short to the beginning of the
// byte-array.
func AddLength(bs []byte) ([]byte, error) {
	if len(bs) > 65536 {
		return nil, ErrLength
	}
	return append(Uint16ToBytes(uint16(len(bs))), bs...), nil
}

// ReadLength reads the first 2 bytes of the io.Reader as a length and then
// reads the length in bytes from the next part of the io.Reader.
//
// Returns an error if the length couldn't be read.
func ReadLength(r io.Reader) ([]byte, error) {
	bs, err := read(2, r)
	if err != nil {
		return nil, err
	}
	length := int(BytesToUint16(bs))
	return read(length, r)
}

func read(tn int, r io.Reader) ([]byte, error) {
	buff := make([]byte, 0, tn)
	temp := make([]byte, tn)
	read := 0
	for read < tn {
		n, err := r.Read(temp)
		if err != nil {
			return nil, err
		}
		buff = append(buff, temp[:n]...)
		read += n
	}
	return buff, nil
}
