VERSION := 1.0

#docker tag msa-otus-api-amd64:1.0 ekod/msa-otus
#docker push ekod/msa-otus

msa-otus:
	docker build \
		-f zarf/docker/dockerfile.msa-otus-api \
		-t msa-otus-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

tidy:
	go mod tidy
	go mod vendor

docker-deploy:
	docker tag msa-otus-api-amd64:$(VERSION) ekod/msa-otus:$(VERSION)
	docker push ekod/msa-otus:$(VERSION)