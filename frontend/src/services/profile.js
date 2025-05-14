import { cookies } from "next/headers";

const getCookie = () => {
  const cookiestore = cookies();
  const cookie = cookiestore.toString();
  return cookie
}

const getUserInfo = async (user_id) => {
  try {
    const cookie = getCookie()
    const response = await fetch(
      `${process.env.BACKEND_URL}/api/user/profile?userid=${user_id}`,
      {
        headers: { "Content-Type": "application/json", cookie },
        cache: "no-store",
      }
    );

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json()
    return await data;
  } catch (err) {
    console.error("Error fetching user info:", err.message);
    return null;
  }
};


let lastId = 0
const getUserPosts = async (user_id, offset) => {
  try {
    const cookie = getCookie()
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/posts`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        cookie,
      },
      body: JSON.stringify({
        id: user_id,
        last_id: lastId == 0 ? offset : lastId
      }),
      cache: "no-store",
    })
    if (!response.ok) throw new Error(`response err ${response.status}`)
    const data = await response.json()
    data ? lastId = await data[data.length - 1].id : ""
    return await data
  } catch (err) {
    console.error(err)
  }
}

export { getUserInfo, getUserPosts };
