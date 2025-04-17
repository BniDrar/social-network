"use client"

import PostList from "@/components/PostList/PostList";
import styles from "./page.module.css";
import { useContext, useEffect } from "react";
import { Context } from "@/context/context.js";


export default function Home() {
  const { username } = useContext(Context);
  useEffect(() => {
    console.log("Username in context has been set:", username); // This will log the updated username
  }, [username]); // Track username changes

  return (
    <div>
      <div>Welcome {username}</div>
      <main className={styles.container}>
        <div className={styles.postsContainer}>
          <PostList />
        </div>
      </main>
    </div>
  );
}
