'use client';

import { use } from 'react';
import styles from "./page.module.css";
import Navbar from "@/components/naveBar/nave";
import Post from "@/components/post/post";
import Menu from "@/components/menu/menu";
import Contacts from "@/components/contacts/contacs";
import Profile from "@/components/profile/profile";
import Comment from "@/components/comment/comment";
import CreateComment from "@/components/createComment/createComment";
import { getPostById } from "@/services/posts";
import { useEffect, useState } from "react";
import { getComments } from '@/services/comment';
import CreateGroup from "@/components/createGroup/createGroup";
import Groups from "@/components/groups/groups";
import SowOnmoble from "@/components/showOnMobile/showOnmoble";

export default function PostPage({ params }) {
  const { id: postId } = use(params)
  const [post, setPost] = useState(null)
  const [comments, setComments] = useState([]);
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function fetchPost() {
      try {
        const fetchedPost = await getPostById(postId);
        setPost(await fetchedPost);
      } catch (error) {
        console.error("Failed to fetch post:", error);
      } finally {
        setLoading(false);
      }
    }
    async function fetchComments() {
      try {
        const fetchedComments = await getComments(postId);
        setComments(fetchedComments);
        setLoading(false);
      } catch (error) {
        console.error("Failed to fetch comments:", error);
      }
    }
    fetchPost();
    fetchComments();
  }, [postId]);

  const handleCommentCreated = (newComment) => {
    setComments((prevComments) => [newComment, ...(prevComments || [])]);
    setPost((prevPost) => ({
      ...prevPost,
      comments: prevPost.comments + 1,
    }));
  };

  return (
    <div className={styles.page}>
      <Navbar />
      <main className={styles.main}>
        <SowOnmoble />
        <div className={styles.leftSidebar} id="left-sidebar">
          <Profile />
          <Menu />
        </div>
        <div className={styles.container}>
          <div className={styles.posts}>
            {loading && <p>Loading post...</p>}
            {!loading && post && <Post post={post} />}
            {!loading && !post && <p>Post not found.</p>}
            <div className={styles.comments}>
              <CreateComment postId={postId} setComments={handleCommentCreated} comments={comments} />
              {comments?.length === 0 && <p>No comments yet.</p>}
              {comments?.length > 0 && <h3>Comments</h3>}
              {comments?.map((comment, index) => (
                <Comment key={`${comment.id}-${index}`} comment={comment} />
              ))}
            </div>
          </div>
        </div>
        <div className={styles.rightSidebar} id='right-sidebar'>
          <Contacts />
          <Groups />
          <CreateGroup />
        </div>
      </main>
    </div>
  );
}