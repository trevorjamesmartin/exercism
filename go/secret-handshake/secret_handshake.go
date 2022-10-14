package secret

var STEPS = map[int]string{1: "wink", 2: "double blink", 4: "close your eyes", 8: "jump"}

func putright(key int, a []string) []string {
	b := STEPS[key]
	return append(a, b)
}

func putleft(key int, a []string) []string {
	b := STEPS[key]
	return append([]string{b}, a...)
}

func Handshake(code uint) []string {
	var hs []string
	var next func(key int, a []string) []string
	if code&16 == 16 {
		next = putleft
	} else {
		next = putright
	}
	if code&1 == 1 {
		hs = next(1, hs)
	}
	if code&2 == 2 {
		hs = next(2, hs)
	}
	if code&4 == 4 {
		hs = next(4, hs)
	}
	if code&8 == 8 {
		hs = next(8, hs)
	}
	return hs
}
