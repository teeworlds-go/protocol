package packer

import (
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
func (u *Unpacker) getByte() byte {
	b := u.data[u.idx]
	u.idx++
	return b
}

func (u *Unpacker) Data() []byte {
	return u.data
}

const (
	Sanitize                = 1
	SanitizeCC              = 2
	SanitizeSkipWhitespaces = 4
)

func (u *Unpacker) GetStringSanitized(sanitizeType int) string {
	bytes := []byte{}

	skipping := sanitizeType&SanitizeSkipWhitespaces != 0

	for {
		b := u.getByte()
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

	return string(bytes)
}

func (u *Unpacker) GetString() string {
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
