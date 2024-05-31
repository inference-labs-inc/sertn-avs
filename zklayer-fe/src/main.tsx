import { render } from "preact";
import App from "./pages/app.tsx";
import "./styles/index.css";
import { Router, Route } from "preact-router";
import { http, createConfig, WagmiConfig } from "wagmi";
import { localhost } from "wagmi/chains";
import { injected, metaMask, safe } from "wagmi/connectors";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
const queryClient = new QueryClient();

export const config = createConfig({
  chains: [localhost],
  connectors: [injected(), metaMask(), safe()],
  transports: {
    [localhost.id]: http(),
  },
});

render(
  <WagmiConfig config={config}>
    <QueryClientProvider client={queryClient}>
      <Router>
        <Route path="/" component={App} />
      </Router>
    </QueryClientProvider>
  </WagmiConfig>,
  document.getElementById("app")!
);
