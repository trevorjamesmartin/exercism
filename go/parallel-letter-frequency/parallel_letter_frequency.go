package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Combinator reads from a [ch]annel, adding the result to the [m]ap supplied by the first parameter.
func Combinator(m map[rune]int, expected int, ch <-chan FreqMap, wg *sync.WaitGroup) {
	defer wg.Done()
	var total int
	for freqs := range ch {
		total++
		// combine with map
		for r := range freqs {
			m[r] += freqs[r]
		}
		if total == expected {
			// when all maps have been received
			break
		}
	}
}

// ReadFrequency writes the result of Frequency() to the channel
func ReadFrequency(s string, ch chan<- FreqMap, wg *sync.WaitGroup) FreqMap {
	defer wg.Done()
	m := Frequency(s)
	ch <- m
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(l) + 1) // our list of strings, plus our combinator
	m := FreqMap{}
	channel := make(chan FreqMap)
	go Combinator(m, len(l), channel, &wg) // 1 writer
	for _, s := range l {
		go ReadFrequency(s, channel, &wg) // many readers/reporters
	}
	wg.Wait()
	return m
}
