//Exercise 7.9:
//Use the html/template package (§4.6) to
//replace printTracks with a function that displays the tracks as
//an HTML table.
//
//Use the solution to the previous exercise to arrange that each click on
//a column head makes an HTTP request to sort the table.

package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length int
}

type Table struct {
	Tracks     []*Track
	SortColumn int
}

func (t *Table) Len() int {
	return len(t.Tracks)
}

func (t *Table) Less(i, j int) bool {
	switch t.SortColumn {
	case 0:
		return t.Tracks[i].Title < t.Tracks[j].Title
	case 1:
		return t.Tracks[i].Artist < t.Tracks[j].Artist
	case 2:
		return t.Tracks[i].Album < t.Tracks[j].Album
	case 3:
		return t.Tracks[i].Year < t.Tracks[j].Year
	case 4:
		return t.Tracks[i].Length < t.Tracks[j].Length
	default:
		return false
	}
}

func (t *Table) Swap(i, j int) {
	t.Tracks[i], t.Tracks[j] = t.Tracks[j], t.Tracks[i]
}

func main() {
	tracks := []*Track{
		{Title: "Track 1", Artist: "Artist 1", Album: "Album 1", Year: 2020, Length: 180},
		{Title: "Track 2", Artist: "Artist 2", Album: "Album 2", Year: 2019, Length: 200},
		{Title: "Track 3", Artist: "Artist 3", Album: "Album 1", Year: 2021, Length: 150},
	}

	table := &Table{Tracks: tracks}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			col, err := strconv.Atoi(r.FormValue("col"))
			if err != nil {
				log.Println("Invalid column value:", err)
				http.Error(w, "Invalid column value", http.StatusBadRequest)
				return
			}
			table.SortColumn = col
			sort.Sort(table)
		}

		tmpl := template.Must(template.New("tracks").Parse(`
			<html>
			<head>
				<title>Track Listing</title>
			</head>
			<body>
				<table>
					<tr>
						<th><a href="/?col=0">Title</a></th>
						<th><a href="/?col=1">Artist</a></th>
						<th><a href="/?col=2">Album</a></th>
						<th><a href="/?col=3">Year</a></th>
						<th><a href="/?col=4">Length</a></th>
					</tr>
					{{range .Tracks}}
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
		`))

		err := tmpl.Execute(w, table)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
