import { UIProvider } from "@yamada-ui/react";
import { Login } from "./components/Login";

function App() {
  return (
    <UIProvider>
      <div>
        <Login />
      </div>
    </UIProvider>
  );
}

export default App;
