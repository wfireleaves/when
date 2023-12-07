package nl_test

import (
	"testing"
	"time"

	"github.com/wfireleaves/when/rules/nl"

	"github.com/wfireleaves/when"
	"github.com/stretchr/testify/require"
)

var null = time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC)

type Fixture struct {
	Text   string
	Index  int
	Phrase string
	Diff   time.Duration
}

func ApplyFixtures(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.NotNil(t, res, "[%s] res #%d", name, i)
		require.Equal(t, f.Index, res.Index, "[%s] index #%d", name, i)
		require.Equal(t, f.Phrase, res.Text, "[%s] text #%d", name, i)
		require.Equal(t, f.Diff, res.Time.Sub(null), "[%s] diff #%d", name, i)
	}
}

func ApplyFixturesNil(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.Nil(t, res, "[%s] res #%d", name, i)
	}
}

func ApplyFixturesErr(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		_, err := w.Parse(f.Text, null)
		require.NotNil(t, err, "[%s] err #%d", name, i)
		require.Equal(t, f.Phrase, err.Error(), "[%s] err text #%d", name, i)
	}
}

func TestAll(t *testing.T) {
	w := when.New(nil)
	w.Add(nl.All...)

	// complex cases
	fixt := []Fixture{
		{"vorige week zondag om 10:00", 0, "vorige week zondag om 10:00", ((-3 * 24) + 10) * time.Hour},
		{"vanavond om 23:10", 0, "vanavond om 23:10", (23 * time.Hour) + (10 * time.Minute)},
		{"op vrijdagmiddag", 3, "vrijdagmiddag", ((2 * 24) + 15) * time.Hour},
		{"komende dinsdag om 14:00", 0, "komende dinsdag om 14:00", ((6 * 24) + 14) * time.Hour},
		{"komende dinsdag 2 uur 's middags", 0, "komende dinsdag 2 uur 's middags", ((6 * 24) + 14) * time.Hour},
		{"komende woensdag om 14:25", 0, "komende woensdag om 14:25", (((7 * 24) + 14) * time.Hour) + (25 * time.Minute)},
		{"om 11 uur afgelopen dinsdag", 3, "11 uur afgelopen dinsdag", -13 * time.Hour},
		{"volgende week dinsdag om 18:15", 0, "volgende week dinsdag om 18:15", (((6 * 24) + 18) * time.Hour) + (15 * time.Minute)},
		{"volgende week vrijdag", 0, "volgende week vrijdag", (9 * 24) * time.Hour},
	}

	ApplyFixtures(t, "nl.All...", w, fixt)
}
