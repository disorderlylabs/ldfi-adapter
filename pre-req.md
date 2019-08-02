Pre-Req :

Install minikube
```shell script
brew install minikube

It is highly recommended to give mikube 4 cpus and atleast 6gb ram ( 8 is better)
Start minikube 
minikube start
```
For working with minikube run the command first, in each of the terminal instances :
# remove the $ if you are running fish
```shell script
eval $(minikube docker-env)
```

If not already done, then build the adapter image :

```shell script
docker build -t hub.tess.io/ldfi/ldfi-adapter:latest -f Dockerfile .
```

Now, you can follow the bookinfo readme and adapter install readme or just do the following :

```shell script
cd bookinfo
chmod +x bookinfo_deploy.sh
bookinfo_deploy.sh
cd ..
chmod +x adapter.sh
adapter.sh
```


Wait for a few minutes to see that all pods are running 

```shell script

$ kubectl get pods --all-namespaces
NAMESPACE      NAME                                      READY   STATUS      RESTARTS   AGE
default        details-v1-6f4c4dfb85-6pcnk               2/2     Running     0          10m
default        productpage-v1-bb79f5cc5-88cg2            2/2     Running     0          10m
default        ratings-v1-69757fc969-8nslj               2/2     Running     0          10m
default        reviews-v1-7db8546f48-jtgj9               2/2     Running     0          10m
default        reviews-v2-576c4977c4-tp66l               2/2     Running     0          10m
default        reviews-v3-778cb5596-f82q4                2/2     Running     0          10m
istio-system   grafana-6575997f54-87lk7                  1/1     Running     0          16m
istio-system   istio-citadel-894d98c85-vg97c             1/1     Running     0          16m
istio-system   istio-cleanup-secrets-1.2.2-pmtjd         0/1     Completed   0          16m
istio-system   istio-egressgateway-9b7866bf5-6lp5r       1/1     Running     0          16m
istio-system   istio-galley-5b984f89b-l8rjm              1/1     Running     0          16m
istio-system   istio-grafana-post-install-1.2.2-z85tq    0/1     Completed   0          16m
istio-system   istio-ingressgateway-75ddf64567-txwjt     1/1     Running     0          16m
istio-system   istio-pilot-5d77c559d4-g276p              2/2     Running     0          16m
istio-system   istio-policy-86478df5d4-g58nr             2/2     Running     3          16m
istio-system   istio-security-post-install-1.2.2-bkwgh   0/1     Completed   0          16m
istio-system   istio-sidecar-injector-7b98dd6bcc-gjpsx   1/1     Running     0          16m
istio-system   istio-telemetry-786747687f-98f7g          2/2     Running     3          16m
istio-system   istio-tracing-555cf644d-lg49d             1/1     Running     0          16m
istio-system   kiali-6cd6f9dfb5-79wbk                    1/1     Running     0          16m
istio-system   mygrpcadapter-5c4cf8cd98-6ncf8            1/1     Running     0          6m51s
istio-system   prometheus-7d7b9f7844-rbnjm               1/1     Running     0          16m
kube-system    coredns-5c98db65d4-bzwxt                  1/1     Running     3          38m
kube-system    coredns-5c98db65d4-wh26l                  1/1     Running     3          38m
kube-system    etcd-minikube                             1/1     Running     1          37m
kube-system    kube-addon-manager-minikube               1/1     Running     1          37m
kube-system    kube-apiserver-minikube                   1/1     Running     1          37m
kube-system    kube-controller-manager-minikube          1/1     Running     1          36m
kube-system    kube-proxy-rlzb7                          1/1     Running     1          38m
kube-system    kube-scheduler-minikube                   1/1     Running     1          36m
kube-system    storage-provisioner                       1/1     Running     2          38m
```