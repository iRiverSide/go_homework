package goTaskOne

func reverseString(s []byte) {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		changeByte := s[sLen-i]
		s[sLen-i] = s[i]
		s[i] = changeByte
	}
}
