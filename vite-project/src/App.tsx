import { UIProvider } from "@yamada-ui/react";
import { Login } from "./pages/Login";
import { Header } from "./components/Header";
import { Routes, Route } from "react-router-dom";
import { UserStore } from "./pages/UserStore";

function App() {
  return (
    <UIProvider>
      <Header title="ログインページ" />
      <Routes>
        <Route path={`/`} element={<Login />} />
        <Route path={`/page1`} element={<UserStore />} />
      </Routes>
    </UIProvider>
  );
}

export default App;
