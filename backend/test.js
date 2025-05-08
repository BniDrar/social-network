
// Function to login a user using username and password
function loginUser(username, password) {
    // Define the login URL for the API endpoint
    const url = 'https://localhost:8080/api/login';

    // Prepare the user credentials as JSON
    const userCredentials = {
        email: username,
        password: password
    };

    // Make the POST request with the user credentials in the body
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json', // Set content type as JSON
        },
        body: JSON.stringify(userCredentials) // Send the user credentials as JSON in the body
    })
    .then(response => {
        if (!response.ok) {
            // Handle non-OK responses (e.g., invalid credentials)
            return response.text().then(errorMessage => {
                throw new Error(errorMessage);
            });
        }
        return response.text(); // Return the response text if login is successful
    })
    .then(data => {
        console.log("Login successful:", data); // Handle success (maybe a success message)
    })
    .catch(error => {
        console.error("Error during login:", error.message); // Handle any errors
    });
}

// Example usage: login with username and password
const username = 'user1@example.com';  // Replace with the actual username
const password = '$2a$10$lWWMbAFxfFFVv5XiMIcRjuu/ovU8burJnM6qSTaFfBfDk4.T3QsRC';  // Replace with the actual password

loginUser(username, password);
fetch("localhost:8080/api/login", )