rm -rf dist
mkdir dist

go build -o my-go-job cmd/main.go
mv my-go-job dist/
cp cmd/env/.env.json dist/