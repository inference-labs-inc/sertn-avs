import { render } from "preact";
import App from "./pages/app.tsx";
import "./styles/index.css";
import { Router, Route } from "preact-router";
import { WagmiConfig } from "wagmi";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { config } from "./common/constants.ts";
const queryClient = new QueryClient();

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
