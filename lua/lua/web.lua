local _M = {}


function _M.webservice(id)

    local mysql = require "luasql.mysql"
    local redis = require "redis"
    local client = redis.connect('127.0.0.1', 6379)

    local env = mysql.mysql()
    local con = env:connect('web','chen','12345678')

    cursor,errorString = con:execute([[select user_id,username,gender,email from user]])
    row = cursor:fetch ({}, "a")
    
    while row do
        if (row.user_id == id) then
            if_exist = client:exists(row.user_id)
            ngx.say(if_exist)
            if (if_exist == true) then
                local output = client:get(id)
                ngx.say("Found the key in redis.")
                ngx.say(output)
            else
                local output = row.username .. " " .. row.gender .. " " .. row.email
                client:set(id, output)
                ngx.say("Didn't find the key in redis. Will insert key.")
                ngx.say(output)
            end
	end
	row = cursor:fetch (row, "a")
    end
end

return _M 

