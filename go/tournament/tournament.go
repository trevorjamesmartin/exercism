package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type ScoreCard struct {
	MP, W, D, L, P int
	Team           string
}

type By func(card1, card2 *ScoreCard) bool

type CardSorter struct {
	scorecards []ScoreCard
	by         By
}

type ScoreBoard map[string]*ScoreCard

func (sb ScoreBoard) Keys() []string {
	keys := make([]string, 0, len(sb))
	for k := range sb {
		keys = append(keys, k)
	}
	return keys
}

func (sb ScoreBoard) Sorted() []ScoreCard {
	cards := make([]ScoreCard, 0, len(sb))
	for _, key := range sb.Keys() {
		cards = append(cards, *sb[key])
	}
	points_desc := func(a, b *ScoreCard) bool {
		if a.P == b.P {
			// ascend
			return a.Team < b.Team
		}
		// descend
		return a.P > b.P
	}
	By(points_desc).Sort(cards)
	return cards
}

func (a ScoreBoard) ReadScore(line string) error {
	v := strings.Split(line, ";")
	if len(v) != 3 {
		return fmt.Errorf("insufficient information: [%v]", line)
	}
	team1, team2, result := v[0], v[1], strings.ToLower(v[2])
	if result != "win" && result != "loss" && result != "draw" {
		return fmt.Errorf("unexpected result: %v", result)
	}
	// create ScoreCards when necessary
	t1, found1 := a[team1]
	t2, found2 := a[team2]
	if !found1 {
		a[team1] = &ScoreCard{Team: team1}
		t1 = a[team1]
	}
	if !found2 {
		a[team2] = &ScoreCard{Team: team2}
		t2 = a[team2]
	}
	// update scores
	t1.MP++
	t2.MP++
	switch result {
	case "win":
		t1.W++    // win
		t2.L++    // lose
		t1.P += 3 // (team1 wins 3 points)
	case "loss":
		t2.W++    // win
		t1.L++    // lose
		t2.P += 3 // (team2 wins 3 points)
	case "draw":
		t1.D++
		t1.P++
		t2.D++
		t2.P++
	}
	return nil
}

func (by By) Sort(cards []ScoreCard) {
	sort.Sort(&CardSorter{scorecards: cards, by: by})
}

func (a *CardSorter) Len() int {
	return len(a.scorecards)
}

func (a *CardSorter) Swap(b, c int) {
	a.scorecards[b], a.scorecards[c] = a.scorecards[c], a.scorecards[b]
}

func (a *CardSorter) Less(b, c int) bool {
	return a.by(&a.scorecards[b], &a.scorecards[c])
}

func Tally(reader io.Reader, writer io.Writer) error {
	var scoreboard = ScoreBoard{}
	input := bufio.NewScanner(reader)
	var err error
	for input.Scan() {
		line := input.Text()
		// (ignore comments)
		if len(line) > 0 && []rune(line)[0] != '#' {
			err = scoreboard.ReadScore(line)
		}
		if err != nil {
			return err
		}
	}
	cards := scoreboard.Sorted()
	fmt.Fprintf(writer, "%v\n", "Team                           | MP |  W |  D |  L |  P")
	for _, t := range cards {
		fmt.Fprintf(writer, "%v\n", fmt.Sprintf("%v%v| %2d | %2d | %2d | %2d | %2d", t.Team, strings.Repeat(" ", 31-len(t.Team)), t.MP, t.W, t.D, t.L, t.P))
	}
	return err
}
