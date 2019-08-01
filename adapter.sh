#!/usr/bin/env bash

kubectl apply -f cluster_service.yaml
kubectl apply -f testdata/attributes.yaml
kubectl apply -f testdata/template.yaml
kubectl apply -f testdata/mygrpcadapter.yaml
kubectl apply -f testdata/sample_operator_cfg.yaml