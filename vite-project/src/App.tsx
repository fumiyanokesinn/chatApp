import { UIProvider } from "@yamada-ui/react";
import "./App.css";
import { Login } from "./components/Login";

function App() {
  return (
    <UIProvider>
      <Login />
    </UIProvider>
  );
}

export default App;
