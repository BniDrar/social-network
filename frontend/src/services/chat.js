export async function postMessages(cursor) {
  try {
    const res = await fetch(`${process.env.BACKEND_URL}/api/messages`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(cursor),
      credentials: "include",
    });

    if (!res.ok) throw new Error(`HTTP ${res.status}`);

    const data = await res.json();
    return data?.reverse();
  } catch (err) {
    console.error("Fetch error:", err);
  }
}
