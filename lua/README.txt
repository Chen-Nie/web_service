export PATH=/usr/local/openresty/nginx/sbin:$PATH
lsof -ti tcp:8080 | xargs kill (kill all process in port 8080)
nginx -p $PWD/
curl 'http://127.0.0.1:8080/?userid=1'
