# Traefik demo plugin

try to see how plug-in in traefik works

### Configuration


```yaml
# Static configuration

experimental:
  plugins:
    example:
      moduleName: github.com/sipp11/traefik-demo-plugino
      version: v0.1.2
```

Here is an example of a file provider dynamic configuration (given here in YAML), where the interesting part is the `http.middlewares` section:

```yaml
# Dynamic configuration

http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - my-plugin

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000

  middlewares:
    my-plugin:
      plugin:
        example:
          headers:
            Foo: Bar
```
