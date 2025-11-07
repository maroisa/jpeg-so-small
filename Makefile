PLATFORM ?= linux/amd64

jpeg-so-small:
	wails build -webview2 embed -tags webkit2_41 -upx -platform $(PLATFORM)
