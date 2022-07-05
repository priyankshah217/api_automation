golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2
	./bin/golangci-lint run ./...

go-test-run:
	go install gotest.tools/gotestsum@latest
	@mkdir -vp "./target" "./target/allure-results"
	~/go/bin/gotestsum --junitfile ./target/allure-results/junit-tests.xml -- ./tests/...
generate-allure-report:
	allure generate ./target/allure-results -o ./target/allure-report
	allure open ./target/allure-report