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

.PHONY: envoy-filter-example
envoy-filter-example:
	docker run --rm \
	    --volume "${PWD}:/usr/local/src/workspace" \
	    --volume "${HOME}/.config/gcloud/application_default_credentials.json:/etc/google_application_default_credentials.json" \
	    --workdir /usr/local/src/workspace/_third_party/envoy-filter-example \
	    envoyproxy/envoy-build-ubuntu:d06dad145694f1a7a02b5c6d0c75b32f753db2dd \
	    bazel build \
	    --google_credentials=/etc/google_application_default_credentials.json \
	    --remote_http_cache=https://storage.googleapis.com/$$BAZEL_REMOTE_HTTP_CACHE_GCS_STORAGE \
	    //:envoy
