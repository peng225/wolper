KIND_VERSION = v0.14.0
HELM_VERSION = v3.9.2
KIND = bin/kind
HELM = bin/helm

.PHONY: test
test:
	go test -v

.PHONY: setup
setup: $(KIND) $(HELM)
	$(KIND) create cluster
	$(HELM) install wolper ../charts/wolper --set deployment.app.replicaCount=1
	sleep 20
	kubectl get all -n wolper
	kubectl wait pods -n wolper -l app.kubernetes.io/name=wolper --for condition=Ready --timeout=30s
	kubectl port-forward -n wolper service/wolper-web-service 8080:8080 &

$(KIND): | bin
	curl -Lo ./kind https://kind.sigs.k8s.io/dl/$(KIND_VERSION)/kind-linux-amd64
	chmod +x ./kind
	mv ./kind $@

$(HELM): | bin
	curl -fsSL https://get.helm.sh/helm-$(HELM_VERSION)-linux-amd64.tar.gz | tar xzv
	mv linux-amd64/helm $@
	rm -rf linux-amd64

bin:
	mkdir $@

.PHONY: clean
clean:
	pkill kubectl
	$(KIND) delete cluster
