import { Button } from "@yamada-ui/react";
import { useState, useEffect } from "react";

export const ToggleDarkMode = () => {
  const [darkMode, setDarkMode] = useState(
    window.matchMedia &&
      window.matchMedia("(prefers-color-scheme: dark)").matches
  );

  useEffect(() => {
    if (darkMode) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [darkMode]);

  return (
    <Button onClick={() => setDarkMode(!darkMode)}>
      {darkMode ? "ダークモード" : "ライトモード"}
    </Button>
  );
};
