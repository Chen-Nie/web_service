package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
    con, err := sql.Open("mysql", "chen:12345678@/web")
    if err != nil {panic(err)}
    var user_id int
    var username string
    var gender string
    var email string
    userid, err := fmt.Fprintf(w, "%s", r.URL.Path[1:])
    id := con.QueryRow("SELECT * FROM user WHERE user_id=?", userid)
    switch err := id.Scan(&user_id, &username, &gender, &email); err {
	case sql.ErrNoRows:
  	    fmt.Println("No rows were returned!")
	case nil:
  	    fmt.Println(username, gender, email)
	default:
  	    panic(err)
    }
	    //fmt.Fprintf(w, "%s!", r.URL.Path[1:])
    defer con.Close();
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
