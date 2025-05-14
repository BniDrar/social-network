'use client'
import styles from './comment.module.css';
import{formatDate} from '@/utils/formateDate';
import Image from 'next/image';
export default function Comment({ comment }) {
      const avatarSrc = comment.avatar ? `${process.env.MEDIA_URL}${comment.avatar}` : "/default-avatar.jpeg";
      const imageSrc = comment.image ? `${process.env.MEDIA_URL}${comment.image}` : null;
      return (
            <div className={styles.comment}>
                  <div className={styles.commentHeader}>
                        <div className={styles.commentUser}>
                              <Image
                                    className={styles.avatar}
                                    src={avatarSrc}
                                    alt="avatar"
                                    width={40}
                                    height={40}
                                    priority
                              />
                              <div className={styles.commentUserInfo}>
                                    <h3>{comment.creater_name}</h3>
                                    <span className={styles.commentDate}>{formatDate(comment.created_at)}</span>
                              </div>
                        </div>
                        <div className={styles.commentContent}>
                              {
                                    imageSrc && (
                                          <Image
                                                className={styles.commentImage}
                                                src={imageSrc}
                                                alt="image"
                                                width={200}
                                                height={200}
                                                priority
                                          />
                                    )
                              }
                              <p>{comment.content}</p>
                        </div>
                  </div>
            </div>

      )
}