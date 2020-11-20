include .env
export

GOOS   := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

ENVOY_BUILD_UBUNTU_VERSION := d06dad145694f1a7a02b5c6d0c75b32f753db2dd
ISTIO_VERSION := 1.8.0
ISTIO_VERSION_CANARY := 1.4.1

KUSTOMIZE_VERSION := 3.5.4
KUSTOMIZE_BIN := docker run \
	--rm \
	--volume $(shell pwd):/usr/local/src/lab \
	--workdir /usr/local/src/lab \
	110y/kustomize:${KUSTOMIZE_VERSION}

.PHONY: istioctl
istioctl:
	curl -sSL "https://storage.googleapis.com/istio-release/releases/${ISTIO_VERSION}/istioctl-${ISTIO_VERSION}-${GOOS}-${GOARCH}.tar.gz" | tar -C /tmp -xzv istioctl
	mv /tmp/istioctl ./bin/istioctl

.PHONY: go-bazel-helloworld-binary
go-bazel-helloworld-binary:
	@cd ./go/bazel/container && \
		bazel build --stamp --workspace_status_command=$$PWD/bazel/status.sh //:binary

.PHONY: go-bazel-helloworld-image
go-bazel-helloworld-image:
	@cd ./go/bazel/container && \
		bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:image

.PHONY: go-bazel-helloworld-publish
go-bazel-helloworld-publish:
	@cd ./go/bazel/container && \
		bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:publish

.PHONY: pb
pb:
	docker-compose run --rm protoc \
		protoc \
		--proto_path=proto/echo \
		--js_out=import_style=commonjs:proto/echo \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:proto/echo \
		--go_out=plugins=grpc:proto/echo \
		proto/echo/*.proto
	docker-compose run --rm protoc \
		protoc \
		--proto_path=proto/grpcserver \
		--js_out=import_style=commonjs:proto/grpcserver \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:proto/grpcserver \
		--go_out=plugins=grpc:proto/grpcserver \
		proto/grpcserver/*.proto

./cc/emscripten/a.out.js: ./cc/emscripten/test.cc
	docker-compose run --rm emscripten /usr/local/src/emscripten/emcc -o ./cc/emscripten/a.out.js ./cc/emscripten/test.cc

.PHONY: envoy-filter-example
envoy-filter-example:
	docker run --rm \
	    --volume "${PWD}:/usr/local/src/workspace" \
	    --volume "${HOME}/.config/gcloud/application_default_credentials.json:/etc/google_application_default_credentials.json" \
	    --workdir /usr/local/src/workspace/_third_party/envoy-filter-example \
	    envoyproxy/envoy-build-ubuntu:${ENVOY_BUILD_UBUNTU_VERSION} \
	    bash -c "bazel build --google_credentials=/etc/google_application_default_credentials.json --remote_http_cache=https://storage.googleapis.com/$$BAZEL_REMOTE_HTTP_CACHE_GCS_STORAGE //http-filter-example:envoy && cp -f ./bazel-bin/http-filter-example/envoy ../../bin/envoy-filter-example"

.PHONY: envoy-filter-hello
envoy-filter-hello:
	docker run --rm \
	    --volume "${PWD}:/usr/local/src/workspace" \
	    --volume "${HOME}/.config/gcloud/application_default_credentials.json:/etc/google_application_default_credentials.json" \
	    --workdir /usr/local/src/workspace \
	    envoyproxy/envoy-build-ubuntu:${ENVOY_BUILD_UBUNTU_VERSION} \
	    bash -c "bazel build --google_credentials=/etc/google_application_default_credentials.json --remote_http_cache=https://storage.googleapis.com/$$BAZEL_REMOTE_HTTP_CACHE_GCS_STORAGE //envoy/filter/hello:envoy && cp -f ./bazel-bin/envoy/filter/hello/envoy ./bin/envoy-filter-hello"

.PHONY: envoy-wasm
envoy-wasm:
	docker-compose run --rm envoy-builder \
	    bash -c "cd /usr/local/src/workspace/_third_party/envoy-wasm && bazel build --google_credentials=/etc/google_application_default_credentials.json --remote_http_cache=https://storage.googleapis.com/$$BAZEL_REMOTE_HTTP_CACHE_GCS_STORAGE --define wasm=enabled //source/exe:envoy-static && cp -f ./bazel-bin/source/exe/envoy-static ../../bin/envoy-wasm"

.PHONY: envoy-wasm-sdk
envoy-wasm-sdk:
	docker build -t envoy-wasm-sdk:v2 -f ./_third_party/envoy-wasm/api/wasm/cpp/Dockerfile-sdk ./_third_party/envoy-wasm/api/wasm/cpp

.PHONY: envoy-wasm-filter
envoy-wasm-filter:
	docker run -v $$PWD/envoy/filter/wasm/:/work -w /work envoy-wasm-sdk:v2 /build_wasm.sh
	# docker run -v $$PWD/envoy/filter/wasm/:/work -w /work proxywasmsdk:v2 /build_wasm.sh

.PHONY: ubuntu
ubuntu:
	cd ./linux/ubuntu && vagrant up && vagrant ssh

.PHONY: ubuntu-reset
ubuntu-reset:
	cd ./linux/ubuntu && vagrant destroy -f && vagrant up && vagrant ssh

.PHONY: bpf-go
bpf-go:
	sudo go run ./linux/bpf/helloworld/hello.go &
	sudo cat /sys/kernel/debug/tracing/trace_pipe

istio-components: istio-discovery istio-autoinject istio-mixer istio-grafana istio-prometheus
istio-components-master: istio-discovery-master istio-autoinject-master istio-mixer-master

.PHONY: istio-crd
istio-crd:
	# @$(KUSTOMIZE_BIN) build ./_third_party/istio/manifests/base > ./kubernetes/crd/istio.yaml
	$(KUSTOMIZE_BIN) build ./_third_party/istio-installer/base > ./kubernetes/crd/istio.yaml

.PHONY: istio-citadel
istio-citadel:
	@cd ./_third_party/istio-installer && HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-system \
		citadel \
		./security/citadel \
		-t \
		> ../../kubernetes/istio-system/citadel.yaml
		# --set 'security.dnsCerts.istio-sidecar-injector-service-account\.istio-control-canary=istio-sidecar-injector\.istio-control-canary\.svc' \

.PHONY: istio-discovery
istio-discovery:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-control \
		istio-discovery-master \
		./istio-control/istio-discovery \
		-t \
		--set global.istioNamespace=istio-system \
		--set global.configNamespace=istio-control \
		--set global.telemetryNamespace=istio-telemetry \
		--set global.policyNamespace=istio-policy \
		--set pilot.policy.enabled=true \
		--set pilot.useMCP=false \
		--set global.mtls.enabled=false \
		> ../../kubernetes/istio-control/discovery.yaml

.PHONY: istio-autoinject
istio-autoinject:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-control \
		istio-autoinject-master \
		./istio-control/istio-autoinject \
		-t \
		--set sidecarInjectorWebhook.enableNamespacesByDefault=false \
		--set global.configNamespace=istio-control \
		--set global.disablePolicyChecks=false \
		--set global.telemetryNamespace=istio-telemetry \
		--set istio_cni.enabled=false \
		--set mixer.policy.enabled==true \
		> ../../kubernetes/istio-control/autoinject.yaml

.PHONY: istio-mixer
istio-mixer:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-telemetry \
		istio-mixer \
		./istio-telemetry/mixer-telemetry/ \
		-t \
		--set global.configNamespace=istio-control \
		--set global.istioNamespace=istio-system \
		--set global.telemetryNamespace=istio-telemetry \
		--set global.policyNamespace=istio-policy \
		--set global.mtls.enabled=false \
		--set policy.enable=true \
		--set mixer.telemetry.useMCP=false \
		> ../../kubernetes/istio-telemetry/mixer.yaml

.PHONY: istio-grafana
istio-grafana:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-telemetry \
		istio-grafana  \
		./istio-telemetry/grafana \
		-t \
		--set global.configNamespace=istio-control \
		--set global.mtls.enabled=false \
		--set policy.enable=true \
		--set mixer.telemetry.useMCP=false \
		> ../../kubernetes/istio-telemetry/grafana.yaml

.PHONY: istio-prometheus
istio-prometheus:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-telemetry \
		istio-prometheus \
		./istio-telemetry/prometheus \
		-t \
		--set global.configNamespace=istio-control \
		--set global.istioNamespace=istio-system \
		--set global.telemetryNamespace=istio-telemetry \
		--set global.policyNamespace=istio-policy \
		--set global.mtls.enabled=false \
		--set policy.enable=true \
		--set mixer.telemetry.useMCP=false \
		> ../../kubernetes/istio-telemetry/prometheus.yaml

.PHONY: istio-policy
istio-policy:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-policy \
		istio-policy \
		./istio-policy \
		-t \
		--set global.configNamespace=istio-control \
		--set global.istioNamespace=istio-system \
		--set global.telemetryNamespace=istio-telemetry \
		--set global.policyNamespace=istio-policy \
		--set global.mtls.enabled=false \
		--set policy.enable=true \
		--set global.useMCP=false \
		> ../../kubernetes/istio-policy/mixer.yaml

.PHONY: istio-discovery-master
istio-discovery-master:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control-master HUB=docker.io/istio TAG=${ISTIO_VERSION_CANARY} ./bin/iop \
		istio-control-master \
		istio-discovery \
		./istio-control/istio-discovery \
		-t \
		--set global.istioNamespace=istio-system \
		--set global.configNamespace=istio-control-master \
		--set global.telemetryNamespace=istio-telemetry-master \
		--set global.policyNamespace=istio-policy-master \
		--set pilot.useMCP=false \
		--set policy.enable=false \
		--set global.mtls.enabled=false \
		> ../../kubernetes/istio-control-master/discovery.yaml

.PHONY: istio-autoinject-master
istio-autoinject-master:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control-master HUB=docker.io/istio TAG=${ISTIO_VERSION_CANARY} ./bin/iop \
		istio-control-master \
		istio-autoinject \
		./istio-control/istio-autoinject \
		-t \
		--set sidecarInjectorWebhook.enableNamespacesByDefault=false \
		--set global.configNamespace=istio-control-master \
		--set global.telemetryNamespace=istio-telemetry-master \
		--set istio_cni.enabled=false \
		> ../../kubernetes/istio-control-master/autoinject.yaml

.PHONY: istio-mixer-master
istio-mixer-master:
	@cd ./_third_party/istio-installer && \
		ISTIO_ENV=istio-control-master HUB=docker.io/istio TAG=${ISTIO_VERSION_CANARY} ./bin/iop \
		istio-telemetry-master \
		istio-mixer \
		./istio-telemetry/mixer-telemetry/ \
		-t \
		--set global.configNamespace=istio-control-master \
		--set global.istioNamespace=istio-system \
		--set global.telemetryNamespace=istio-telemetry-master \
		--set global.policyNamespace=istio-policy-master \
		--set global.mtls.enabled=false \
		--set policy.enable=false \
		--set mixer.telemetry.useMCP=false \
		> ../../kubernetes/istio-telemetry-master/mixer.yaml

.PHONY: istio-config
istio-config:
	@cd ./_third_party/istio-installer && \
		HUB=docker.io/istio TAG=${ISTIO_VERSION} ./bin/iop \
		istio-control \
		istio-config \
		./istio-control/istio-config \
		--set configValidation=true

.PHONY: kind-cluster
kind-cluster:
	kind delete cluster --name lab
	kind create cluster --name lab --image kindest/node:v1.19.1
	make lab-istio
	# skaffold run --filename=./kubernetes/skaffold/skaffold.yaml
	docker build -t ghcr.io/110y/lab/grpcserver1:latest ./go/grpc/grpcserver1
	kind load docker-image ghcr.io/110y/lab/grpcserver1:latest --name lab
	docker build -t ghcr.io/110y/lab/grpcserver2:latest ./go/grpc/grpcserver2
	kind load docker-image ghcr.io/110y/lab/grpcserver2:latest --name lab
	kubectl apply -f kubernetes/grpcserver1/deployment.yaml
	kubectl apply -f kubernetes/grpcserver2/deployment.yaml
	kubectl config set-context $(shell kubectl config current-context) --namespace=grpcserver1

.PHONY: lab-istio
lab-istio:
	./bin/istioctl install -y

.PHONY: controller-runtime-example-container
controller-runtime-example-container:
	docker build -f ./kubernetes/controller-runtime/Dockerfile -t registry:5000/controller-runtime-example:latest .
	docker push registry:5000/controller-runtime-example:latest

.PHONY: gotools
gotools:
	go build -o ./bin/envsubst github.com/a8m/envsubst/cmd/envsubst

.PHONY: authority-container
authority-container:
	docker build -t $${CONTAINER_REGISTRY}/authority:$(shell git rev-parse --short HEAD) -f ./envoy/authority/Dockerfile .

.PHONY: authority-registry
authority-registry: authority-container
	docker push $${CONTAINER_REGISTRY}/authority:$(shell git rev-parse --short HEAD)

.PHONY: authority-deploy
authority-deploy:
	docker-compose run --rm go ./bin/envsubst < ./kubernetes/authority/deployment.yaml | kubectl apply -f -

.PHONY: kube-istio-lab
kube-istio-lab:
	docker-compose run --rm go ./bin/envsubst < ./kubernetes/istio-lab/origin-service.yaml | kubectl apply -f -
	docker-compose run --rm go ./bin/envsubst < ./kubernetes/istio-lab/replication-service.yaml | kubectl apply -f -
	docker-compose run --rm go ./bin/envsubst < ./kubernetes/istio-lab/destination.yaml | kubectl apply -f -
	docker-compose run --rm go ./bin/envsubst < ./kubernetes/istio-lab/virtual_service.yaml | kubectl apply -f -

.PHONY: rust-hello
rust-hello:
	rustc -o bin/rust_hello ./rust/hello/main.rs
