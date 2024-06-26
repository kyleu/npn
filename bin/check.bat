@ECHO OFF

rem Runs code statistics, checks for outdated dependencies, then runs linters

cd %~dp0\..

echo "=== linting ==="
@ECHO ON
golangci-lint run --fix --max-issues-per-linter=0 --sort-results --skip-dirs "/vue" ./...
