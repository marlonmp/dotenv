package dotenv

type reader struct {
	size, bookmark int
	buff           []byte

	exceptionFounded bool
}

func newReader(b []byte) *reader {
	return &reader{
		size: len(b),
		buff: b,
	}
}

// get the variable name
func (R *reader) readIdx() []byte {
	var bytesRead []byte

	i := R.bookmark

Loop:
	for ; i < R.size; i++ {
		switch R.buff[i] {
		// end of idx
		case '=':
			R.exceptionFounded = true
			break Loop
		// end of line
		case '\n':
			R.exceptionFounded = false
			break Loop
		// skip spaces and tabs
		case ' ', '\t':
			continue Loop
		}
		// add byte if not match
		bytesRead = append(bytesRead, R.buff[i])
	}

	R.bookmark = i + 1

	return bytesRead
}

// get the data in the quotes (", ', `)
func (R *reader) readQuotes(i int) ([]byte, int) {
	var bytesRead []byte

	quote := R.buff[i]

	i++

Loop:

	for ; i < R.size; i++ {
		switch R.buff[i] {
		case '\\':
			if i+1 < R.size {
				i++
			}
		case quote:
			break Loop
		}

		bytesRead = append(bytesRead, R.buff[i])
	}

	return bytesRead, i
}

// get the data after the variable name
func (R *reader) readVal() []byte {
	var bytesRead []byte

	i := R.bookmark

Loop:

	for ; i < R.size; i++ {
		switch R.buff[i] {
		case '\\':
			if i+1 < R.size {
				i++
			}
		case '"', '\'', '`':
			bytesRead, i = R.readQuotes(i)
			break Loop
		case ' ', '\t':
			continue Loop
		case '\n':
			break Loop
		}

		bytesRead = append(bytesRead, R.buff[i])
	}

	R.bookmark = i + 1

	return bytesRead
}
