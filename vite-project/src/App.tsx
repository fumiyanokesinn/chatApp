import { UIProvider } from "@yamada-ui/react";
import { Login } from "./pages/Login";
import { Header } from "./components/Header";
import { Routes, Route } from "react-router-dom";
import { CreateAccount } from "./pages/CreateAccount";

function App() {
  return (
    <UIProvider>
      <Header title="ログインページ" />
      <Routes>
        <Route path={`/`} element={<Login />} />
        <Route path={`/create-account`} element={<CreateAccount />} />
      </Routes>
    </UIProvider>
  );
}

export default App;
