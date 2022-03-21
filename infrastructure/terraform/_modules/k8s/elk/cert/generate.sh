ELASTICSEARCH_IMAGE="docker.elastic.co/elasticsearch/elasticsearch:7.11.1" && \
aws eks --region ap-southeast-1 update-kubeconfig --name ${cluster_name} --profile ${profile} && \
docker rm -f elastic-helm-charts-certs || true && \
rm -f elastic-certificates.p12 elastic-certificate.pem elastic-certificate.crt elastic-stack-ca.p12 || true && \
password=$(docker run --rm busybox:1.31.1 /bin/sh -c "< /dev/urandom tr -cd '[:alnum:]' | head -c20")
encryptionkey=$(docker run --rm busybox:1.31.1 /bin/sh -c "< /dev/urandom tr -cd '[:alnum:]' | head -c32")
docker run --name elastic-helm-charts-certs -i -w /tmp \
		$ELASTICSEARCH_IMAGE \
		/bin/sh -c " \
			elasticsearch-certutil ca --out /tmp/elastic-stack-ca.p12 --pass '' && \
			elasticsearch-certutil cert --name elasticsearch-master --dns elasticsearch-master --ca /tmp/elastic-stack-ca.p12 --pass '' --ca-pass '' --out /tmp/elastic-certificates.p12" && \ 
docker cp elastic-helm-charts-certs:/tmp/elastic-certificates.p12 ./  && \
docker rm -f elastic-helm-charts-certs || true && \
openssl pkcs12 -nodes -passin pass:'' -in elastic-certificates.p12 -out elastic-certificate.pem  && \
openssl x509 -outform der -in elastic-certificate.pem -out elastic-certificate.crt  && \
kubectl create namespace elk || true  && \
kubectl delete secret  elastic-certificates -n elk|| true  && \
kubectl delete secret  elastic-certificate-pem -n elk|| true  && \
kubectl delete secret  elastic-certificate-crt -n elk|| true && \
kubectl delete secret  elastic-credentials -n elk || true && \
kubectl delete secret  kibana -n elk || true && \

kubectl create secret generic elastic-certificates --from-file=elastic-certificates.p12 -n elk  && \
kubectl create secret generic elastic-certificate-pem --from-file=elastic-certificate.pem -n elk && \
kubectl create secret generic elastic-certificate-crt --from-file=elastic-certificate.crt -n elk  && \
kubectl create secret generic elastic-credentials --from-literal=password=$password --from-literal=username=elastic -n elk && \
kubectl create secret generic kibana --from-literal=encryptionkey=$encryptionkey --from-literal=username=elastic -n elk && \
rm -f elastic-certificates.p12 elastic-certificate.pem elastic-certificate.crt