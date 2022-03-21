#build stage
FROM registry.gitlab.com/greyhole/packages/postman:v1.0.3-postman-faafa85c

COPY ./hawkeye/configs /app
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]