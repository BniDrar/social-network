import styles from './NewPostCard.module.css'

const NewPostCard = () => {
    return (
        <div className={styles.card}>
            <div className={styles.cardHeader}>
                {/* <Image className={styles.createPostAvatar} src={avatar} width={150} height={150} alt="" /> */}
                <form method="post">
                    <input type="text" name="" placeholder="What's on your mind..." className={styles.createPost} />
                </form>
            </div>
            <div className={styles.cardHeader}></div>
        </div>
    );
}

export default NewPostCard;