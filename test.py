import os

os.system("go test ./src/... -cover -v -coverprofile=c.out && go tool cover -html=c.out -o coverage.html")