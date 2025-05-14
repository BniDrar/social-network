'use client';

import React, { useEffect, useState } from 'react';
import styles from './groupPostsTab.module.css';
import { GetPostsByGroupId } from '@/services/posts';
import Post from '@/components/post/post';
import CreatePost from '@/components/createPost/createPost';

export default function GroupPostsTab({ id }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);

  const fetchPosts = async () => {
    if (loading || !hasMore) return;
    setLoading(true);

    const body = { id: parseInt(id), limit: 10 };
    if (posts.length > 0) {
      const lastPost = posts[posts.length - 1];
      body.last_id = lastPost.id;
      body.creation_time = lastPost.created_at;
    }

    try {
      const fetchedPosts = await GetPostsByGroupId(body);
      if (!fetchedPosts || fetchedPosts.length === 0) {
        setHasMore(false);
      } else {
        // Remove duplicates by ID
        setPosts((prev) => {
          const existingIds = new Set(prev.map((p) => p.id));
          const newPosts = fetchedPosts.filter((p) => !existingIds.has(p.id));
          return [...prev, ...newPosts];
        });
      }
    } catch (error) {
      console.error("Error fetching group posts:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    setPosts([]);
    setHasMore(true);
    fetchPosts();
  }, [id]);

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

    window.addEventListener('scroll', handleScroll);
    return () => {
      clearTimeout(debounceTimer);
      window.removeEventListener('scroll', handleScroll);
    };
  }, [loading, hasMore, posts]);

  return (
    <div className={styles.groupPostsTab}>
      <CreatePost groupId={id}/>
      <div className={styles.posts}>
        {posts.length > 0 ? (
          posts.map((post) => <Post key={post.id} post={post} />)
        ) : (
          <div className={styles.noPosts}>No posts available</div>
        )}
        {loading && <div className={styles.loading}>Loading more posts...</div>}
        {!hasMore && <div className={styles.noMorePosts}>No more posts to load</div>}
      </div>
    </div>
  );
}
