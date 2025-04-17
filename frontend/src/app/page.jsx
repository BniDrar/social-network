import PostList from "@/components/PostList/PostList";
import styles from "./page.module.css"
import Navebar from '@/components/Navebar/Navebar';
export default function Home() {
  return (
    <div>
      <Navebar />
      <main className={styles.container}>
        <div className={styles.postsContainer}>
          <PostList />
        </div>
      </main>
    </div>
  );
}
