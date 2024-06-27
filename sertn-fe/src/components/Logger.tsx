import { InferenceResponse, LogMessage, LogStyles } from "../common/types";

const Logger = ({
  logs,
  inference,
}: {
  logs: LogMessage[];
  inference: InferenceResponse;
}) => {
  return (
    <div
      class={`${
        inference.taskIndex >= 0 ? "h-60 p-6 mb-4" : "h-0"
      } text-white bg-slate-800 w-full rounded-sm transition-all duration-500 overflow-y-auto scrollbar-hide flex flex-col`}
    >
      {logs
        .filter(
          (log) => !!log.message.length && log.taskIndex === inference.taskIndex
        )
        .map((log, i) => (
          <div key={log.message + "" + i}>
            <span class={LogStyles[log.level]}>
              {log.level}: [Block Number - {log.blockNumber}] [TaskIndex -{" "}
              {log.taskIndex}]
            </span>
            {" " + log.message}
          </div>
        ))}
    </div>
  );
};
export default Logger;
