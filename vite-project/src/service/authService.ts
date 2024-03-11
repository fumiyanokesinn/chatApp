// authService.js

type LoginRequest = {
  email: string;
  password: string;
};

export async function login(request: LoginRequest) {
  const response = await fetch("http://localhost:8080/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(request),
  });
  if (!response.ok) {
    throw new Error("Login failed");
  }
  return await response.json();
}
