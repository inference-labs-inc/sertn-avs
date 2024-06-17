############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

FORK_RPC=https://holesky.gateway.tenderly.co
CHAINID=17000

AGGREGATOR_ECDSA_PRIV_KEY=0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
CHALLENGER_ECDSA_PRIV_KEY=0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
MAIN_ECDSA_KEY=${AGGREGATOR_ECDSA_PRIV_KEY}
OPERATOR_ECDSA_PRIV_KEY=c6373064c6eb0c3b09d55ec5a85cf9a39003ffc39245157584d8eeed8bed041d

DEPLOYMENT_FILES_DIR=contracts/script/output/${CHAINID}

WETH_ADDRESS=0x94373a4919b3240d86ea41593d5eba789fef3848
STRATEGY_ADDRESS=0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9

OPERATOR_ADDRESS=0x6dBC2B9174B0b51B7B308e064358a31E50beeBfa
CHALLENGER_ADDRESS=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
AGGREGATOR_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
-----------------------------: ## 

___CONTRACTS___: ## 
setup-python:
	python3 -m venv ./.venv && source ./.venv/bin/activate && pip install -r ./python/requirements.txt

start-interface: ## builds all contracts
	cd sertn-fe && npm i && npm run dev

build-contracts: ## builds all contracts
	cd contracts && forge build --sizes

setup: setup-python build-contracts

deploy-contracts: 
	./tests/anvil/deploy-avs-save-anvil-state.sh

start-chain: ## starts anvil from a saved state file (with el and avs contracts deployed)
	anvil --fork-url ${FORK_RPC}

bindings: ## generates contract bindings
	cd contracts && ./generate-go-bindings.sh

__CLI__: ## 

cli-setup-operator: send-fund cli-register-operator-with-eigenlayer cli-deposit-into-mocktoken-strategy cli-register-operator-with-avs ## registers operator with eigenlayer and avs

cli-register-operator-with-eigenlayer: ## registers operator with delegationManager
	go run cli/main.go --config config-files/operator.anvil.yaml register-operator-with-eigenlayer

cli-deposit-into-mocktoken-strategy: ## 
	./scripts/deposit-into-mocktoken-strategy.sh

cli-register-operator-with-avs: ## 
	go run cli/main.go --config config-files/operator.anvil.yaml register-operator-with-avs

cli-deregister-operator-with-avs: ## 
	go run cli/main.go --config config-files/operator.anvil.yaml deregister-operator-with-avs

cli-print-operator-status: ## 
	go run cli/main.go --config config-files/operator.anvil.yaml print-operator-status

send-funds-operator: 
	cast send ${OPERATOR_ADDRESS} --value 0.25ether --private-key ${MAIN_ECDSA_KEY}

send-funds: ## sends fund to the operator challwnger and aggregator saved in tests/keys/test.ecdsa.key.json
	cast send ${AGGREGATOR_ADDRESS} --value 0.25ether --private-key ${MAIN_ECDSA_KEY} && cast send ${CHALLENGER_ADDRESS} --value 0.25ether --private-key ${MAIN_ECDSA_KEY}

wrap-eth:
	cast send ${WETH_ADDRESS} "deposit()" --value 0.125ether --private-key ${OPERATOR_ECDSA_PRIV_KEY}

-----------------------------: ## 
# We pipe all zapper logs through https://github.com/maoueh/zap-pretty so make sure to install it
# TODO: piping to zap-pretty only works when zapper environment is set to production, unsure why
____OFFCHAIN_SOFTWARE___: ## 
start-aggregator: ## 
	source ./.venv/bin/activate && go run aggregator/cmd/main.go --config config-files/aggregator.yaml \
		--sertn-deployment ${DEPLOYMENT_FILES_DIR}/sertn_avs_deployment_output.json \
		--ecdsa-private-key ${AGGREGATOR_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

start-operator: send-funds-operator wrap-eth start-operator-go

start-operator-go: ## 
	source ./.venv/bin/activate && go run operator/cmd/main.go --config config-files/operator.anvil.yaml \
		2>&1 | zap-pretty

start-challenger: ## 
	source ./.venv/bin/activate && go run challenger/cmd/main.go --config config-files/challenger.yaml \
		--sertn-deployment ${DEPLOYMENT_FILES_DIR}/sertn_avs_deployment_output.json \
		--ecdsa-private-key ${CHALLENGER_ECDSA_PRIV_KEY} \
		2>&1 | zap-pretty

run-plugin: ## 
	go run plugin/cmd/main.go --config config-files/operator.anvil.yaml
-----------------------------: ## 
_____HELPER_____: ## 
mocks: ## generates mocks for tests
	go install go.uber.org/mock/mockgen@v0.3.0
	go generate ./...

tests-unit: ## runs all unit tests
	go test $$(go list ./... | grep -v /integration) -coverprofile=coverage.out -covermode=atomic --timeout 15s
	go tool cover -html=coverage.out -o coverage.html

tests-contract: ## runs all forge tests
	cd contracts && forge test

tests-integration: ## runs all integration tests
	go test ./tests/integration/... -v -count=1

