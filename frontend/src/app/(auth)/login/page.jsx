'use client';

import styles from "./style.module.css"

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
    <div className={styles.loginContainer}>
      <h1>Login</h1>
      <form id="login-form" className={styles.loginForm}>
        <label htmlFor="nickname" className={styles.label}>nickname:</label>
        <input type="nickname" id="nickname" name="nickname" required className={styles.input} />

        <label htmlFor="password" className={styles.label}>Password:</label>
        <input type="password" id="password" name="password" required className={styles.input} />

        <button type="submit" className={styles.button}>Login</button>

        <p id="error-message" style={{ color: 'red' }}></p>
      </form>

    </div>
  );
}
