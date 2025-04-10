'use client';
import { useEffect } from 'react';

export default function RegisterPage() {
  useEffect(() => {
    const form = document.getElementById('register-form');
    const errorMsg = document.getElementById('error-message');

    form.addEventListener('submit', async (e) => {
      e.preventDefault();

      const userData = {
        nickname: form.nickname.value,
        email: form.email.value,
        password: form.password.value,
        first: form.first.value,
        last: form.last.value,
        date_of_birth: form.date_of_birth.value,
        about_me: form.about_me.value,
        status: 0, // default, not shown to user
      };

      try {
        const response = await fetch('http://localhost:8080/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(userData),
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
    <div className="register-container">
      <h1>Register</h1>
      <form id="register-form">
        <label htmlFor="nickname">Nickname:</label>
        <input type="text" name="nickname" required />

        <label htmlFor="email">Email:</label>
        <input type="email" name="email" required />

        <label htmlFor="password">Password:</label>
        <input type="password" name="password" required />

        <label htmlFor="first">First Name:</label>
        <input type="text" name="first" required />

        <label htmlFor="last">Last Name:</label>
        <input type="text" name="last" required />

        <label htmlFor="date_of_birth">Date of Birth:</label>
        <input type="date" name="date_of_birth" required />

        <label htmlFor="about_me">About Me:</label>
        <textarea name="about_me" rows="3"></textarea>

        <button type="submit">Register</button>

        <p id="error-message" style={{ color: 'red' }}></p>
      </form>

      <style jsx>{`
        .register-container {
          max-width: 500px;
          margin: 50px auto;
          padding: 25px;
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

        input,
        textarea {
          padding: 8px;
          margin-top: 5px;
          font-size: 16px;
        }

        button {
          margin-top: 20px;
          padding: 10px;
          font-size: 16px;
          background-color: #28a745;
          color: white;
          border: none;
          border-radius: 6px;
          cursor: pointer;
        }

        button:hover {
          background-color: #218838;
        }
      `}</style>
    </div>
  );
}
