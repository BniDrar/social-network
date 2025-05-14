export async function GetContacts(contacts) {
  try {
    const response = await fetch(
      `${process.env.BACKEND_URL}/api/contacts`, // 
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        //   body: JSON.stringify(contacts),
        credentials: "include",
      }
    );
    const data = await response.json();
    return data;
    // return response
  } catch (error) {
    console.error("Error fetching Contact:", error);
    throw error;
  }
}
