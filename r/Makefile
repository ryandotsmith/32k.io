all: clean
	md2html
	aws s3 sync ./site/ s3://r.32k.io/ --acl public-read --content-type text/html --profile 32k
	aws cloudfront create-invalidation --distribution-id E1GZ3VINCGTLGF --paths '/*' --profile 32k

clean:
	mkdir -p ./site
	rm -f ./site/*
