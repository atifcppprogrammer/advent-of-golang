setup:
	@chmod +x ./scripts/*
	@chmod +x .git-hooks/*
	@cp .git-hooks/* .git/hooks

init: ./scripts/makefile-init.sh
	@./scripts/makefile-init.sh $(YEAR) $(DAY) $(NAME)

run:
	@go run ./solutions/year-$(YEAR)/day-$$(printf %02d $(DAY)) $(INPUT)

test: 
	@go test -v ./packages/*
	@go test -v ./solutions/*/*

solution:
	@git add go.work && git add solutions; \
	git commit -m "solution: year-$(YEAR)/day-$$(printf %02d $(DAY))"

improve:
	@git add go.work && git add solutions; \
	./scripts/makefile-improve.sh $(YEAR) $(DAY)

format: 
	@go work edit -fmt
	@gofmt -l -s -w  .
