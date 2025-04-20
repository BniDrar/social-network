'use client'

import React, { useEffect, useState } from 'react'
import styles from './page.module.css'
"use client";

import styles from "./style.module.css";
import { useEffect} from "react";
import { useRouter } from "next/navigation";

export default function page() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  })

  const handleChange = (e) => {
    const { id, value } = e.target;
    console.log("id : ", id);
    console.log("value : ", value);
    
    
    setFormData(prevData => ({
      ...prevData,
      [id]: value
    }));
  }

  const handleSubmit = () => {
    useEffect(() => {
      console.log(formData);
      // const response = await fetch("/api/Auth", {
      // body: formData,
      // });
      // const JsonRes = await response.json();
      // if (!JsonRes.error) {
      // location.href = "/";
      // }
    })
  }

export default function LoginPage() {
  console.log('login page ')
  // const { username, setUsername } = useContext(Context);
  // console.log('we started with this username:', username )
  const router = useRouter();

  useEffect(() => {
    const form = document.getElementById("login-form");
    const errorMsg = document.getElementById("error-message");

    form.addEventListener("submit", async (e) => {
      e.preventDefault();

      const nickname = form.nickname.value;
      const password = form.password.value;

      console.warn(nickname, password);

      try {
        const response = await fetch("http://localhost:8080/api/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ nickname, password }),
          credentials: "include", // Correctly sends cookies
        });

        if (response.ok) {
          //setUsername(nickname); 
          //localStorage.setItem('username', nickname); // Persist to localStorage
          router.push("/"); // Use client-side navigation // navigat is for server side
        }
      } catch (err) {
        console.error(err.message);
        errorMsg.textContent = err.message; // Display error message to the user
      }
    });
  }, []);

  return (
    <div className={styles.container}>
      <h1 className={styles.bigTitle}>zellcode</h1>
      <form className={styles.form} onSubmit={handleSubmit}>
        <input className={styles.input} id='username' placeholder='enter email or username' required onChange={handleChange} />
        <input className={styles.input} type='password' placeholder='enter your password' required onChange={handleChange} />
        <button className={styles.button} type='submit'>Login</button>
    <div className={styles.loginContainer}>
      <h1>Login</h1>
      <form id="login-form" className={styles.loginForm}>
        <label htmlFor="nickname" className={styles.label}>
          nickname:
        </label>
        <input
          type="nickname"
          id="nickname"
          name="nickname"
          required
          className={styles.input}
        />

        <label htmlFor="password" className={styles.label}>
          Password:
        </label>
        <input
          type="password"
          id="password"
          name="password"
          required
          className={styles.input}
        />

        <button type="submit" className={styles.button}>
          Login
        </button>
        <p className={styles.linkText}>
          Don't have an account?{" "}
          <a href="/register" className={styles.link}>
            Register here
          </a>
        </p>
        <p id="error-message" style={{ color: "red" }}></p>
      </form>
    </div>
  );
}
