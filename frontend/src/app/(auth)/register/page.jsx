'use client';
import { useEffect } from 'react';
import styles from './style.module.css';

export default function RegisterPage() {
  useEffect(() => {
    const form = document.getElementById('register-form');
    const errorMsg = document.getElementById('error-message');

    // Set max date for date_of_birth input to today
    const today = new Date().toISOString().split('T')[0];
    const dobInput = form.date_of_birth;
    if (dobInput) {
      dobInput.max = today;

    }
    form.addEventListener('submit', async (e) => {
      e.preventDefault();

      if (form.password.value !== form.confirm_password.value) {
        errorMsg.textContent = 'Passwords do not match';
        return;
      }

      if (form.password.value !== form.confirm_password.value) {
        errorMsg.textContent = 'Passwords do not match';
        return;
      }

      const userData = {
        nickname: form.nickname.value,
        email: form.email.value,
        password: form.password.value,
        first: form.first.value,
        last: form.last.value,
        date_of_birth: form.date_of_birth.value,
        about_me: form.about_me.value,
        status: 0,
      };

      try {
        const response = await fetch('http://localhost:8080/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(userData),
        });

        if (response.ok) {
          console.log('Registration successful');
          window.location.href = '/';
        } else {
          const errorData = await response.text();
          throw new Error(errorData || 'Registration failed');
        }
      } catch (err) {
        console.error(err.message);
        errorMsg.textContent = err.message;
      }
    });
  }, []);

  return (
    <div className={styles.container}>
      <h1>Register</h1>
      <form id="register-form" className={styles.form}>
        <label htmlFor="nickname" className={styles.label}>Nickname:</label>
        <input type="text" name="nickname" required className={styles.input} />

        <label htmlFor="email" className={styles.label}>Email:</label>
        <input type="email" name="email" required className={styles.input} />

        <label htmlFor="password" className={styles.label}>Password:</label>
        <input type="password" name="password" required className={styles.input} />

        <label htmlFor="confirm_password" className={styles.label}>Confirm Password:</label>
        <input type="password" name="confirm_password" required className={styles.input} />

        <label htmlFor="first" className={styles.label}>First Name:</label>
        <input type="text" name="first" required className={styles.input} />

        <label htmlFor="last" className={styles.label}>Last Name:</label>
        <input type="text" name="last" required className={styles.input} />

        <label htmlFor="date_of_birth" className={styles.label}>Date of Birth:</label>
        <input type="date" name="date_of_birth" required className={styles.input} max={new Date().toISOString().split('T')[0]} />

        <label htmlFor="about_me" className={styles.label}>About Me:</label>
        <textarea name="about_me" rows="3" className={styles.textarea}></textarea>

        <button type="submit" className={styles.button}>Register</button>

        <p className={styles.linkText}>
          Already have an account?{' '}
          <a href="/login" className={styles.link}>Login here</a>
        </p>

        <p id="error-message" className={styles.errorMessage}></p>
      </form>
    </div>
  );
}
