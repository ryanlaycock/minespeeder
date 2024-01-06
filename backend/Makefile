dependencies:
	echo "Installing dependencies..."
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

generate:
	echo "Generating..."
	oapi-codegen -package api --config=api/config.yaml ../MineSpeederAPI.json > api/minespeeder.gen.go
	echo "Done"