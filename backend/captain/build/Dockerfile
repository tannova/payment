#build stage
FROM registry.gitlab.com/greyhole/packages/myrole:v1.15.0-myrole-c629a4fa

COPY ./captain/configs /app
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]