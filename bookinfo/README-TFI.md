get the adapter running

kubectl apply -f ../cluster_service.yaml 
kubectl apply -f ../testdata/mygrpcadapter.yaml 
kubectl apply -f ../testdata/attributes.yaml
kubectl apply -f ../testdata/template.yaml 
kubectl apply -f ../testdata/tfi.yaml 


N.B.: this works only nondeterministically now, since the adapter is not always invoked :( :(

try it a couple of times.

curl -H 'x-ebay-ldfi:after=details|details' http://localhost/productpage





