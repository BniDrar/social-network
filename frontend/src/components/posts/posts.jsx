"use client"
import { useEffect, useState } from "react";
import styles from "./posts.module.css"
import Post from "../post/post";
import { getPosts } from "@/services/posts";

const Posts = ({ body, lien }) => {
    const [posts, setPosts] = useState([]);
    const [last_id, setLastId] = useState(null);
    const [loading, setLoading] = useState(false);
    const [hasMore, setHasMore] = useState(true);

    const fetchPosts = async () => {
        if (loading || !hasMore) return;
        setLoading(true);
        if (posts.length > 0) {
            const lastPost = posts[posts.length - 1];
            body.last_id = lastPost.id;
            body.creation_time = lastPost.created_at;
        }
        body.limit = 10
        try {
            const fetchedPosts = await getPosts(body, lien);

            if (!fetchedPosts || fetchedPosts.length === 0) {
                setHasMore(false);
            } else {
                setPosts((prev) => [...prev, ...fetchedPosts]);
            }
        } catch (error) {
            console.error("Error fetching posts:", error);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchPosts();
    }, []);

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
        <>
            {posts && posts.length > 0 ? posts.map((post, index) => {
                return <Post key={`${post.id}-${index}`} post={post} />
            }) : <div className={styles.noPosts} style={{ textAlign: "center" }}>no posts!</div>}
        </>
    );
}

export default Posts;