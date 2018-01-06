all: clean
	go run main.go

clean:
	mkdir -p Site
	rm -f Site/*

deploy:
	aws s3 sync ./Site/ s3://r.32k.io/ --acl public-read --content-type text/html --profile 32k
