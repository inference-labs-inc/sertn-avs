import { localhost } from "viem/chains";
import { createConfig, http } from "wagmi";
import { injected, metaMask, safe } from "wagmi/connectors";

export const backendUrl = "http://localhost:8080/inference";

export const taskManagerAddress = "0xa62835D1A6bf5f521C4e2746E1F51c923b8f3483";

export const config = createConfig({
  chains: [localhost],
  connectors: [injected(), metaMask(), safe()],
  transports: {
    [localhost.id]: http(),
  },
});
