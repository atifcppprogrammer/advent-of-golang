define stage_solution
	@git add go.work
	@git add "solutions/year-$(1)/day-$$(printf %02d $(2))"
endef

define run_solution
	@PART=$(4) go run ./solutions/year-$(1)/day-$$(printf %02d $(2)) $(3)
endef

.PHONY: setup init run test solution improve format version

setup:
	@chmod +x ./scripts/* 
	@chmod +x .git-hooks/* && cp .git-hooks/* .git/hooks

init: ./scripts/makefile-init.sh
	@./scripts/makefile-init.sh $(YEAR) $(DAY) $(NAME)

run:
ifdef PART
	$(call run_solution,$(YEAR),$(DAY),$(INPUT),$(PART))
else
	$(call run_solution,$(YEAR),$(DAY),$(INPUT),1)
endif

test: ./scripts/makefile-test.sh
	@./scripts/makefile-test.sh

solution:
	$(call stage_solution,$(YEAR),$(DAY)); \
	git commit -m "solution: year-$(YEAR)/day-$$(printf %02d $(DAY))"

improve:
	$(call stage_solution,$(YEAR),$(DAY)); \
	./scripts/makefile-improve.sh $(YEAR) $(DAY)

format: 
	@go work edit -fmt
	@gofmt -l -s -w  .

version: ./scripts/makefile-version.sh
	@./scripts/makefile-version.sh $(GO)