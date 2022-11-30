build:
	@go build -o mytens.exe main.go

v%:
	@git tag v$(*)
	@auto-changelog
	@git tag -d v$(*)
	@echo '{"version": "v$(*)","description": "Simple command to implement git flow and other command when developing MyTEnS, especially service-api-data","homepage": "https://buatan.id","bin": "mytens-v$(*)-windows-amd64.exe","architecture": {"64bit": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/v$(*)/mytens-v$(*)-windows-amd64.exe"},"32bit": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/v$(*)/mytens-v$(*)-windows-386.exe"},"arm64": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/v$(*)/mytens-v$(*)-windows-arm.exe"}}\n}' > mytens.json
	@git add .
	@git commit -m "Update v$(*)"
	@git tag v$(*)

.PHONY: build