// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import "@eigenlayer/contracts/strategies/StrategyBase.sol";
import "../../src/ERC20Mock.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract Utils is Script {
    // Note that this fct will only work for the ERC20Mock that has a public mint function
    function _mintTokens(
        address strategyAddress,
        address[] memory tos,
        uint256[] memory amounts
    ) internal {
        for (uint256 i = 0; i < tos.length; i++) {
            ERC20Mock underlyingToken = ERC20Mock(
                address(StrategyBase(strategyAddress).underlyingToken())
            );
            underlyingToken.mint(tos[i], amounts[i]);
        }
    }

    // TODO: this doesn't actually advance by n blocks... maybe because broadcasting batches txs somehow..?
    function advanceChainByNBlocks(uint256 n) public {
        for (uint256 i = 0; i < n; i++) {
            // we transfer eth to ourselves to advance the block
            vm.broadcast(msg.sender);
            payable(msg.sender).transfer(1 wei);
        }
    }

    function convertBoolToString(
        bool input
    ) public pure returns (string memory) {
        if (input) {
            return "true";
        } else {
            return "false";
        }
    }

    function convertOperatorStatusToString(
        IRegistryCoordinator.OperatorStatus operatorStatus
    ) public pure returns (string memory) {
        if (
            operatorStatus ==
            IRegistryCoordinator.OperatorStatus.NEVER_REGISTERED
        ) {
            return "NEVER_REGISTERED";
        } else if (
            operatorStatus == IRegistryCoordinator.OperatorStatus.REGISTERED
        ) {
            return "REGISTERED";
        } else if (
            operatorStatus == IRegistryCoordinator.OperatorStatus.DEREGISTERED
        ) {
            return "DEREGISTERED";
        } else {
            return "UNKNOWN";
        }
    }

    // Forge scripts best practice: https://book.getfoundry.sh/tutorials/best-practices#scripts
    function readInput(
        string memory inputFileName
    ) internal view returns (string memory) {
        string memory inputDir = string.concat(
            vm.projectRoot(),
            "/script/input/"
        );
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory file = string.concat(inputFileName, ".json");
        return vm.readFile(string.concat(inputDir, chainDir, file));
    }

    function readOutput(
        string memory outputFileName
    ) internal view returns (string memory) {
        string memory inputDir = string.concat(
            vm.projectRoot(),
            "/script/output/"
        );
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory file = string.concat(outputFileName, ".json");
        return vm.readFile(string.concat(inputDir, chainDir, file));
    }

    function writeOutput(
        string memory outputJson,
        string memory outputFileName
    ) internal {
        string memory outputDir = string.concat(
            vm.projectRoot(),
            "/script/output/"
        );
        string memory chainDir = string.concat(vm.toString(block.chainid), "/");
        string memory outputFilePath = string.concat(
            outputDir,
            chainDir,
            outputFileName,
            ".json"
        );
        vm.writeJson(outputJson, outputFilePath);
    }
}
