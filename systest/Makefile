test_name ?= TestSmeshing
version_info ?= $(shell git rev-parse --short HEAD)
org ?= spacemeshos
image_name ?= $(org)/systest:$(version_info)
smesher_image ?= $(org)/go-spacemesh-dev:$(version_info)
test_id ?= systest-$(version_info)
test_pod_name ?= systest-$(version_info)-$(shell date +'%s')
keep ?= false
clusters ?= 1
size ?= 10
level ?= debug
bootstrap ?= 5m
storage ?= standard=1Gi
node_selector ?=
namespace ?=
label ?=


.PHONY: docker
docker:
	@DOCKER_BUILDKIT=1 docker build ../ -f Dockerfile -t $(image_name)

.PHONY: push
push:
	docker push $(image_name)

.PHONY: run
run:
	@echo "launching test pod with name=$(test_pod_name) and testid=$(test_id)"
	@kubectl run --image $(image_name) $(test_pod_name) \
	--restart=Never \
	--labels="testid=$(test_id)" \
    --image-pull-policy=IfNotPresent -- \
    tests -test.v -test.timeout=0 -test.run=$(test_name) -namespace=$(namespace) -clusters=$(clusters) -storage=$(storage) -size=$(size) -image=$(smesher_image) -level=$(level) -node-selector=$(node_selector) -bootstrap=$(bootstrap) -keep=$(keep) -testid=$(test_id) -labels=$(label)
	-@kubectl wait --timeout=5s --for=condition=ready pod/$(test_pod_name)
	@kubectl logs $(test_pod_name) -f

.PHONY: clean
clean:
	@echo "deleting test pods with testid=$(test_id)"
	@kubectl delete pod --selector=testid=$(test_id)
	@echo "deleting ephemeral namespaces with testid=$(test_id)"
	@kubectl delete ns --selector=testid=$(test_id),keep=false

.PHONY: cleanall
cleanall: clean
	@echo "deleting all namespaces with testid=$(test_id)"
	@kubectl delete ns --selector=testid=$(test_id)
