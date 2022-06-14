VERSION := 1.0

#docker tag msa-otus-api-amd64:1.0 ekod/msa-otus
#docker push ekod/msa-otus

msa-otus:
	docker build \
		-f zarf/docker/dockerfile.msa-otus-api \
		-t msa-otus-api-amd64 \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

tidy:
	go mod tidy
	go mod vendor

docker-deploy:
	docker tag msa-otus-api-amd64 ekod/msa-otus
	docker push ekod/msa-otus
apply:
	kubectl apply -f zarf/k8s/

delete:
	kubectl delete -f zarf/k8s/