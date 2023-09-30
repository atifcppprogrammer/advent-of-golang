setup:
	@chmod +x ./init-solution.sh
	@chmod +x .git-hooks/*
	@cp .git-hooks/* .git/hooks

init: init-solution.sh
	@./init-solution.sh $(YEAR) $(DAY) $(NAME)

run:
	@go run ./solutions/year-$(YEAR)/day-$$(printf %02d $(DAY)) $(PATH)

test: 
	@go test -v ./packages/*
	@go test -v ./solutions/*/*

solution:
	@git add go.work && git add solutions; \
	git commit -m "solution: year-$(YEAR)/day-$$(printf %02d $(DAY))"

format: 
	@go work edit -fmt
	@gofmt -l -s -w  .
