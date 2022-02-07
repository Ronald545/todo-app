dev: 
	concurrently "npm run dev --prefix=frontend" "reflex -r '\\.go' -s -- go run main.go"

build: 
	npm run build --prefix=frontend && go build main.go

start:
	concurrently "npm run preview --prefix=frontend" "./main"
