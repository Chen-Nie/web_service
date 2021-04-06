package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
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
    userids, ok := r.URL.Query()["gouserid"]
    if !ok || len(userids[0]) < 1 {
        fmt.Println("No input")
        return
    }
    var userid string
    userid = userids[0]
    input, err := strconv.Atoi(userid)
    
    id := con.QueryRow("SELECT * FROM user WHERE user_id=?", input)
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

