go-reverse-proxy
================

Let's say you have an existing service that you want to peel off some endpoints to a different service, without changing the client.

That's what this is for.

### Example: Rails App, Add Go Service

So we've got a Rails app on port 3000 and a Go service we want to introduce into the mix on port 2112.

Spin up go-reverse-proxy on port 3000, change the Rails app to port 3001, and have a config.json which looks like:

```json
{
  "Routes":
    [
      {
        "Verb": "get",
        "Source": "/rush_endpoint",
        "Destination": "http://localhost:2112"
      },
      {
        "Verb": "get",
        "Source": "/",
        "Destination": "http://localhost:3001"
      }
    ]
}
```

And then in your Procfile:

```
proxy: PORT=3000 go-reverse-proxy
web: bundle exec rails server -p 3001
rush: PORT=2112 rush-server
```

Now when a request comes in for ```http://localhost:3000/rush_endpoint``` it will be seemlessly proxied to ```http://localhost:2112/rush_endpoint```, all other requests will be proxied to ```http://localhost:3001```

### Proof of Concept

Just something [@danott](http://www.github.com/danott) and I were thinking about.

### Acknowledgements

Uses [Martini](https://github.com/codegangsta/martini) to make a super fast reverse proxy, thanks Jeremy!
