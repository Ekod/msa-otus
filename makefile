VERSION := 1.0

all: msa-otus

msa-otus:
	docker build \
		-f zarf/docker/dockerfile.msa-otus-api \
		-t msa-otus-api-amd64 \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.


KIND_CLUSTER := msa-otus-cluster

kind-up:
	kind create cluster \
		--image kindest/node:v1.24.0@sha256:0866296e693efe1fed79d5e6c7af8df71fc73ae45e3679af05342239cdc5bc8e \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml
	kubectl config set-context --current --namespace=msa-otus-system

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	kind load docker-image msa-otus-api-amd64 --name $(KIND_CLUSTER)

kind-apply:
	cat zarf/k8s/base/msa-otus-pod/base-service.yaml | kubectl apply -f -

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

tidy:
	go mod tidy
	go mod vendor

kind-logs:
	kubectl logs -l app=msa-otus --all-containers=true -f --tail=100 --namespace=msa-otus-system

docker-deploy:
	docker tag msa-otus-api-amd64 ekod/msa-otus
	docker push ekod/msa-otus
apply:
	kubectl apply -f zarf/k8s/

delete:
	kubectl delete -f zarf/k8s/

build-deploy: msa-otus docker-deploy

kind-restart:
	kubectl rollout restart deployment msa-otus-pod

u: msa-otus kind-up kind-load kind-apply

	