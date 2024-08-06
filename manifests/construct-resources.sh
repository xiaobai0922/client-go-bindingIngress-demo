kubectl create deployment ingress-controller --image bailu/ingress-controller:1.0.0 --dry-run=client -o yaml > ingress-controller-deployment.yaml

kubectl create serviceaccount ingress-controller --dry-run=client -o yaml > ingress-controller-serviceaccount.yaml

kubectl create clusterrole ingress-controller-clusterrole --resource=services,ingresses --verb=create,update,delete,get,list,watch --dry-run=client -o yaml > ingress-controller-clusterrole.yaml

kubectl create clusterrolebinding ingress-controller-clusterrolebinding --clusterrole=ingress-controller-clusterrole --serviceaccount=default:ingress-controller --dry-run=client -o yaml > ingress-controller-clusterrolebinding.yaml