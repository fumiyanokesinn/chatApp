import { Switch } from "@yamada-ui/react";
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
    <Switch
      colorScheme="green"
      defaultIsChecked={darkMode}
      onChange={() => setDarkMode(!darkMode)}
    />
  );
};
