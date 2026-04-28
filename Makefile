DIST := dist

.PHONY: build clean

build:
	GOOS=windows GOARCH=amd64 go build -o $(DIST)/gitviz.exe . && zip -j $(DIST)/gitviz_windows_amd64.zip $(DIST)/gitviz.exe     && rm $(DIST)/gitviz.exe
	GOOS=darwin  GOARCH=amd64 go build -o $(DIST)/gitviz     . && zip -j $(DIST)/gitviz_darwin_amd64.zip  $(DIST)/gitviz         && rm $(DIST)/gitviz
	GOOS=darwin  GOARCH=arm64 go build -o $(DIST)/gitviz     . && zip -j $(DIST)/gitviz_darwin_arm64.zip  $(DIST)/gitviz         && rm $(DIST)/gitviz
	GOOS=linux   GOARCH=amd64 go build -o $(DIST)/gitviz     . && zip -j $(DIST)/gitviz_linux_amd64.zip   $(DIST)/gitviz         && rm $(DIST)/gitviz
	GOOS=linux   GOARCH=arm64 go build -o $(DIST)/gitviz     . && zip -j $(DIST)/gitviz_linux_arm64.zip   $(DIST)/gitviz         && rm $(DIST)/gitviz

clean:
	rm -rf $(DIST)
