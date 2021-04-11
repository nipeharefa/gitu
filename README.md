Gitu
====

Simple Static File server with config.
Inspired from Vercel.


### Getting Started

TODO


### Example

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