package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	cipher := ""
	for _, i := range plain {
		switch {
		case i > 96 && i < 123:
			cipher += string(rune(((shiftKey + int(i) - 97) % 26) + 97))
		case i > 64 && i < 91:
			cipher += string(rune(((shiftKey + int(i) - 65) % 26) + 65))
		default:
			cipher += string(rune(int(i)))
		}
	}
	return cipher
}
