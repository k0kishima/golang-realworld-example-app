.PHONY: test-e2e

test-e2e:
	@echo "Running end-to-end tests..."
	@APIURL=http://localhost:8080/api ./tests/e2e/run-api-tests.sh
