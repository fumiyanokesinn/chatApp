// useLogin.js
import { useState } from "react";

type LoginRequest = {
  email: string;
  password: string;
};

export default function useLogin() {
  const [message, setMessage] = useState<string>("");

  const login = async (request: LoginRequest) => {
    setMessage("");
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    });
    const data = await response.json();

    setMessage(data.message);
  };

  return { login, message };
}
