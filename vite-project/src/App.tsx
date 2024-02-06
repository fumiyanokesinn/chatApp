import { UIProvider } from "@yamada-ui/react";
import { Login } from "./components/Login";

function App() {
  return (
    <UIProvider>
      <Login />
    </UIProvider>
  );
}

export default App;
