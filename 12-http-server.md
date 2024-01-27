# Create an HTTP server to server "wasm micro services"

## Load testing

Do some load tests, first:

```bash
hey -n 100 -c 1 -m POST \
-d 'John Doe' \
"http://localhost:8081" 
```
> note the number of requests per second

Then, do a second test:

```bash
hey -n 100 -c 10 -m POST \
-d 'John Doe' \
"http://localhost:8081" 
```

- What happens when you increase the number of concurrent requests?
- Why?

## Fix the issue

- Fix the issue
- Try again the load tests
- Note the number of requests per second


