package message

import (
	"encoding/binary"
	"math"
	"time"
)

// bytesToInt converts the bytes to an int64.
func bytesToInt(bs []byte) uint64 {
	return binary.BigEndian.Uint64(bs)
}

// intToBytes converts the int64 to bytes.
func intToBytes(x int64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(x))
	return bs
}

// bytesToFloat converts the bytes to a float32.
func bytesToFloat(bs []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(bs))
}

// floatToBytes converts the float32 to bytes.
func floatToBytes(x float32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, math.Float32bits(x))
	return bs
}

// timeToInt converts the time.Time to an int64.
func timeToInt(t time.Time) int64 {
	return t.Unix()
}

// intToTime converts the int64 to a time.Time.
func intToTime(x int64) time.Time {
	return time.Unix(x, 0)
}
