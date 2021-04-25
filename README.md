Gitu
====

Simple Static File server with config, inspired from Vercel. With gitu you can configure cache, add custome header and redirect like rules in another platform like firebase, vercel, and other.
## Getting Started

## Using Docker

```
FROM alpine:3.10

WORKDIR /app

# USER apps
RUN mkdir static

COPY --from=docker.io/nipeharefa/gitu:0.0.6 /app/main main
COPY --from=builder /app/build ./static/
COPY now.json now.json

CMD [ "/app/main" ]
```


### Example

Create static directory, and create static content
```
mkdir static && cd static
touch index.html
mkdir css jss
touch css/index.css
```

Create new file configuration
```
touch now.json
```

and add simple config to file
```
{
  "routes": [
    {
      "src": "/",
      "headers": {
        "cache-control": "no-cache"
      }
    },
    {
      "src": "/(js|css)/",
      "headers": {
        "cache-control": "public, s-maxage=15552000, max-age=15552000, must-revalidate"
      }
    },
    {
      "src": "/(.*)",
      "rewrite": "/"
    }
  ]
}
```
and run gitu
```
gitu
```
