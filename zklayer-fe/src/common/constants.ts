import { localhost } from "viem/chains";
import { createConfig, http } from "wagmi";
import { injected, metaMask, safe } from "wagmi/connectors";

export const backendUrl = "http://localhost:8080/inference";

export const taskManagerAddress = "0x9E545E3C0baAB3E08CdfD552C960A1050f373042";

export const config = createConfig({
  chains: [localhost],
  connectors: [injected(), metaMask(), safe()],
  transports: {
    [localhost.id]: http(),
  },
});
