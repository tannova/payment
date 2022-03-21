# How to migrate from current helm chart template to new helm chart template?

Step 1: cd to `<app_name>/<env>`

Step 2: delete templates folder

Step 3: Create Symbolic link to templates folder with  command:  
```bash
ln -s ../../_main-chart/templates templates
```
Only work with MacOs/Linux system.
 

Step 4: Migrate values.yaml
    
    - Remove global section and reformat
    - Move env section to root
    - If use vaultsecret or load balancer, add enabled: true
    - Add valutsecret.secretName field if vaultsecret.enabled
    - Rename loadbalancer.sslCert to loadBalancer.sslCertArn
    - Rename images.image to images.name
    - Rename images to image
    - Optional: Add resources section
    Example

Before: 

```yaml
  replicaCount: 1
  namespace: default
  vaultsecret:
    key: user
    path: kv/apps/echo
  image:
    secret: registry-credentials
    repository: registry.gitlab.com
    pullPolicy: Always
    name: mcuc/monorepo/echo
    tag: latest
    env:
      - name: USER
        valueFrom:
          secretKeyRef:
            name: echo
            key: user
  resources: {}
  service:
    type: NodePort
    port: 8080
    targetPort: 8080
  loadBalancer:
    enabled: true
    sslCertArn: arn:aws:acm:ap-southeast-1:732572206546:certificate/faa46be2-0539-4d66-8eb2-85a94b88ef0b
    host: echo.stg.cmgp.inspirelab.io

```

After:

```yaml
replicaCount: 1
namespace: default
vaultsecret:
  enabled: true
  secretName: echo
  path: kv/apps/echo
image:
  secret: registry-credentials
  repository: registry.gitlab.com
  pullPolicy: Always
  name: mcuc/monorepo/echo
  tag: latest
env:
  - name: USER
    valueFrom:
      secretKeyRef:
        name: echo
        key: user
resources: {}
service:
  type: NodePort
  ports:
    - name: "grpc"
      port: 8080
      targetPort: 8080
loadBalancer:
  enabled: true
    sslCertArn: arn:aws:acm:ap-southeast-1:732572206546:certificate/faa46be2-0539-4d66-8eb2-85a94b88ef0b
    host: echo.stg.cmgp.inspirelab.io
```


# How to create new app

Step 1: Create folder <app_name>/\<env>

Step 2: Create Symbolic link to templates folder with  command:  
```bash
ln -s ../../_main-chart/templates templates
```
Only work with MacOs/Linux system.

Step 3:
    Copy `Chart.yaml` & `values.yaml` from `../../main-chart` to  `<app_name>/<env>` folder

Step 4: Edit `Chart.yaml` & `values.yaml`
 


# How to package & publish chart?

```
helm package <app_name></env>

helm s3 push ..
```