build:
	go build -o mytens.exe main.go

v%:
	git tag v$(*)
	auto-changelog
	git tag -d v$(*)
	git add .
	git commit -m "Update v$(*)"
	git tag v$(*)

.PHONY: build