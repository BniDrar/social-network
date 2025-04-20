"use client";

import PostList from "@/components/PostList/PostList";
import styles from "./page.module.css";
import Navebar from "@/components/Navebar/Navebar";

import {LeftSideBar, RightSideBar} from "@/components/SideBar/SideBar";

export default function Home() {
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
