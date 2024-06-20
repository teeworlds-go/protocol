package chunk7

// data has to be the uncompressed payload of a teeworlds packet
// without the packet header
//
// It will return all the chunks (messages) in that packet
func UnpackChunks(data []byte) []Chunk {
	chunks := []Chunk{}
	payloadSize := len(data)
	i := 0

	for i < payloadSize {
		chunk := Chunk{}
		chunk.Header.Unpack(data[i:])
		i += 2 // header
		if chunk.Header.Flags.Vital {
			i++
		}
		end := i + chunk.Header.Size
		chunk.Data = make([]byte, end-i)
		copy(chunk.Data[:], data[i:end])
		i += chunk.Header.Size
		chunks = append(chunks, chunk)
	}

	return chunks
}
