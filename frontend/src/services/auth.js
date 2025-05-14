 async function login(credentials){
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
      credentials: "include",
    });

    if (!response.ok) {
      const resp = await response.json(); // Assign resp here
      return {
        status: response.status,
        error: resp.error || "An error occurred",
      };
    } else {
      return response;
    }
  } catch (error) {
    console.error("Login failed:", error);
    return { status: 500, error: "Failed to connect to the server" };
  }
};

async function getProfile(id = 0) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/profile?userid=${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    if (!response.ok) {
      return {
        status: response.status,
        error: "An error occurred",
      };
    }
  const data = await response.json();
  return data;
  } catch (error) {
    console.error("Error fetching profile:", error);
    return {
      status: 500,
      error: "Internal server error",
    };
  }
}


async function register(ReqData) {
  try {
    const formData = new FormData();
    for (const key in ReqData) {
      formData.append(key, ReqData[key]);
    }

    const response = await fetch(`${process.env.BACKEND_URL}/api/user/register`, {
      method: "POST",
      body: formData,
    });

    if (!response.ok) {
      const resp = await response.json();
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

async function IsAuth() {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/authenticate`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    if (!response.ok) {
      return {
        status: response.status,
        error: "An error occurred",
      };
    }
    return response;
  } catch (error) {
    console.error("Error checking authentication:", error);
    return {
      status: 500,
      error: "Internal server error",
    };
  }
}

async function logout() {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/logout`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    if (!response.ok) {
      return {
        status: response.status,
        error: "An error occurred",
      };
    }
    return response;
  } catch (error) {
    console.error("Error logging out:", error);
    return {
      status: 500,
      error: "Internal server error",
    };
  }
}
async function checkLogin() {
  try {
    const res = await fetch(`${process.env.BACKEND_URL}/ping/user`, {
      credentials: "include",
    });
    return res.ok;
  } catch (err) {
    console.error("Login check failed:", err);
    return false;
  }
}
export {getProfile, login, register, IsAuth, logout, checkLogin };