FROM envoyproxy/envoy:v1.18.4 as builder
ARG service
#Backend dir context
WORKDIR /app
ADD $service /app/
# Create if not exist
RUN mkdir -p /app/scripts && test -f /app/envoy.yaml || touch /app/envoy.yaml
FROM envoyproxy/envoy:v1.18.4
WORKDIR /app
RUN apt-get update && apt-get install -y gettext \
  && rm -rf /var/lib/apt/lists/*
# For default
COPY --from=builder /app/envoy.yaml /etc/envoy.yaml
COPY --from=builder /app/descriptors* /etc/descriptors
COPY --from=builder /app/scripts /app/scripts
EXPOSE 8080 8081

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value