details=`kubectl get pods | grep details | cut -d " " -f 1`
product=`kubectl get pods | grep productpage | cut -d " " -f 1`

kubectl cp ~/go/src/istio.io/istio/samples/bookinfo/src/productpage/productpage.py $product:/opt/microservices/
#kubectl cp ~/go/src/istio.io/istio/samples/bookinfo/src/details/details.rb $details:/opt/microservices/
