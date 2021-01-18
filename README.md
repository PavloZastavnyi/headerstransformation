# Headers Transformation

Headers Transformation is a middleware plugin for [Traefik](https://github.com/containous/traefik) which transform headers.

## Configuration

### Static

```yaml
pilot:
    token: {pilot token}

api:
  dashboard: true
  insecure: true

experimental:
  plugins:
    headerstransformation:
      moduleName: "github.com/PavloZastavnyi/headerstransformation"
      version: "v0.0.1"

entryPoints:
  http:
    address: ":8000"
    forwardedHeaders:
      insecure: true

providers:
  file:
    filename: rules-headerstransformation.yaml
```

### Dynamic

```yaml
http:
  routers:
    router:
      entryPoints:
      - http
      middlewares:
      - headerstransformation
      service: service-whoami
      rule: Path(`/whoami`)

  services:
    service-whoami:
      loadBalancer:
        servers:
        - url: http://localhost:5000/
        passHostHeader: false
  
  middlewares:
    headerstransformation:
      plugin:
        headerstransformation:
         HeaderName: "X-Traefik-Test"
```

## Docker
Run whoami docker container for demo
```bash
docker run -d --network host containous/whoami -port 5000
```