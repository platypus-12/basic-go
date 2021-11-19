package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var itemPage = template.Must(template.New("itemPage").Parse(`
<html>
<head>
<title>track</title>
</head>
<body>
 <table border="1">
    <tr>
      <th>Item</th>
      <th>Price</th>
    </tr>
	{{range $i, $v := .}}
    <tr>
      <td>{{$i}}</td>
	  <td>{{$v}}</td>
    </tr>
    {{end}}
  </table>
</body>
</html>
</html>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemPage.Execute(w, db); err != nil {
		log.Fatal(err)
	}
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item:%q\n", item)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_price := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(_price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s created\n", item, dollars(price))
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, existKey := db[item]; existKey == false {
		fmt.Fprintf(w, "there isn't key named %s", item)
		return
	}
	_price := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(_price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s\n", err.Error())
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s updated\n", item, dollars(price))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, existKey := db[item]; existKey == false {
		fmt.Fprintf(w, "there isn't key named %s", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s deleted", item)
}
