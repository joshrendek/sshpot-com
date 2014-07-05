all:
	go test
	go build
	
deploy:
	git push heroku master
