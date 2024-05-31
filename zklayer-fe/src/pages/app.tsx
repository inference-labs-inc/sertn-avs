import { useEffect, useRef, useState } from "preact/hooks";
import { backendUrl } from "../common/constants";
import { useAccount, useConnect } from "wagmi";
import { connectWallet, formatAddress } from "../common/eth";

type InferenceResponse = {
  inputs: string;
  taskIndex: number;
};
const ProofTypes = ["Challenge", "Pre Prove"] as const;

type InferenceForm = {
  elements: {
    type: {
      value: string;
    };
    input: {
      value: string;
    };
  };
};

const Logger = (setLogs: any) => {};

import { useWatchContractEvent } from "wagmi";
import { zkLayer } from "../ABI/TaskManager.abi";

const fetchInference = async (inputs: string) => {
  const response = await fetch(backendUrl + `?inputs=${inputs}`);
  const text = await response.text();
  console.log({ text });
  return JSON.parse(text) as InferenceResponse;
};

const App = () => {
  const [inference, setInference] = useState({
    inputs: "",
    taskIndex: -1,
  } as InferenceResponse);

  const [_signer, setSigner] = useState(null);

  const [proofType, setProofType] = useState(0);

  useEffect(() => {
    console.log({ inference });
  }, [inference.taskIndex]);

  const { connectors, connect } = useConnect();

  const { address, isConnected } = useAccount();
  const logsRef = useRef();

  useEffect(() => {
    if (!isConnected) return;
    connectWallet(setSigner);
  }, [isConnected, setSigner]);

  // useWatchContractEvent({
  //   address: "0x9E545E3C0baAB3E08CdfD552C960A1050f373042",
  //   abi: zkLayer.abi,
  //   eventName: "*",
  //   batch: false,
  //   onLogs(logs) {
  //     console.log("New logs!", logs[0]);
  //   },
  // });

  return (
    <>
      {/*credit @ https://stackoverflow.com/questions/68138453/how-create-a-true-sticky-header-footer-using-tailwindcss-sticks-to-bottom-even for the useful sticky header*/}
      <div class="flex flex-col min-h-screen transition-all ease-in-out duration-300">
        <header class="sticky z-50 top-0 p-5 flex justify-between">
          <a href="https://inferencelabs.com">About</a>
          <a
            href="#"
            onClick={() => {
              connect({ connector: connectors[0] });
            }}
          >
            {isConnected ? formatAddress(address as string) : "Connect Wallet"}
          </a>
        </header>
        <div class="flex-grow justify-center items-center flex">
          <main class="font-thin w-full max-w-xl px-5 -translate-y-10">
            <h1 class="text-8xl w-full text-center p-10 pt-0">
              ZK<span class="text-slate-500">Layer</span>
            </h1>
            <form
              class="w-full transition-all"
              onSubmit={(e) => {
                e.preventDefault();
                (async () => {
                  const { elements: inferenceForm } =
                    e.target as unknown as InferenceForm;
                  setInference(await fetchInference(inferenceForm.input.value));
                })();
              }}
            >
              <div class="flex tems-center mb-6">
                <input
                  class="bg-gray-100 border border-r-0 border-gray-400 rounded-sm rounded-r-none py-2 px-4 text-gray-700 leading-tight focus:outline-none  focus:border-slate-500 flex-grow w-3/5"
                  id="inline-full-name"
                  type="text"
                  name="input"
                  placeholder="Inputs"
                />
                <select
                  disabled={true}
                  class="appearance-none rounded-l-none bg-white border border-gray-400 hover:border-gray-500  p-2 rounded-sm leading-tight focus:outline-none w-30 h-full"
                  value={ProofTypes[proofType]}
                >
                  {ProofTypes.map((proofType) => {
                    return (
                      <option key={proofType} value={proofType}>
                        <span class="text-xs">{proofType}</span>
                      </option>
                    );
                  })}
                </select>
              </div>
              <div
                class={`${
                  inference.taskIndex >= 0 ? "h-60" : "h-0"
                } text-white bg-slate-600 w-full rounded-sm mb-6 transition-all duration-500 overflow-y-auto scrollbar-hide`}
              ></div>
              <div class="flex justify-center gap-4 w-full">
                <button
                  class="shadow bg-slate-200 hover:bg-slate-400 focus:shadow-outline focus:outline-none hover:text-white  py-2 px-4 rounded"
                  type="submit"
                >
                  Submit Inference
                </button>

                <button
                  class="shadow bg-slate-200 hover:bg-slate-400 focus:shadow-outline focus:outline-none hover:text-white  py-2 px-4 rounded"
                  type="button"
                  onClick={() => {
                    setProofType((proofType) => (proofType + 1) % 2);
                  }}
                >
                  Proof type
                </button>
              </div>
            </form>
          </main>
        </div>
        <footer class="sticky z-50 bottom-0 p-5">
          Made with 🥩 in Hamilton, Ontario
        </footer>
      </div>
    </>
  );
};

export default App;
