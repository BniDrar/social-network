'use client';

import { useEffect } from 'react';

export default function LoginPage() {
  useEffect(() => {
    const form = document.getElementById('login-form');
    const errorMsg = document.getElementById('error-message');

    form.addEventListener('submit', async (e) => {
      e.preventDefault();

      const nickname = form.nickname.value;
      const password = form.password.value;

      console.log(nickname, password);
      try {
        const response = await fetch('http://localhost:8080/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ nickname, password }),
        });

        // Check if the response is OK (status 2xx)
        if (response.ok) {
          // If the response is successful, redirect to the homepage
          console.log('Login successful');
          window.location.href = '/';  // Redirect to homepage
        } else {
          // If the response is not OK, handle the error message
          const errorData = await response.text();  // Get the raw text (no JSON)
          throw new Error(errorData || 'Login failed');
        }
      } catch (err) {
        console.error(err.message);
        errorMsg.textContent = err.message;  // Display error message to the user
      }
    });
  }, []);

  return (
    <div className="login-container">
      <h1>Login</h1>
      <form id="login-form">
        <label htmlFor="nickname">nickname:</label>
        <input type="nickname" id="nickname" name="nickname" required />

        <label htmlFor="password">Password:</label>
        <input type="password" id="password" name="password" required />

        <button type="submit">Login</button>

        <p id="error-message" style={{ color: 'red' }}></p>
      </form>

      <style jsx>{`
        .login-container {
          max-width: 400px;
          margin: 100px auto;
          padding: 20px;
          border: 1px solid #ccc;
          border-radius: 12px;
          box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
          font-family: sans-serif;
        }

        form {
          display: flex;
          flex-direction: column;
        }

        label {
          margin-top: 10px;
        }

        input {
          padding: 8px;
          margin-top: 5px;
          font-size: 16px;
        }

        button {
          margin-top: 20px;
          padding: 10px;
          font-size: 16px;
          background-color: #0070f3;
          color: white;
          border: none;
          border-radius: 6px;
          cursor: pointer;
        }

        button:hover {
          background-color: #0059c1;
        }
      `}</style>
    </div>
  );
}
