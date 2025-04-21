"use client";

import PostList from "@/components/PostList/PostList";
import styles from "./page.module.css";

import {LeftSideBar, RightSideBar} from "@/components/SideBar/SideBar";

export default function Home() {
  console.log('home page');
  return (
    <>
      <div className={styles.home}>
      <LeftSideBar />
       <main className={styles.container}>
          <div className={styles.postsContainer}>
            <PostList />
          </div>
        </main>
      <RightSideBar/>
      </div>
    </>
  );
}
