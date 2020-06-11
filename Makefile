.PHONY: generate

generate:
	@go generate ./...
	@echo "[OK] Files added to embed box!"