
export const getUserInfoClient = async (user_id) => {
    try {
        const response = await fetch(
            `${process.env.BACKEND_URL}/api/user/profile?userid=${user_id}`, // Use relative URL
            {
                credentials: "include", // let browser send cookies
                cache: "no-store",
            }
        );

        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return await response.json();
    } catch (err) {
        console.error("Client error fetching user info:", err.message);
        return null;
    }
};
