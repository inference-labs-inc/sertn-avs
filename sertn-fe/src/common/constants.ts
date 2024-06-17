import { localhost } from "viem/chains";
import { createConfig, http } from "wagmi";
import { injected, metaMask, safe } from "wagmi/connectors";

export const backendUrl = "http://localhost:8080/inference";

export const taskManagerAddress = "0x8ac5eE52F70AE01dB914bE459D8B3d50126fd6aE";

export const config = createConfig({
  chains: [localhost],
  connectors: [injected(), metaMask(), safe()],
  transports: {
    [localhost.id]: http(),
  },
});
