import styles from "./profile.module.css";

function Profile() {
  return (
    <div className={styles.profile}>
      <div className={styles.pic}>
        <div className={styles.picContainer}>
          <img
            src="profile.webp" // Reference from the public folder
            alt="profile"
            className={styles.img}
          />
        </div>
      </div>

      <div className={styles.input}>
        <input
          type="text"
          placeholder="Search..."
          className={styles.searchInput}
        />
      </div>
      <div className={styles.info}>
        <div className={styles.left}>
          <button className={styles.inbox} role="button">Inbox</button>
        </div>
        <div className={styles.right}>
          <button className={styles.compose} role="button">+Compose</button>
        </div>
      </div>
    </div>
  );
}

export default Profile;
