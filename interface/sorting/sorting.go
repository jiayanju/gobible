package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	printTracks(tracks)

	sort.Sort(byArtist(tracks))
	fmt.Println()
	fmt.Println("Sort by Artist")
	fmt.Println()
	printTracks(tracks)

	sort.Sort(sort.Reverse(byArtist(tracks)))

	fmt.Println()
	fmt.Println("Reverse Sort by Artist")
	fmt.Println()
	printTracks(tracks)

	fmt.Println()
	fmt.Println("Custom Sort")
	fmt.Println()
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Year != y.Year {
			return x.Year < y.Year
		}

		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)

	// printTracksHTML(tracks)

	const port = 8088
	fmt.Printf("Start web server at %d\n", port)

	http.HandleFunc("/tracks", func(w http.ResponseWriter, req *http.Request) {
		printTracksHTML(tracks, w)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("error : ", err)
	}
}

//Track track record
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go Home", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m32s")},
	{"Go", "Alica Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m11s")},
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
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func printTracksHTML(tracks []*Track, w io.Writer) {
	const tpl = `
        <html>
            <head>
                <title>Tracks</title>
            </head>
            <body>
                <p>Tracks</p>
                <table border="1">
                    <tr>
                        <th>Title</th>
                        <th>Artist</th>
                        <th>Album</th>
                        <th>Year</th>
                        <th>Length</th>
                    </tr>
                    {{range .}}
                        <tr>
                            <td style="text-align: center">{{.Title}}</td>
                            <td style="text-align: center">{{.Artist}}</td>
                            <td style="text-align: center">{{.Album}}</td>
                            <td style="text-align: center">{{.Year}}</td>
                            <td style="text-align: center">{{.Length}}</td>
                        </tr>
                    {{end}}
                </table>
            </body>
        </html>
    `

	t, _ := template.New("table").Parse(tpl)
	err := t.Execute(w, tracks)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	tracks []*Track
	less   func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.tracks)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.tracks[i], c.tracks[j])
}

func (c customSort) Swap(i, j int) {
	c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i]
}
