import { validbackendUrl } from "@/utils/ustil.js";
async function login(ReqData) {
  try {
    const response = await fetch(`${validbackendUrl}/api/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(ReqData),
      credentials: "include", // Correctly sends cookies
    });
    if (!response.ok) {
      let resp = await response.json();
      return {
        status: response.status,
        error: resp.error || "An error occurred",
      };
    }
    return response;
  } catch (error) {
    console.error("Error logging in:", error);
    return {
      status: 500,
      error: "Internal server error",
    };
  }
}

async function register(ReqData) {
  try {
    const response = await fetch(`${validbackendUrl}/api/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(ReqData),
    });
    if (!response.ok) {
      let resp = await response.json();
      return {
        status: response.status,
        error: resp.error || "An error occurred",
      };
    }
    return response;
  } catch (error) {
    console.error("Error registering:", error);
    return {
      status: 500,
      error: "Internal server error",
    };
  }
}

export { login, register };
