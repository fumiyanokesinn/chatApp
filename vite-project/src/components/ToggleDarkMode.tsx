import { Switch } from "@yamada-ui/react";
import { useState, useEffect } from "react";
import { Icon as FontAwesomeIcon } from "@yamada-ui/fontawesome";
import { faMoon, faSun } from "@fortawesome/free-solid-svg-icons";

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
    <div>
      <FontAwesomeIcon icon={faSun} />

      <Switch
        colorScheme="green"
        defaultIsChecked={darkMode}
        onChange={() => setDarkMode(!darkMode)}
      />
      <FontAwesomeIcon icon={faMoon} />
    </div>
  );
};
