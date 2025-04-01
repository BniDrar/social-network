import PostList from "@/components/PostList/PostList";
import styles from "./page.module.css"

export default function Home() {
  return (
    <div>
      <main className={styles.container}>
        <div className={styles.postsContainer}>
          <PostList />
        </div>
      </main>
    </div>
  );
}
