# webproxy

1. Build Go:
$GOOS=linux GARCH=amd64 go build webproxy

2. Build Docker:

$docker build -t webproxy:1.0.0

3. Run Docker:

$docker run -e "SERVER_HOST=localhost" -e "SERVER_PORT=3000" -p 5000:5000 webproxy:1.0.0

4. Test:

$curl http://localhost:5000/proxy