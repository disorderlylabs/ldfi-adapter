set up bookinfo.

set up destination rules:

kubectl apply -f samples/bookinfo/networking/destination-rule-all.yaml

set up virtual services:

kubectl apply -f samples/bookinfo/networking/virtual-service-all-v1.yaml

set up service-specific fault paths:

kubectl apply -f details_bad.yaml 
kubectl apply -f ratings_bad.yaml 
kubectl apply -f reviews_bad.yaml 

#  kubectl delete -f samples/bookinfo/platform/kube/bookinfo.yaml

#  kubectl apply -f bookinfo_local.yaml

set up propagation: 

kubectl apply -f ../testdata/rlfi.yaml


now request-level fault injection should "just work."  e.g.



curl -H x-ebay-ldfi:fail=ratings http://localhost/productpage
