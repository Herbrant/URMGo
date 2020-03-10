PROJECTNAME := "URMGo"

# Go related variables.
GOBIN := ./bin
GOFILES := ./src/

build:
	@echo "  >  Building binary..."
	@go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)
	@echo "  >  Done."

clean:
	@echo "  >  Cleaning..."
	@rm -f $(GOBIN)/*
	@echo "  >  Done."