import { UIProvider } from "@yamada-ui/react";
import { Login } from "./components/Login";
import { Header } from "./components/Header";

function App() {
  return (
    <UIProvider>
      <Header title="ログインページ" />
      <Login />
    </UIProvider>
  );
}

export default App;
