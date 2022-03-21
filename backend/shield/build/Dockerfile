#build stage
FROM registry.gitlab.com/greyhole/packages/myid:v1.61.6-myid-84fbdafa

COPY ./shield/configs /app
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]
