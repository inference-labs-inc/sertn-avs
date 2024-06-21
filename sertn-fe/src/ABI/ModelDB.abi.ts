export const ModelDBAbi = [
  {
    type: "function",
    name: "createNewModel",
    inputs: [
      {
        name: "modelVerifierAddress",
        type: "address",
        internalType: "address",
      },
      { name: "title", type: "string", internalType: "string" },
      { name: "description", type: "string", internalType: "string" },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "currentModelId",
    inputs: [],
    outputs: [{ name: "", type: "uint256", internalType: "uint256" }],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "modelAddresses",
    inputs: [{ name: "", type: "uint256", internalType: "uint256" }],
    outputs: [{ name: "", type: "address", internalType: "address" }],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "modelVerifiers",
    inputs: [{ name: "", type: "address", internalType: "address" }],
    outputs: [
      { name: "title", type: "string", internalType: "string" },
      { name: "description", type: "string", internalType: "string" },
      { name: "modelVerifier", type: "address", internalType: "address" },
    ],
    stateMutability: "view",
  },
];
