export async function GetNotifs() {
    try {
        const response = await fetch(
            `${process.env.BACKEND_URL}/api/user/notifications`,
            {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            }
        );
        if (!response.ok) {
            return null;
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching Contact:", error);
        throw error;
    }
}