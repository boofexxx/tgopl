package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("%.2f$", d)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	p, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price: %v", price)
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "price to %s has been changed to %v", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "item %s has been deleted", item)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "such item is already exists: %q\n", item)
		return
	}
	p, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price: %v", price)
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "item %s has been created with price %v", item, db[item])
}

type database map[string]dollars

var itemList = template.Must(template.New("itemList").Parse(`
<html>
<body>
<table>
<tr>
	<th>Item</th>
	<th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
	<td>{{$key}}</td>
	<td>{{$value}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
