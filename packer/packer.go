package packer

import (
	"errors"
	"fmt"
	"slices"
)

type Unpacker struct {
	data []byte
	idx  int
}

func (u *Unpacker) Reset(data []byte) {
	u.data = slices.Clone(data)
	u.idx = 0
}

// first byte of the current buffer
func (u *Unpacker) byte() byte {
	return u.data[u.idx]
}

// consume one byte
func (u *Unpacker) GetByte() (byte, error) {
	if u.RemainingSize() < 1 {
		return 0x00, errors.New("GetByte not enough data")
	}
	b := u.data[u.idx]
	u.idx++
	return b, nil
}

func (u *Unpacker) GetMsgAndSys() (msgId int, system bool, err error) {
	msg := u.GetInt()
	sys := msg&1 != 0
	msg >>= 1
	return msg, sys, nil
}

func (u *Unpacker) GetRaw(size int) ([]byte, error) {
	if size < 0 {
		return nil, fmt.Errorf("GetRaw called with negative size %d", size)
	}
	end := u.idx + size
	if end > u.Size() {
		return nil, fmt.Errorf("GetRaw can not read size %d not enough data", size)
	}
	b := u.data[u.idx:end]
	u.idx += size
	return b, nil
}

// get the full payload from the very beginning
//
// see also:
// - Rest()
// - RemainingData()
func (u *Unpacker) Data() []byte {
	return u.data
}

// consume raw data until the end
//
// see also:
// - Data()
// - RemainingData()
func (u *Unpacker) Rest() []byte {
	rest := u.data[u.idx:]
	u.idx = u.Size()
	return rest
}

// read only operation
// does not consume the data
// if you need to consume the data use Rest() instead
// this method is mostly used for debugging
//
// see also:
// - Rest()
// - Data()
func (u *Unpacker) RemainingData() []byte {
	return u.data[u.idx:]
}

func (u *Unpacker) Size() int {
	return len(u.data)
}

func (u *Unpacker) RemainingSize() int {
	return u.Size() - u.idx
}

const (
	Sanitize                = 1
	SanitizeCC              = 2
	SanitizeSkipWhitespaces = 4
)

func (u *Unpacker) GetStringSanitized(sanitizeType int) (string, error) {
	bytes := []byte{}

	skipping := sanitizeType&SanitizeSkipWhitespaces != 0

	for {
		b, err := u.GetByte()
		if err != nil {
			return "", err
		}
		if b == 0x00 {
			break
		}

		if skipping {
			if b == ' ' || b == '\t' || b == '\n' {
				continue
			}
			skipping = false
		}

		if sanitizeType&SanitizeCC != 0 {
			if b < 32 {
				b = ' '
			}
		} else if sanitizeType&Sanitize != 0 {
			if b < 32 && !(b == '\r') && !(b == '\n') && !(b == '\t') {
				b = ' '
			}
		}

		bytes = append(bytes, b)
	}

	return string(bytes), nil
}

func (u *Unpacker) GetString() (string, error) {
	return u.GetStringSanitized(Sanitize)
}

func (u *Unpacker) GetInt() int {
	sign := int(u.byte()>>6) & 1
	res := int(u.byte() & 0x3F)
	// fake loop should only loop once
	// its the poor mans goto hack
	for {
		if (u.byte() & 0x80) == 0 {
			break
		}
		u.idx += 1
		res |= int(u.byte()&0x7F) << 6

		if (u.byte() & 0x80) == 0 {
			break
		}
		u.idx += 1
		res |= int(u.byte()&0x7F) << (6 + 7)

		if (u.byte() & 0x80) == 0 {
			break
		}
		u.idx += 1
		res |= int(u.byte()&0x7F) << (6 + 7 + 7)

		if (u.byte() & 0x80) == 0 {
			break
		}
		u.idx += 1
		res |= int(u.byte()&0x7F) << (6 + 7 + 7 + 7)
		break
	}

	u.idx += 1
	res ^= -sign
	return res
}

func PackStr(str string) []byte {
	return slices.Concat(
		[]byte(str),
		[]byte{0x00},
	)
}

func PackBool(b bool) []byte {
	if b {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func PackInt(num int) []byte {
	res := []byte{0x00}
	idx := 0
	if num < 0 {
		res[0] |= 0x40 // set sign bit
		num = ^num
	}

	res[0] |= byte(num & 0x3F) // pack 6 bit into dst
	num >>= 6                  // discard 6 bits
	for num != 0 {
		res = append(res, 0x00)

		res[idx] |= 0x80 // set extend bit
		idx++
		res[idx] = byte(num & 0x7F) // pack 7 bit
		num >>= 7                   // discard 7 bits
	}

	return res
}

func UnpackInt(data []byte) int {
	sign := int(data[0]>>6) & 1
	res := int(data[0] & 0x3F)
	i := 0
	// fake loop should only loop once
	// its the poor mans goto hack
	for {
		if (data[i] & 0x80) == 0 {
			break
		}
		i += 1
		res |= int(data[i]&0x7F) << 6

		if (data[i] & 0x80) == 0 {
			break
		}
		i += 1
		res |= int(data[i]&0x7F) << (6 + 7)

		if (data[i] & 0x80) == 0 {
			break
		}
		i += 1
		res |= int(data[i]&0x7F) << (6 + 7 + 7)

		if (data[i] & 0x80) == 0 {
			break
		}
		i += 1
		res |= int(data[i]&0x7F) << (6 + 7 + 7 + 7)
		break
	}

	res ^= -sign
	return res
}

func UnpackMsgAndSys(data []byte) (msgId int, system bool) {
	msg := UnpackInt(data)
	sys := msg&1 != 0
	msg >>= 1
	return msg, sys
}
