package packer

import "errors"

func PackInt(num int) ([]byte, error) {
	dstLen := 4
	res := []byte{0x00}
	idx := 0
	if num < 0 {
		res[0] |= 0x40 // set sign bit
		num = ^num
	}

	res[0] |= byte(num & 0x3F) // pack 6 bit into dst
	num >>= 6                  // discard 6 bits
	for num != 0 {
		if idx > dstLen {
			return nil, errors.New("Int too big")
		}

		res = append(res, 0x00)

		res[idx] |= 0x80 // set extend bit
		idx++
		res[idx] = byte(num & 0x7F) // pack 7 bit
		num >>= 7                   // discard 7 bits
	}

	return res, nil
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
