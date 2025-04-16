'use client'

import { Suspense, useEffect, useState } from "react"
import styles from './page.module.css'

import Post from "@/components/Post/Post";

// import { backendUrl } from "@/utils/ustil";

export default function Page(props) {
    const [post, setPost] = useState({
        "id": 1,
        "avatar": "/assets/images/no-face.jpg",
        "username": "yrahhaou",
        "content": "🍲🍖 Classic Homestyle Meatloaf with a Tangy Glaze 🍋✨",
        "image": "/assets/images/post1.jpg",
        "likes": "15k",
        "comments": "5k",
        "user_like": false,
        "created_at": "7h"
    })

    useEffect(() => async () => {
        // fetch Post and use State to change default post data
    })
    return (
        <>
            <Suspense fallback={<Loading />}>
                <div>
                    <Post key={1} data={post} />
                    {/* add Comments Here */}
                    <Comment comment={{ image: "https://randomuser.me/api/portraits/men/8.jpg", user: "mostapha Benada", content: "Okay but why does this look like it came straight out of a dream? 😍🔥”", likes: 41 }} />
                    <Comment comment={{ image: "https://randomuser.me/api/portraits/men/1.jpg", user: "mohammed Malawi", content: "BRB drooling. Can I get this delivered to my soul, please?", likes: 13 }} />
                    <Comment comment={{ image: "https://randomuser.me/api/portraits/men/13.jpg", user: "ayoub Beghdadi", content: "I need the recipe, the chef’s name, and a map to this exact location. Now.", likes: 21 }} />
                    <Comment comment={{ image: "https://randomuser.me/api/portraits/men/7.jpg", user: "saad Sfakssi", content: "This is what happiness looks like in edible form. I’m obsessed!", likes: 101 }} />
                </div>
            </Suspense>
        </>);
}

function Loading() {
    return (
        <h1>Loading...</h1>
    )
}

function Comment({ comment }) {
    return (
        <div className={styles.comment}>
            <div className={styles.commentHeader}>
                <img src={comment.image} className={styles.commentAvatar} />
                <div>
                    <div className={styles.commentUser}>{comment.user}</div>
                    <div className={styles.commentTime}>2 hours ago</div>
                </div>
            </div>

            <div className={styles.commentContent}>
                <p>{comment.content}</p>
            </div>

            <div className={styles.commentActions}>
                <div className={styles.commentAction}>
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                    </svg>
                    <span>{comment.likes}</span>
                </div>
                <div className={styles.commentAction}>
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                    </svg>
                    <span>Reply</span>
                </div>
            </div>
        </div>
    )
}