.PHONY: rabbit clean

rabbit:
	go build -o out/rabiit ./cmd/rabbit.go

clean:
	rm -rf ./out

generate:
	bash ./hack/update-codegen.sh