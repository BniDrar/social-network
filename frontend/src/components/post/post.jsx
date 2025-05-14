'use client';

import { useState } from 'react';
import styles from "./post.module.css";
import Image from "next/image";
import { useRouter, usePathname } from 'next/navigation';
import { BiHeart, BiComment, BiShareAlt } from "react-icons/bi";
import { formatDate } from "@/utils/formateDate";
import isValidUrl from "@/utils/valideUrl";
import { likePost } from '@/services/posts';

export default function Post({ post }) {
  const router = useRouter();
  const pathname = usePathname();
  const goToProfile = () => {
    !pathname.startsWith("/profile") ? router.push(`/profile/${post.user_id}`) : null;
  };

  const avatarSrc = post.avatar
    ? `${process.env.MEDIA_URL}${post.avatar}`
    : "/default-avatar.jpeg";

  const imageSrc = post.image
    ? `${process.env.MEDIA_URL}${post.image}`
    : "";

  const username = post.username || `${post.first_name} ${post.last_name}`

  const [likes, setLikes] = useState(post.likes_count || 0);
  const [liked, setLiked] = useState(post.engagement === 1);
  const [isLiking, setIsLiking] = useState(false);

  const handleLike = async () => {
    if (isLiking) return;

    const originalLikes = likes;
    const newLikedState = !liked;
    const newLikesCount = newLikedState ? likes + 1 : likes - 1;

    setIsLiking(true);
    setLiked(newLikedState);
    setLikes(newLikesCount);

    try {
      await likePost(post.id);
    } catch (error) {
      setLiked(!newLikedState);
      setLikes(originalLikes);
      console.error("Failed to like/unlike post:", error);
    } finally {
      setIsLiking(false);
    }
  };

  return (
    <div className={styles.post}>
      <div className={styles.postHeader}>
        <div className={styles.postUser}>
          <Image
            className={styles.avatar}
            src={avatarSrc}
            alt="avatar"
            width={40}
            height={40}
            priority
            onClick={goToProfile}
          />
        </div>
        <div className={styles.postUserInfo} onClick={goToProfile}
        >
          <h3>{username}</h3>
          <span className={styles.postDate}>{formatDate(post.created_at)}</span>
        </div>
        <span className={styles.groupName}>{post.group_name}</span>
      </div>

      <div className={styles.postContent}>
        {isValidUrl(imageSrc) && (
          <Image
            className={styles.postImage}
            src={imageSrc}
            alt="post image"
            width={500}
            height={300}
          />
        )}
        <p className={styles.content}>{post.content}</p>
      </div>

      <div className={styles.postActions}>
        <div className={styles.postActionsLeft}>
          <button onClick={handleLike} disabled={isLiking}>
            <BiHeart className={liked ? styles.postIconLiked : styles.postIcon} />
            <span className={styles.postLikeCount}>{likes}</span>
          </button>
          <button onClick={() => router.push(`/post/${post.id}`)}>
            <BiComment className={styles.postIcon} />
            <span className={styles.postCommentCount}>{post.comments || 0}</span>
          </button>
        </div>
        <div className={styles.postActionsRight}>
          <button>
            <BiShareAlt className={styles.postIcon} />
          </button>
        </div>
      </div>
    </div>
  );
}
