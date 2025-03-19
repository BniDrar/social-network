import styles from '@/components/post/post.module.css'

const Post = () => {
    return (
        <article className={styles.card}>
            <div className={styles.card_header}>
                {/* avatar - username - creation_date */}
            </div>
            <div className={styles.card_body}>
                {/* text - image - reactions_len - comments_len */}
            </div>
            <div className={styles.card_footer}>
                {/* add_reaction - add_comment */}
            </div>
        </article>
    );
}

export default Post;