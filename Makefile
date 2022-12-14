build:
	@go build -o mytens.exe main.go

v%:
	@git tag v$(*)
	make m$(*)
	@auto-changelog
	@git tag -d v$(*)
	@git add .
	@git commit -m "Update v$(*)"
	@git tag v$(*)

m%:
	@echo '{"version": "v$(*)","description": "Simple command to implement git flow and other command when developing MyTEnS, especially service-api-data","homepage": "https://buatan.id","bin": "mytens.exe","architecture": {"64bit": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/v$(*)/mytens-v$(*)-windows-amd64.zip"},"32bit": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/v$(*)/mytens-v$(*)-windows-386.zip"},"arm64": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/v$(*)/mytens-v$(*)-windows-arm.zip"}},"checkver": {"github": "https://github.com/buatan/service-api-data-cli","regex": "tag/(v[\\\\w.-]+)"},"autoupdate": {"architecture": {"64bit": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/$$version/mytens-$$version-windows-amd64.zip"},"32bit": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/$$version/mytens-$$version-windows-386.zip"},"arm64": {"url": "https://github.com/buatan/service-api-data-cli/releases/download/$$version/mytens-$$version-windows-arm.zip"}}}}' > bucket/mytens.json

.PHONY: build