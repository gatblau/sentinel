# Installing on OpenShift

```bash
# create new project
oc new-project sentinel

# deploy the app
oc new-app https://raw.githubusercontent.com/gatblau/sentinel/dev/install/sentinel.yaml
```