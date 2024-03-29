package lib

import "strconv"

func F정수_쉼표_추가[T T숫자](값 T) string {
	return FormatInt(int64(값), 3, ',')
}

// COPIED FROM 'github.com/icza/gox/fmtx'
func FormatInt(n int64, groupSize int, grouping byte) string {
	if groupSize < 1 {
		groupSize = 3
	}

	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / groupSize

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == groupSize {
			j, k = j-1, 0
			out[j] = grouping
		}
	}
}
