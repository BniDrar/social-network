"use client";

import styles from "./profile.module.css";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { useUser } from "@/context/userContext";

export default function Profile() {
  const { user } = useUser();
  const router = useRouter();
  const profileUrl = `/profile/${user.myID}`;

  return (
    <div className={styles.profile}>
      <div className={styles.profileContainer}>
        <div className={styles.profileHeader} onClick={() => router.push(profileUrl)}>
          <Image
            className={styles.avatar}
            src={user.avatar || "/default-avatar.jpeg"}
            alt="avatar"
            width={50}
            height={50}
            priority
          />
          <div className={styles.profileInfo}>
            <h3>{`${user.first || ""} ${user.last || ""}`}</h3>
            <span className={styles.nickname}>
              {user.myNickname ? `@${user.myNickname}` : `${user.first || ""} ${user.last || ""}`}
            </span>
          </div>
        </div>
        <div className={styles.profileStats}>
          <div className={styles.stat}>
            <span className={styles.statNumber}>{user.followers_count || 0}</span>
            <span className={styles.statLabel}>Followers</span>
          </div>
          <div className={styles.stat}>
            <span className={styles.statNumber}>{user.following_count || 0}</span>
            <span className={styles.statLabel}>Following</span>
          </div>
        </div>
      </div>
    </div>
  );
}
