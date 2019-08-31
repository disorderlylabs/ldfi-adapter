pod=`kubectl get pods --all-namespaces|cut -d " " -f 4|grep mygrpc`

kubectl -n istio-system logs -f $pod





