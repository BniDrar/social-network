"use client";

import styles from "./page.module.css";
import Navbar from "@/components/naveBar/nave";
import Post from "@/components/post/post";
import CreatePost from "@/components/createPost/createPost";
import Menu from "@/components/menu/menu";
import Contacts from "@/components/contacts/contacs";
import Profile from "@/components/profile/profile";
import { getPosts } from "@/services/posts";
import { useEffect, useState } from "react";
import CreateGroup from "@/components/createGroup/createGroup";
import Groups from "@/components/groups/groups";
import { usePathname } from "next/navigation";
import SowOnmoble from "@/components/showOnMobile/showOnmoble";


export default function Home() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const pathname = usePathname();

  const fetchPosts = async () => {
    if (loading || !hasMore) return;
    setLoading(true);

    const body = {};
    if (posts.length > 0) {
      const lastPost = posts[posts.length - 1];
      body.last_id = lastPost.id;
      body.creation_time = lastPost.created_at;
    }
    body.limit = 10;

    try {
      const fetchedPosts = await getPosts(
        body,
        `${process.env.BACKEND_URL}/api/posts`
      );

      if (!fetchedPosts || fetchedPosts.length === 0) {
        setHasMore(false);
      } else {
        setPosts((prev) => {
          const existingIds = new Set(prev.map((p) => p.id));
          const newUniquePosts = fetchedPosts.filter(
            (p) => !existingIds.has(p.id)
          );
          return [...prev, ...newUniquePosts];
        });
      }
    } catch (error) {
      console.error("Error fetching posts:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    setPosts([]);
    setHasMore(true);
    fetchPosts();
  }, [pathname]);

  useEffect(() => {
    let debounceTimer;

    const handleScroll = () => {
      clearTimeout(debounceTimer);
      debounceTimer = setTimeout(() => {
        if (
          window.innerHeight + window.scrollY >=
            document.body.offsetHeight - 100 &&
          !loading &&
          hasMore
        ) {
          fetchPosts();
        }
      }, 200);
    };

    window.addEventListener("scroll", handleScroll);
    return () => {
      clearTimeout(debounceTimer);
      window.removeEventListener("scroll", handleScroll);
    };
  }, [loading, hasMore, posts]);

  return (
    <div className={styles.page}>
      <Navbar />
      <main className={styles.main}>
        <SowOnmoble/>
        <div className={styles.leftSidebar} id="left-sidebar">
          <Profile />
          <Menu />
        </div>
        <div className={styles.container}>
          <CreatePost />
          <div className={styles.posts}>
            {posts && posts.length > 0 ? (
              posts.map((post) => <Post key={post.id} post={post} />)
            ) : (
              <div className={styles.noPosts}>No posts available</div>
            )}
            {loading && (
              <div className={styles.loading}>Loading more posts...</div>
            )}
            {!hasMore && (
              <div className={styles.noMorePosts}>No more posts to load</div>
            )}
          </div>
          <div className={styles.rightSidebar} id="right-sidebar">
            <Contacts />
            <Groups />
            <CreateGroup />
          </div>
        </div>
      </main>
    </div>
  );
}