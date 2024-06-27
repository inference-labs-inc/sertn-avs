import { render } from "preact";
import App from "./pages/app.tsx";
import "./styles/index.css";
import { Router, Route } from "preact-router";
import { WagmiProvider } from "wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { config } from "./common/constants.ts";
const queryClient = new QueryClient();

render(
  <WagmiProvider config={config}>
    <QueryClientProvider client={queryClient}>
      <Router>
        <Route path="/" component={App} />
      </Router>
    </QueryClientProvider>
  </WagmiProvider>,
  document.getElementById("app")!
);
