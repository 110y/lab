FROM istio/proxyv2:1.7.3 AS istio-proxy

# FROM envoyproxy/envoy:v1.15.2
# COPY --from=istio-proxy /usr/local/bin/envoy /usr/local/bin/envoy

ENTRYPOINT /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l debug --service-cluster proxy
