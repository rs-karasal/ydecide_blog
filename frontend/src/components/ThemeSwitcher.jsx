import { useEffect, useState } from "react";
import { SunOutlined, MoonOutlined } from "@ant-design/icons";

function ThemeSwitcher() {
  const [theme, setTheme] = useState("light");

  useEffect(() => {
    const savedTheme = localStorage.getItem("theme");
    if (savedTheme) {
      setTheme(savedTheme);
      document.documentElement.classList.add(savedTheme);
    } else {
      const prefersDarkMode = window.matchMedia(
        "(prefers-color-scheme: dark)"
      ).matches;
      const defaultTheme = prefersDarkMode ? "dark" : "light";
      setTheme(defaultTheme);
      document.documentElement.classList.add(defaultTheme);
    }
  }, []);



  const toggleTheme = () => {
    const newTheme = theme === "light" ? "dark" : "light";
    setTheme(newTheme);

    document.documentElement.classList.remove(theme);
    document.documentElement.classList.add(newTheme);

    localStorage.setItem("theme", newTheme);
  };

  return (
    <button
      className="p-2 bg-gray-200 dark:bg-gray-800 text-black dark:text-white"
      onClick={toggleTheme}
    >
      {theme === "light" ? <SunOutlined /> : <MoonOutlined />}
    </button>
  );
}

export default ThemeSwitcher;
