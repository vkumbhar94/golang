package main

// var glob = "*"
//
// func index(s, pattern string) int {
// 	if len(pattern) > len(s) {
// 		return -1
// 	}
//
// 	patternLen := len(pattern)
// 	for i := 0; i < len(s)-patternLen+1; i++ {
// 		ss := s[i : i+patternLen]
// 		if strings.EqualFold(ss, pattern) {
// 			return i
// 		}
// 	}
// 	return -1
// }
//
// func hasSuffix(s, pattern string) bool {
// 	if len(pattern) > len(s) {
// 		return false
// 	}
// 	return index(s[len(s)-len(pattern):], pattern) == 0
// }
//
// func GlobBytes(globParts [][]byte, s []byte) bool {
// 	// if globParts is 0 we assume that query is *
// 	if len(globParts) == 0 {
// 		return true
// 	}
//
// 	remaining := s
// 	for _, part := range globParts {
// 		remaining = trimPart(part, remaining)
// 		if remaining == nil {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func globParts(p string) [][]byte {
// 	s := bytes.Split([]byte(p), []byte(glob))
// 	res := [][]byte{}
// 	for _, n := range s {
// 		if len(n) > 0 {
// 			res = append(res, n)
// 		}
// 	}
// 	return res
// }
//
// func extractRune(s []byte, i int) (rune, int) {
// 	if s[i] < utf8.RuneSelf {
// 		return rune(s[i]), 1
// 	}
// 	return utf8.DecodeRune(s[i:])
// }
//
// func trimPart(part, s []byte) []byte {
// 	if len(part) == 0 {
// 		return s
// 	}
//
// 	lower, pstart := extractRune(part, 0)
// 	upper := unicode.ToUpper(lower)
//
// 	for i := 0; i < len(s); {
// 		r, size := extractRune(s, i)
// 		i += size
// 		if r != lower && r != upper {
// 			continue
// 		}
// 		if pstart >= len(part) {
// 			return s[i:]
// 		}
// 		potentialStart := false
// 		si, pi := i, pstart
// 		for si < len(s) && pi < len(part) {
// 			r, size = extractRune(s, si)
// 			si += size
// 			if !potentialStart {
// 				i = si
// 				potentialStart = r == lower || r == upper
// 			}
//
// 			pr, psize := extractRune(part, pi)
// 			pi += psize
// 			if pr == r || unicode.ToUpper(pr) == r {
// 				if pi >= len(part) {
// 					return s[si:]
// 				}
// 				continue
// 			}
// 			if !potentialStart {
// 				break
// 			}
// 			potentialStart = false
// 			si, pi = i, pstart
// 		}
// 	}
// 	return nil
// }
//
// func hasSuffixBytes(s, pattern []byte) bool {
// 	if len(pattern) > len(s) {
// 		return false
// 	}
// 	return bytes.EqualFold(s[len(s)-len(pattern):], pattern)
// }
