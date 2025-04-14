'use client'

import React, { useEffect, useState } from 'react'
import styles from './page.module.css'

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

  return (
    <div className={styles.container}>
      <h1 className={styles.bigTitle}>zellcode</h1>
      <form className={styles.form} onSubmit={handleSubmit}>
        <input className={styles.input} id='username' placeholder='enter email or username' required onChange={handleChange} />
        <input className={styles.input} type='password' placeholder='enter your password' required onChange={handleChange} />
        <button className={styles.button} type='submit'>Login</button>
      </form>
    </div>
  )
}
