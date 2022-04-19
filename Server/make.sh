GOOS=windows GOARCH=amd64  go build -ldflags  "-w -s" -o ./bin/win/release/x64/akhomelab.exe ./src
GOOS=windows GOARCH=386  go build -ldflags  "-w -s" -o ./bin/win/release/win32/akhomelab.exe ./src
GOOS=linux GOARCH=amd64  go build -ldflags  "-w -s" -o ./bin/linux/release/x64/akhomelab ./src
GOOS=linux GOARCH=386  go build -ldflags  "-w -s" -o ./bin/linux/release/win32/akhomelab ./src
GOOS=darwin GOARCH=amd64  go build -ldflags  "-w -s" -o ./bin/osx/release/x64/akhomelab ./src

