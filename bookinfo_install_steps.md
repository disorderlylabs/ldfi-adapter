
Follow link : https://istio.io/docs/setup/kubernetes/install/kubernetes/

```shell script
for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl apply -f $i; done

kubectl apply -f install/kubernetes/istio-demo.yaml

Confirm if things are in order :
kubectl get svc -n istio-system

```

If your cluster is running in an environment that does not support an external load balancer (e.g., minikube), 
**the EXTERNAL-IP of istio-ingressgateway will say <pending>**.

To access the gateway, use the service’s NodePort, or use port-forwarding instead.

Ensure that istio pods running:

```shell script
kubectl get pods -n istio-system
```

Deploying Bookinfo :

```shell script
kubectl label namespace default istio-injection=enabled

kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
```

Check if services are running :

```shell script
kubectl get services

# there should be a cluster-ip 

kubectl get pods --all-namespaces

# status of the pods should be running, this will take a few minutes
```

Check if the app is running :

```shell script
kubectl exec -it $(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}') -c ratings -- curl productpage:9080/productpage

some html will be there as output
```

Next, we need a gateway to be able to access bookinfo

```shell script
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml

Output :
gateway.networking.istio.io/bookinfo-gateway created
virtualservice.networking.istio.io/bookinfo created

# check if the gateway is there
$ kubectl get gateway
NAME               AGE
bookinfo-gateway   44s
```

On minikube the ingress-gateway won't have externalip, we need to take some steps to expose the gateway for local developement

```shell script

-- general k8s with external LB
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
```

```shell script

-- for minkube we will need to use the minikube ip and node ports
export INGRESS_HOST=$(minikube ip)
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')

--export the gateway url
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT

-- your env should look like this ( notice the higher port range, if your ports are 80 and 443 then something isn't right
INGRESS_HOST=192.168.99.100
INGRESS_PORT=31380
SECURE_INGRESS_PORT=31390
GATEWAY_URL=192.168.99.100:31380

```

# cleanup Bookinfo

```shell script
samples/bookinfo/platform/kube/cleanup.sh
```
