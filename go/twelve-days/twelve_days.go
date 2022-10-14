package twelve

import (
	"fmt"
	"strings"
)

type ShoppingList struct {
	nth  string
	gift string
}

func Verse(i int) string {
	nthDay := map[int]ShoppingList{
		1:  {nth: "first", gift: "a Partridge in a Pear Tree"},
		2:  {nth: "second", gift: "two Turtle Doves"},
		3:  {nth: "third", gift: "three French Hens"},
		4:  {nth: "fourth", gift: "four Calling Birds"},
		5:  {nth: "fifth", gift: "five Gold Rings"},
		6:  {nth: "sixth", gift: "six Geese-a-Laying"},
		7:  {nth: "seventh", gift: "seven Swans-a-Swimming"},
		8:  {nth: "eighth", gift: "eight Maids-a-Milking"},
		9:  {nth: "ninth", gift: "nine Ladies Dancing"},
		10: {nth: "tenth", gift: "ten Lords-a-Leaping"},
		11: {nth: "eleventh", gift: "eleven Pipers Piping"},
		12: {nth: "twelfth", gift: "twelve Drummers Drumming"}}

	nthVerse := fmt.Sprintf("On the %v day of Christmas my true love gave to me: %v", nthDay[i].nth, nthDay[i].gift)
	for n := i - 1; n > 0; n-- {
		if n == 1 {
			nthVerse += fmt.Sprintf(", and %v", nthDay[n].gift)
		} else {
			nthVerse += fmt.Sprintf(", %v", nthDay[n].gift)
		}
	}
	return nthVerse + "."
}

func Song() string {
	lyrics := []string{}
	for i := 1; i < 13; i++ {
		lyrics = append(lyrics, Verse(i))
	}
	return strings.Join(lyrics, "\n")
}
