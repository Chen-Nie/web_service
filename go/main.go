package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "database/sql"
    "../src"
    "encoding/json"
)

sql := src.SQLConnection()
redis := src.redisConnection()

func handler(w http.ResponseWriter, r *http.Request) {
    var user_id int
    var username string
    var gender string
    var email string
    type User struct {
    	Username string
    	Gender string
    	Email string
    }
    
    userids, ok := r.URL.Query()["gouserid"]
    if !ok || len(userids[0]) < 1 {
        fmt.Println("No input")
        return
    }
    var userid string
    userid = userids[0]
    input, err := strconv.Atoi(userid)
    redis.Do("FLUSHALL")
    
    id := sql.QueryRow("SELECT * FROM user WHERE user_id=?", input)
    switch err := id.Scan(&user_id, &username, &gender, &email); err {
	case sql.ErrNoRows:
  	    fmt.Println("No rows were returned!")
	case nil:
	    if_exist, err := redis.Int(redis.Do("EXISTS", userid))
	    if err != nil {fmt.Println("redis error")}
	    if if_exist == 1{
	        fmt.Println("Found the key in redis.")
	    	output, err := redis.String(redis.Do("GET", userid))
	    	if err != nil {log.Fatal(err)}
	    	var o User
	    	err := json.Unmarshal(output, $o)
	    	fmt.Println("Username: " + o.Username + ", Gender: " + o.Gender + ", Email: " + o.Email)
	    } else {
	        fmt.Println("Didn't find the key in redis. Will insert key.")
	        output := User {
	        	Username: username,
	        	Gender: gender,
	        	Email: email,
	        }
	        o, err:= json.Marshal(output)
	        if err != nil {
		    fmt.Println("error:", err)
	        }
	        fmt.Println("Username: " + username + ", Gender: " + gender + ", Email: " + email)
  	        redis.Do("SET", userid, o)
  	    }
	default:
  	    panic(err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

