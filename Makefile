build:
	go build -o go-web-ui.exe main.go

clean:
	del /q/s *.exe

.PHONY: clean