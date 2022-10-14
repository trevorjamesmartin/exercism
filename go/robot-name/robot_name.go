package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	major    string
	minor    string
	revision int
}

func randomName() (string, string, int) {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []byte("01234567891234567891234567")
	rand.Shuffle(len(letters), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		letters[i], letters[j] = letters[j], letters[i]
	})
	major := string(letters[rand.Int()%25])
	minor := string(letters[rand.Int()%25])
	revision := rand.Int() % 999
	return major, minor, revision
}

type robotfarm struct {
	limit int
	index map[string]*Robot
	count int
}

var robots = &robotfarm{limit: 26 * 26 * 10 * 10 * 10, index: map[string]*Robot{}}

// 2 uppercase letters, followed by 3 digits
func (r *Robot) Name() (string, error) {
	var nameError error
	var major, minor, stringName string
	var revision int
	if r.revision > 0 {
		return fmt.Sprintf("%v%v%03d", r.major, r.minor, r.revision), nil
	}
	if robots.count >= robots.limit {
		return "", fmt.Errorf("limit reached %v", robots.count)
	}
	for r.revision == 0 {
		major, minor, revision = randomName()
		stringName = fmt.Sprintf("%v%v%03d", major, minor, revision)
		_, found := robots.index[stringName]
		if !found {
			r.major = major
			r.minor = minor
			r.revision = revision
			robots.index[stringName] = r
			robots.count++
		}
	}
	return stringName, nameError
}

func (r *Robot) Reset() {
	r.major = ""
	r.minor = ""
	r.revision = 0
}
