generate:
	@echo "updating resolvers to match the schema"
	go run github.com/99designs/gqlgen generate

develop:
	@echo "starting the server ..."
	go run server.go