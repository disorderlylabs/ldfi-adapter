eval $(minikube docker-env)
docker build -t hub.tess.io/ldfi/ldfi-adapter:latest -f Dockerfile .


pod=`kubectl get pods --all-namespaces|cut -d " " -f 4|grep mygrpc`

kubectl -n istio-system delete pod $pod


pod=`kubectl get pods --all-namespaces|cut -d " " -f 4|grep mygrpc`

kubectl -n istio-system logs -f $pod 



