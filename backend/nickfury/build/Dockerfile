#build stage
FROM registry.gitlab.com/greyhole/packages/mylog:v1.4.1-mylog-f3efbd12

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.6 && \
    wget -qO /grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /grpc_health_probe
COPY ./nickfury/configs /app
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]
