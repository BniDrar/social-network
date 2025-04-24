// export const validbackendUrl = `http://localhost:8080`
import { backendUrl } from "@/utils/ustil";

export default async function GetLastChats() {
  try {
    const response = await fetch(`${backendUrl}/last-chats`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching :", error);
    throw error;
  }
}
