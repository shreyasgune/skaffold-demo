kubectl delete deployment nginx-ingress-controller skaffold-demo 
kubectl delete svc nginx-ingress skaffold-demo 
kubectl delete secret sgune-cert nginx-ingress-serviceaccount-token-smwcv default-token-lcll5
minikube stop 