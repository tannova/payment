#build stage
FROM registry.gitlab.com/greyhole/packages/hose:v1.0.0-hose-9f774daf

COPY ./falcon/configs /app
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]