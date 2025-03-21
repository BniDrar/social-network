import styles from '@/components/post/post.module.css'
import avatar from "@/assets/images/no-face.jpg"
import post from "@/assets/images/post1.jpg"
import likes from '@/assets/icons/likes.svg'
import Image from 'next/image';

const Post = () => {
    return (
        <article className={styles.card}>
            <div className={styles.card_header}>
                {/* avatar - username - creation_date */}
                <Image width={100} height={100} alt='avatar' src={avatar} className={styles.post_avatar} />
                <div className="">
                    <h4 className={styles.post_username}>yrahhaou</h4>
                    <span className={styles.post_created_at}>7h</span>
                </div>
            </div>
            <div className={styles.card_body}>
                {/* text - image - reactions_len - comments_len */}
                <p className={styles.post_text}>🍲🍖 Classic Homestyle Meatloaf with a Tangy Glaze 🍋✨</p>
                {/* <img src="" className={styles.post_image} alt="post image" /> */}
                <br />
                <Image width={500} height={500} alt='post' src={post} className={styles.post_image} />
                <div className="">
                    <div className={styles.post_reactions_len}>
                        <span className={styles.postLikesNum}>150</span>
                        <Image src={likes} width={30} height={30} alt='likes' className={styles.postLikesSvg} />
                    </div>
                    <div className={styles.post_comments_len}></div>
                </div>
            </div>
            <div className={styles.card_footer}>
                {/* add_reaction - add_comment */}

            </div>
        </article>
    );
}

export default Post;