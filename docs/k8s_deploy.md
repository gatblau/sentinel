# K8S deployment

In order to run Sentinel in Kubernetes, it is necessary to create a cluster role that contains the rules granting it access to the API objects that need watching.

The script [sl_rbac.yaml](../scripts/kube/sl_rbac.yaml) creates a **resource_watcher** cluster wide role, a **sentinel** service account and assigns the privileges in the role to the service account using a role binding.

To apply the above configuration:

```bash
# second, create a role, account and binding
kubectl create -f ./deploy/sl_rbac.yaml
```

Now that the service account has the required privileges, you might want to overlay a different [config.toml](../config.toml) configuration file. In this case, a config map like the one shown [here](../scripts/kube/sl_config_map.yaml) can be used.

If you want to create a config map from a file do the following:

```bash
# create config map from config.toml file
kubectl create configmap sentinel --from-file=config.toml

# get the content in yaml format
kubectl get configmaps sentinel -o yaml
```

or otherwise, load the map [here](../scripts/kube/sl_config_map.yaml):

```bash
# load config map from yaml file
kubectl create -f ./deploy/sl_config_map.yaml
```

Alternatively, if no config map is set, Sentinel will use the default config.toml file in the image, or the override values based on the equivalent environment variables passed in as required.

With the above ready, can now deploy the Sentinel pod:

```bash
# finally deploy the Sentinel pod
kubectl create -f ./deploy/sl_pod.yaml

# check status
kubectl get pods
```

To see the Sentinel events logged to the stdout:

```bash
kubectl logs sentinel
time="2019-06-08T13:43:24Z" level=info msg="Loading configuration."
time="2019-06-08T13:43:24Z" level=info msg="TRACE has been set as the logger level." platform=kube-01
...
...
```
