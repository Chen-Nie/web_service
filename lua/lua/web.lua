local _M = {}


function _M.webservice(id)

    mysql = require "luasql.mysql"

    local env = mysql.mysql()
    local con = env:connect('web','chen','12345678')

    cursor,errorString = con:execute([[select user_id,username,gender,email from user]])
    row = cursor:fetch ({}, "a")
	
    while row do
        if (row.user_id == id) then
            ngx.say(row.user_id)
            ngx.say(row.username)
            ngx.say(row.gender)
            ngx.say(row.email)
	end
	row = cursor:fetch (row, "a")
    end
end

return _M 
