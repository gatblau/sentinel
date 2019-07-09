# Installing on OpenShift

First of all, create a project to deploy Sentinel in:
```bash
# create new project
oc new-project sentinel --display-name="Sentinel" --description="Hosts the Sentinel application."
```
The Sentinel template can be imported as follows:
```bash
# import the template in the catalogue
oc create -f sentinel.yml -n openshift
```
Once the template is imported in OpenShift, it shows in the catalogue and can be deployed using the web console.

Otherwise, to install it from the command line:

```bash
# deploy the app
oc new-app https://raw.githubusercontent.com/gatblau/sentinel/dev/install/openshift/sentinel.yml
```

