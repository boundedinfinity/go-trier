makefile_dir		:= $(abspath $(shell pwd))

.PHONY: test publish

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

test:
	go test ./...

publish:
	git add . || true
	git commit -m "$(tag)"
	git push origin master
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)
