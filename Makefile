build:
	
docker:
	docker build -t harry671003/kube-state-metrics-mock .

push:
	docker push harry671003/kube-state-metrics-mock:latest