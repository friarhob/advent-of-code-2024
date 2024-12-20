# Variables
INPUT_DIR := inputs
EXAMPLES_DIR := examples
SOLUTIONS_DIR := solutions

# Targets
.PHONY: all run-day run-example run-all tidy clean

# Run a specific day's solution
run-day:
	@echo "Running solution for Day $(DAY)..."
	go run ./$(SOLUTIONS_DIR)/day$(DAY)/. -input $(INPUT_DIR)/day$(DAY).txt

run-example:
	@echo "Running example test for Day $(DAY)..."
	go run ./$(SOLUTIONS_DIR)/day$(DAY)/. -input $(EXAMPLES_DIR)/day$(DAY).txt

# Run all solutions
run-all:
	@for day in $(shell ls $(SOLUTIONS_DIR)); do \
		echo "Running solution for $$day..."; \
		go run ./$(SOLUTIONS_DIR)/$$day/. -input $(INPUT_DIR)/$$day.txt; \
	done

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Clean generated files
clean:
	@echo "Cleaning up test cache..."
	go clean -testcacheHow