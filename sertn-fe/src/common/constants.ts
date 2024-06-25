import { localhost } from "viem/chains";
import { createConfig, http } from "wagmi";
import { injected, metaMask, safe } from "wagmi/connectors";

export const backendUrl = "http://localhost:8080/inference";

export const taskManagerAddress = "0x28410d159554A95545571A1362c7F51c37D31021";
export const ModelDBAddress = "0x2Fe80512816f10F254E0d9593199cdc54d5B2456";

export const config = createConfig({
  chains: [localhost],
  connectors: [injected(), metaMask(), safe()],
  transports: {
    [localhost.id]: http(),
  },
});
