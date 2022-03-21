kubectl delete svc --namespace=prod kimpton-echoserver-external
kubectl delete svc --namespace=prod kimpton-echoserver
kubectl delete svc --namespace=prod kimpton-echoweb
kubectl delete deploy --namespace=prod kimpton-echoserver
kubectl delete deploy --namespace=prod kimpton-echoclient
kubectl delete deploy --namespace=prod kimpton-echoweb