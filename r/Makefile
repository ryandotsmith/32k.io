all: clean
	md2html
	aws-vault exec 32k -- aws s3 sync ./site/ s3://r.32k.io/ --acl public-read --content-type text/html

clean:
	mkdir -p ./site
	rm -f ./site/*
