import { useCallback, useEffect, useMemo, useState } from "preact/hooks";
import { config, taskManagerAddress } from "../common/constants";
import { useAccount, useConnect, usePublicClient } from "wagmi";
import { formatAddress } from "../common/eth";
import { sertn } from "../ABI/TaskManager.abi";
import {
  InferenceForm,
  InferenceResponse,
  Log,
  LogLevel,
  LogMessage,
  ProofTypes,
} from "../common/types";
import { fetchInference, handleEvents } from "../common/utils";
import Logger from "../components/Logger";

let blockNumber = 0n;

const App = () => {
  const [logs, setLogs] = useState([
    {
      level: "ERROR",
      message: "",
      blockNumber: "0",
      taskIndex: 0,
    } as LogMessage,
  ]);

  const [inference, setInference] = useState({
    inputs: "",
    taskIndex: -1,
  } as InferenceResponse);

  const [proofType, setProofType] = useState(0);
  const { connectors, connect } = useConnect();
  const { address, isConnected } = useAccount();

  const client = usePublicClient({ config: config });

  const logger = useMemo(
    () => ({
      log: useCallback(
        (
          level: LogLevel,
          message: string,
          index: number,
          blockNumber: string
        ) => {
          setLogs((logs) =>
            [...logs, { level, message, blockNumber, taskIndex: index }]
              .sort((a, b) => +a.blockNumber - +b.blockNumber)
              .filter(
                (log, i, logs) => i == 0 || log.message !== logs[i - 1].message
              )
          );
        },
        [logs]
      ),
    }),
    []
  );

  useEffect(() => {
    if (!client) return;
    //Implementing the setInterval method
    const interval = setInterval(() => {
      (async () => {
        const logs = await client.getLogs({
          address: taskManagerAddress,
          events: sertn.abi.filter((element) => element.type === "event"),
          fromBlock: blockNumber,
        });
        if (logs.length) {
          blockNumber = logs[logs.length - 1].blockNumber + 1n;
          handleEvents(logs as unknown as Log[], logger);
        }
      })();
    }, 2000);
    return () => clearInterval(interval);
  }, [client]);

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
              zk<span class="text-slate-500">Layer</span>
            </h1>
            <form
              class="w-full transition-all mb-24"
              onSubmit={(e) => {
                e.preventDefault();
                (async () => {
                  const { elements: inferenceForm } =
                    e.target as unknown as InferenceForm;
                  setInference(
                    await fetchInference(
                      inferenceForm.input.value,
                      inferenceForm.type.value === "Pre Prove"
                    )
                  );
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
                  name="type"
                  value={ProofTypes[proofType]}
                >
                  {ProofTypes.map((proofType) => (
                    <option key={proofType} value={proofType}>
                      <span class="text-xs">{proofType}</span>
                    </option>
                  ))}
                </select>
              </div>
              <Logger logs={logs} inference={inference} />
              <div class="flex justify-center gap-4 w-full">
                <button
                  class="shadow bg-slate-200 hover:bg-slate-400 focus:shadow-outline focus:outline-none hover:text-white py-2 px-4 rounded"
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
