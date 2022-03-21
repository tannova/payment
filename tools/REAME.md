# Build Docker CI
Following code will build images of registry.gitlab.com/mcuc/monorepo/ci

```bash
# ci
docker login registry.gitlab.com
docker build -t registry.gitlab.com/mcuc/monorepo/ci -f Dockerfile-ci .
docker push registry.gitlab.com/mcuc/monorepo/ci
```

# Build Docker Custon 19.03.12
Following code will build images of registry.gitlab.com/mcuc/monorepo/docker:19.03.12

```bash
# ci
docker login registry.gitlab.com
docker build -t registry.gitlab.com/mcuc/monorepo/docker:19.03.12 -f Docker-custom-19.03.12 .
docker push registry.gitlab.com/mcuc/monorepo/docker:19.03.12
```

