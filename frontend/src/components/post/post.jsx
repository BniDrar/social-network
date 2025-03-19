import styles from '@/components/post/post.module.css'

const Post = () => {
    return (
        <article className={styles.card}>
            <div className={styles.card_header}>
                {/* avatar - username - creation_date */}
                <img src="" className={styles.post_avatar} alt="avatar" title='avatar' />
                <div className="">
                    <h4 className={styles.post_username}></h4>
                    <span className={styles.post_created_at}></span>
                </div>
            </div>
            <div className={styles.card_body}>
                {/* text - image - reactions_len - comments_len */}
                <p className={styles.post_text}></p>
                <img src="" className={styles.post_image} alt="post image" />
                <div className="">
                    <div className={styles.post_reactions_len}></div>
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