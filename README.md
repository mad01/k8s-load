# k8s-load

[![Docker Repository on Quay](https://quay.io/repository/mad01/k8s-load/status "Docker Repository on Quay")](https://quay.io/repository/mad01/k8s-load)

small golang service to generate cpu load. and a container with the load generator hey. the service is configured to autoscale from 1-50 pods based on pod load of `50%` . 

this is done for testing of different horizontal pod autoscaler algoritms. for more load just change the hey `-c` concurrent connection flag, `-n` total number of requests and the replica count of the hey deployment. 


## Build with bazel

This repo is works with the standard go tools. It also works with [Bazel](https://www.bazel.build).

Building with Bazel requires a bazel 5.3.0+ installation and running the following command:

```
bazel build //...
```

After the build is complete the `k8s-load` binary lives under the bazel-bin directory: 

```
bazel-bin/k8s-load
```  


### kubernetes deployment
```bash
kubectl apply -f deployment-hey.yaml
kubectl apply -f deployment-service.yaml
```
