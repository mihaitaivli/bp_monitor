generate:
	@echo "updating resolvers to match the schema"
	gqlgen generate

develop:
	@echo "starting the server ..."
	go run server.go