# Installing on OpenShift

## Configuring Security

Sentinel requires *watch* and *list* privileges on the resources is watching for changes across the Kubernetes cluster. These privileges are described in a new [cluster role](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#role-and-clusterrole) that can be found [here](cluster_role.yaml).

Before running the template, the above cluster role needs to be created with the privileges required by Sentinel to run.

The role then needs to be bound to the sentinel service account.

To create the cluster role, service account, role binding and namespace, log in as cluster admin and execute the following:

```bash
make oc-setup
```

## Installing using the web console catalogue
 
Import the Sentinel template as follows:

```bash
# import the template in the catalogue
make oc-import-template
```
Once the template is imported in OpenShift, it shows in the catalogue and can be run using the web console.

## Installing using the command line

```bash
# you are log as system admin!
oc login -u system:admin

# create a new project, role, account and bindings
sh ./scripts/openshift/setup.sh

# deploy the app from the file system
oc new-app ./scripts/openshift/sentinel.yml
```

## Cleanup operations

To delete the template:
```bash
make oc-delete-template
```

To remove all Sentinel resources:
```bash
make oc-cleanup
```
