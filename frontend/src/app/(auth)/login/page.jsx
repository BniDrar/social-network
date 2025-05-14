"use client";

import { useUser } from "@/context/userContext";
import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import styles from "./login.module.css";
import { login } from "@/services/auth";

export default function Login() {
  const { setIsLoggedIn } = useUser();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const loginData = {
      nickname: username,
      password: password,
    };
    let resp = await login(loginData);
    if (resp.ok) {
      setIsLoggedIn(true);
      router.push("/");
    } else {
      let errorForm = document.querySelector("#errorForm");
      errorForm.innerHTML = resp.error || "An error occurred";
      setInterval(() => {
        errorForm.innerHTML = "";
      }, 3000);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.formWrapper}>
        <h1 className={styles.title}>Welcome Back</h1>
        <p className={styles.subtitle}>Please enter your details to sign in</p>

        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.inputGroup}>
            <input
              type="text"
              placeholder="Email or Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className={styles.input}
              required
            />
          </div>

          <div className={styles.inputGroup}>
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className={styles.input}
              required
            />
          </div>

          <div className={styles.options}>
            <label className={styles.remember}>
              <input type="checkbox" disabled /> Remember me
            </label>
            <Link href="/forgot-password" className={styles.forgot}>
              Forgot password?
            </Link>
          </div>
          <div className="error" id="errorForm"></div>
          <button type="submit" className={styles.button}>
            Sign in
          </button>
        </form>

        <p className={styles.register}>
          Don't have an account?{" "}
          <Link href="/register" className={styles.link}>
            Sign up
          </Link>
        </p>
      </div>
    </div>
  );
}
