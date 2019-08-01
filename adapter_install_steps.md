We need to create a service in the cluster. This service needs to expose the port for grpc communication.
This port is the same as the one we supplied in the docker file for the adapter.
We will use the same port in the operator config.

This is followed by applying the attributes and the template spec. 
Finally, since we are running in the cluster, we will supply the address in the sample_operator_cfg.yml as 

```
mygrpcadapterservice:44225
```

For working with minikube run the command first, in each of the terminal instances :
```shell script
eval $(minikube docker-env)
```

The commands are in adapter.sh. Summarized below

```shell script
kubectl apply -f cluster_service.yaml
kubectl apply -f testdata/attributes.yaml
kubectl apply -f testdata/template.yaml
kubectl apply -f testdata/mygrpcadapter.yaml
kubectl apply -f testdata/sample_operator_cfg.yaml
```

To comeback to your local environment run the following :
```shell script
eval $(docker-machine env -u)
```

To confirm if the sidecar is running the READY column should be 2/2, see an example below :

```shell script
$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
details-v1-6f4c4dfb85-ts4gc      2/2     Running   0          30h
productpage-v1-bb79f5cc5-xtsnd   2/2     Running   0          30h
ratings-v1-69757fc969-gdh2b      2/2     Running   0          30h
reviews-v1-7db8546f48-v5jhl      2/2     Running   0          30h
reviews-v2-576c4977c4-frz8q      2/2     Running   0          30h
reviews-v3-778cb5596-kw2sz       2/2     Running   0          30h
```

Next, let's try and hit the bookinfo with a bad header
```shell script

⋊> araina1 at LM-SJC-11016345 in ~/work/myootadapter (master|●4✚5…389) since Thu Aug 01 00:06:58
$ curl -vk -H "x-ebay-ldfi:bar" http://192.168.99.100:31380/productpage
*   Trying 192.168.99.100...
* TCP_NODELAY set
* Connected to 192.168.99.100 (192.168.99.100) port 31380 (#0)
> GET /productpage HTTP/1.1
> Host: 192.168.99.100:31380
> User-Agent: curl/7.54.0
> Accept: */*
> x-ebay-ldfi:bar
>
< HTTP/1.1 403 Forbidden
< content-length: 57
< content-type: text/plain
< date: Thu, 01 Aug 2019 07:07:10 GMT
< server: istio-envoy
<
* Connection #0 to host 192.168.99.100 left intact
PERMISSION_DENIED:h1.handler.istio-system:Unauthorized...⏎
```

Now with the right header 

```shell script
⋊> araina1 at LM-SJC-11016345 in ~/work/myootadapter (master|●4✚5…389) since Thu Aug 01 00:06:41
$ curl -vk -H "x-ebay-ldfi:abc" http://192.168.99.100:31380/productpage
*   Trying 192.168.99.100...
* TCP_NODELAY set
* Connected to 192.168.99.100 (192.168.99.100) port 31380 (#0)
> GET /productpage HTTP/1.1
> Host: 192.168.99.100:31380
> User-Agent: curl/7.54.0
> Accept: */*
> x-ebay-ldfi:abc
>
< HTTP/1.1 200 OK
< content-type: text/html; charset=utf-8
< content-length: 3769
< server: istio-envoy
< date: Thu, 01 Aug 2019 07:06:57 GMT
< x-envoy-upstream-service-time: 35
<
<!DOCTYPE html>
<html>
  <head>
    <title>Simple Bookstore App</title>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">

.... some more html content
* Connection #0 to host 192.168.99.100 left intact
```