<img src="./../pics/sentinel_small.png" align="right" height="200" width="200"/>

# Getting Started

The easiest way to getting started is to deploy Sentinel from a container image in [minikube](https://github.com/kubernetes/minikube).

Alternatively, you can try run the binary outside of an image and redirect the output to the file system.

## Installing required tools

To build it from code, make sure you have the latest version of [golang](https://golang.org/dl/) installed on your machine.

You will also need to have a [git](https://www.atlassian.com/git/tutorials/install-git) client and [make](https://www.gnu.org/software/make/) installed.

Finally, as Sentinel is meant to observe changes in a Kubernetes cluster, you need to have Kubernetes running and point Sentinel to it. The simplest way to have kubernetes running locally is by using [minikube](https://github.com/kubernetes/minikube). You need to [install it](https://kubernetes.io/docs/tasks/tools/install-minikube/) on your local machine. And to talk to minikube you also need to [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), the Kubernetes command line tool.

## Deploying from a container image

To deploy Sentinel from the container image into K8S [see here](./k8s_deploy.md).

## Trying the binary out in the local machine

If you want to try it out on your local machine [see here](./binary_deploy.md)

## Publishing to a webhook

In order to publish to a web hook, update the Sentinel configuration file to use the webhook publisher. Set the URI of the publisher to the address the web consumer application is listening to.

Start up a web consumer application where the web hook is pointing to, [such as the one here](../consumer/web_consumer.py).

Run the sentinel process.

Deploy an application on minikube and see the changes appearing on the web consumer application terminal.

[*] _The Sentinel icon was made by [Freepik](https://www.freepik.com) from [Flaticon](https://www.flaticon.com) and is licensed by [Creative Commons BY 3.0](http://creativecommons.org/licenses/by/3.0)_