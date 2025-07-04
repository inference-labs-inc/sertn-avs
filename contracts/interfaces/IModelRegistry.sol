// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

interface IModelRegistry {
    /**
     * @notice The strategy used to verify the model
     * XXX: changing this struct, don't forget to take care of `client/src/common/contract_constants.py`
     */
    enum VerificationStrategy {
        None,
        Onchain,
        Offchain
    }
    /**
     * @notice The event emitted when a new model is created
     * @param modelId The id of the model
     * @param modelVerifier The address of the model verifier
     * @param verificationStrategy The strategy used to verify the model
     * @param modelURI The URI of the model
     * @param computeCost The compute cost of the model
     */
    event ModelCreated(
        uint256 indexed modelId,
        address indexed modelVerifier,
        VerificationStrategy indexed verificationStrategy,
        string modelURI,
        uint256 computeCost
    );
    /**
     * @notice The event emitted when a model is updated
     * @param modelId The id of the model
     * @param modelURI The URI of the model
     */
    event ModelURIUpdated(uint256 indexed modelId, string indexed modelURI);
    /**
     * @notice The event emitted when the compute cost of a model is updated
     * @param modelId The id of the model
     * @param computeCost The new compute cost of the model
     */
    event ComputeCostUpdated(uint256 indexed modelId, uint256 indexed computeCost);

    /**
     * @notice The event emitted when a model verifier is updated
     * @param modelId The id of the model
     * @param modelVerifier The address of the model verifier
     */
    event ModelVerifierUpdated(uint256 indexed modelId, address indexed modelVerifier);

    /**
     * @notice The event emitted when a verification strategy is updated
     * @param modelId The id of the model
     * @param verificationStrategy The new verification strategy of the model
     */
    event VerificationStrategyUpdated(
        uint256 indexed modelId,
        VerificationStrategy indexed verificationStrategy
    );

    /**
     * @notice The error emitted when a model already exists
     * @param modelId The id of the model
     */
    error ModelAlreadyExists(uint256 modelId);
    /**
     * @notice The error emitted when a model verifier is invalid
     */
    error InvalidModelVerifier();
    /**
     * @notice The error emitted when a model does not exist
     */
    error ModelDoesNotExist();

    /**
     * @notice The function to create a new model
     * @param modelVerifier The address of the model verifier
     * @param verificationStrategy The strategy used to verify the model
     * @param modelURI The URI of the model
     */
    function createNewModel(
        address modelVerifier,
        VerificationStrategy verificationStrategy,
        string memory modelURI,
        uint256 computeCost
    ) external returns (uint256 modelId);

    /**
     * @notice The function to update the URI of a model
     * @param modelId The id of the model
     * @param modelURI The new URI of the model
     */
    function updateModelURI(uint256 modelId, string memory modelURI) external;

    /**
     * @notice The function to update the compute cost of a model
     * @param modelId The id of the model
     * @param computeCost The new compute cost of the model
     */
    function updateComputeCost(uint256 modelId, uint256 computeCost) external;

    /**
     * @notice The function to update the model verifier of a model
     * @param modelId The id of the model
     * @param modelVerifier The new address of the model verifier
     */
    function updateModelVerifier(uint256 modelId, address modelVerifier) external;

    /**
     * @notice The function to update the verification strategy of a model
     * @param modelId The id of the model
     * @param verificationStrategy The new verification strategy of the model
     */
    function updateVerificationStrategy(
        uint256 modelId,
        VerificationStrategy verificationStrategy
    ) external;
}
