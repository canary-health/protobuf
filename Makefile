.DEFAULT_GOAL := pb

########################
## Helpers variables
########################
M=$(shell printf "\033[34;1m▶\033[0m")
TIMESTAMP := $(shell /bin/date "+%Y-%m-%d_%H-%M-%S")

#####
## Protobuf commands
#####
.PHONY: pb 

pb: rm-pb; $(info $(M) Generating protobuf files ...)
	go generate ./types/gen.go

######
## Clean up commands
######
.PHONY: rm-pb rm-tools
rm-pb: ; $(info $(M) Removing generated protobuf files... )
	for pb in ./types/*.pb.go; do \
		$(RM) $$pb; \
		echo Removed: $$pb; \
	done

rm-tools: ; $(info $(M) Removing ./_tools files... )
	sudo rm -rf ./_tools