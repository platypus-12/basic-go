package main

import (
	"fmt"
	"os"
	"sort"
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

var tracks = []*Track{
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

func printTracks(tracks []*Track) {
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
	tracks     []*Track
	byHeadings []func(*Track, *Track) bool
}

func (ch *ByChosenHeading) Len() int {
	return len(ch.tracks)
}

//最新に選ばれたsortロジックからsortする
func (ch *ByChosenHeading) Less(i, j int) bool {
	for x := len(ch.byHeadings) - 1; x >= 0; x-- {
		if ch.byHeadings[x](ch.tracks[i], ch.tracks[j]) == true {
			return true
		}
	}
	return false
}

func (ch *ByChosenHeading) Swap(i, j int) {
	ch.tracks[i], ch.tracks[j] = ch.tracks[j], ch.tracks[i]
}

//昔に選んだsortロジックも保存したい意
func (ch *ByChosenHeading) AppendHeading(byHeading func(*Track, *Track) bool) {
	ch.byHeadings = append(ch.byHeadings, byHeading)
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

func main() {
	byChosenHeading := &ByChosenHeading{tracks, nil}
	byChosenHeading.AppendHeading(ByTitle)
	sort.Sort(byChosenHeading)
	fmt.Printf("\n\n【ByTitle】\n")
	printTracks(byChosenHeading.tracks)

	byChosenHeading.AppendHeading(ByYear)
	sort.Sort(byChosenHeading)
	fmt.Printf("\n\n【ByYear, ByTitle】\n")
	printTracks(byChosenHeading.tracks)

	fmt.Printf("\n@@@@@@@@@@@@@@@@@@@@@@\n\n")

	byChosenHeading1 := &ByChosenHeading{tracks, nil}
	byChosenHeading1.AppendHeading(ByTitle)
	byChosenHeading1.AppendHeading(ByYear)

	sort.Stable(byChosenHeading1)
	fmt.Printf("\n\n【ByTitle】\n")
	printTracks(byChosenHeading1.tracks)

	sort.Stable(byChosenHeading1)
	fmt.Printf("\n\n【ByYear, ByTitle】\n")
	printTracks(byChosenHeading1.tracks)
}
