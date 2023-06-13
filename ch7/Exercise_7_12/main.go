//“Exercise 7.12:
//Change the handler for /list to print its output as an HTML
//table, not text. You may find the html/template package (§4.6) useful.”

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.New("list").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Inventory List</title>
		</head>
		<body>
			<table>
				<tr>
					<th>Item</th>
					<th>Price</th>
				</tr>
				{{ range $item, $price := . }}
					<tr>
						<td>{{ $item }}</td>
						<td>{{ $price }}</td>
					</tr>
				{{ end }}
			</table>
		</body>
		</html>
	`))

	err := tmpl.Execute(w, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "item and price parameters are required")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "invalid price value")
		return
	}

	db[item] = dollars(price)
	fmt.Fprintln(w, "Item created successfully")
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "item and price parameters are required")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "invalid price value")
		return
	}

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	db[item] = dollars(price)
	fmt.Fprintln(w, "Item updated successfully")
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "item parameter is required")
		return
	}

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintln(w, "Item deleted successfully")
}
