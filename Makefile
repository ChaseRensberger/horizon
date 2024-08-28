build:
	docker build -t horizon .
	docker tag horizon chaserensberger/horizon:latest
	docker push chaserensberger/horizon:latest
