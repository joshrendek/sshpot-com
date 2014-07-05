all:
	go test
	go build
	./sshd_honey_web
	
deploy:
	git push heroku master
