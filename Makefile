############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

-----------------------------: ##

___ANVIL_STATE___: ##

init:
	@chmod +x ./init.sh
	./init.sh  # TODO: modify init.sh to actually install all stuff

build-anvil-state-with-deployed-contracts: ## builds anvil state with deployed contracts and generates a state
	@chmod +x ./contracts/anvil/build-state.sh
	./contracts/anvil/build-state.sh

start-anvil:
	@chmod +x ./contracts/anvil/start-anvil.sh
	./contracts/anvil/start-anvil.sh

___CONTRACTS___: ##

build-contracts: ## builds all contracts and generates ABIs
	cd contracts && forge build
	cd contracts && forge inspect SertnTaskManager abi --json > ../abis/SertnTaskManager.abi.json
	cd contracts && forge inspect SertnServiceManager abi --json > ../abis/SertnServiceManager.abi.json

deploy-eigenlayer-contracts:
	@chmod +x ./contracts/anvil/deploy-el.sh
	./contracts/anvil/deploy-el.sh

deploy-sertn-contracts:
	@chmod +x ./contracts/anvil/deploy-sertn.sh
	./contracts/anvil/deploy-sertn.sh

add-operator-model:
	pwd # TODO: ...

__CLI__: ##

send-fund: ## sends fund to the operator saved in tests/keys/test.ecdsa.key.json
	cast send 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5 --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

-----------------------------: ##
# We pipe all zapper logs through https://github.com/maoueh/zap-pretty so make sure to install it
# TODO: piping to zap-pretty only works when zapper environment is set to production, unsure why
____OFFCHAIN_SOFTWARE___:
start-operator: ## start operator (part of quickstart)
	uv run task_operator.py

start-aggregator: ## start aggregator (part of quickstart)
	uv run aggregator.py

-----------------------------: ##
_____HELPER_____: ##
tests-contract: ## runs all forge tests
	cd contracts && forge test
