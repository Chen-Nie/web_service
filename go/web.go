package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gomodule/redigo/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {
    con, err := sql.Open("mysql", "chen:12345678@/web")
    if err != nil {panic(err)}
    conn, erro := redis.Dial("tcp", "localhost:6379")
    if erro != nil {log.Fatal(erro)}
    
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
	    if_exist, err := redis.Int(conn.Do("EXISTS", userid))
	    if err != nil {fmt.Println("redis error")}
	    if if_exist == 1{
	        fmt.Println("Found the key in redis.")
	    	output, err := redis.String(conn.Do("GET", userid))
	    	if err != nil {log.Fatal(err)}
	    	fmt.Println(output)
	    } else {
	        fmt.Println("Didn't find the key in redis. Will insert key.")
	        output := "Username: " + username + ", Gender: " + gender + ", Email: " + email
	        fmt.Println(output)
  	        conn.Do("SET", userid, output)
  	    }
	default:
  	    panic(err)
    }
    defer con.Close();
    defer conn.Close()
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

