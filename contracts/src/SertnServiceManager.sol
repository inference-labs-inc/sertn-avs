// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./ISertnTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from Sertn.
 * @author Layr Labs, Inc.
 */
contract SertnServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    ISertnTaskManager public immutable sertnTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlySertnTaskManager() {
        require(
            msg.sender == address(sertnTaskManager),
            "onlySertnTaskManager: not from sertn task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        address _sertnTaskManager
    )
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)),
            _registryCoordinator,
            _stakeRegistry
        )
    {
        sertnTaskManager = ISertnTaskManager(_sertnTaskManager);
    }

    function initialize(address _sertnOwner) public initializer {
        __Ownable_init();
        super._transferOwnership(_sertnOwner);
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlySertnTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
