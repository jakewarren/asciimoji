all:
	@go-bindata emojis.json
	@go install