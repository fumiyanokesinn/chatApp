// useLogin.js
import { useState } from "react";

type CreateAccountRequest = {
  name: string;
  email: string;
  password: string;
};

export default function useCreateAccount() {
  const [message, setMessage] = useState<string>("");

  const createAccount = async (request: CreateAccountRequest) => {
    setMessage("");
    const response = await fetch("http://localhost:8080/create_account", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    });
    const data = await response.json();

    setMessage(data.message);
  };

  return { createAccount, message };
}
