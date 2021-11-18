package sorting

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var Tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Delilah", "As I Am", 2012, length("4m36s")},
	{"Go", "Delilah2", "As I Am", 2012, length("4m36s")},
	{"not Go", "Delilah2", "As I Am", 2012, length("4m36s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type ByChosenHeading struct {
	Tracks     []*Track
	ByHeadings []func(*Track, *Track) bool
}

func (ch *ByChosenHeading) Len() int {
	return len(ch.Tracks)
}

//最新に選ばれたsortロジックからsortする
func (ch *ByChosenHeading) Less(i, j int) bool {
	for x := len(ch.ByHeadings) - 1; x >= 0; x-- {
		if ch.ByHeadings[x](ch.Tracks[i], ch.Tracks[j]) == true {
			return true
		}
	}
	return false
}

func (ch *ByChosenHeading) Swap(i, j int) {
	ch.Tracks[i], ch.Tracks[j] = ch.Tracks[j], ch.Tracks[i]
}

//昔に選んだsortロジックも保存したい意
func (ch *ByChosenHeading) AppendHeading(byHeading func(*Track, *Track) bool) {
	ch.ByHeadings = append(ch.ByHeadings, byHeading)
}

func ByTitle(i, j *Track) bool {
	return i.Title < j.Title
}

func ByArtist(i, j *Track) bool {
	return i.Artist < j.Artist
}

func ByAlbum(i, j *Track) bool {
	return i.Album < j.Album
}

func ByYear(i, j *Track) bool {
	return i.Year < j.Year
}

func ByLength(i, j *Track) bool {
	return i.Length < j.Length
}
