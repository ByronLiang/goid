#!/bin/bash

kubectl create configmap $1 -n $2 --from-file $3 -o yaml --dry-run | kubectl apply -f -
