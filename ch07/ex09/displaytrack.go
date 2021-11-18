package main

import (
	"basic-go/ch07/ex09/sorting"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
)

var trackPage = template.Must(template.New("trackPage").Parse(`
<html>
<head>
<title>track</title>
</head>
<body>
 <table border="1">
    <tr>
      <th><a href="/?heading=Title">Title</a></th>
      <th><a href="/?heading=Artist">Artist</a></th>
	  <th><a href="/?heading=Album">Album</a></th>
	  <th><a href="/?heading=Year">Year</a></th>
	  <th><a href="/?heading=Length">Length</a></th>
    </tr>
	{{range .}}
    <tr>
      <td>{{.Title}}</td>
	  <td>{{.Artist}}</td>
	  <td>{{.Album}}</td>
	  <td>{{.Year}}</td>
	  <td>{{.Length}}</td>
    </tr>
    {{end}}
  </table>
</body>
</html>
</html>
`))

func printTracks(w io.Writer, tracks []*sorting.Track) {
	if err := trackPage.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

func displayTrackPage(w http.ResponseWriter, r *http.Request) {
	var byChosenHeading = &sorting.ByChosenHeading{Tracks: sorting.Tracks, ByHeadings: nil}
	heading := r.URL.Query().Get("heading")
	if heading == "Title" {
		byChosenHeading.AppendHeading(sorting.ByTitle)
	} else if heading == "Artist" {
		byChosenHeading.AppendHeading(sorting.ByArtist)
	} else if heading == "Album" {
		byChosenHeading.AppendHeading(sorting.ByAlbum)
	} else if heading == "Year" {
		byChosenHeading.AppendHeading(sorting.ByYear)
	} else if heading == "Length" {
		byChosenHeading.AppendHeading(sorting.ByLength)
	}
	sort.Sort(byChosenHeading)
	printTracks(w, byChosenHeading.Tracks)
}

func main() {
	http.HandleFunc("/", displayTrackPage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
