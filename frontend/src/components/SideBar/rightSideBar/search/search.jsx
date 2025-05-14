import styles from "./search.module.css";

function Profile() {
  return (
    <div className={styles.search}>
      <div className={styles.text}> Messages</div>
      <div className={styles.input}>
        <input
          type="text"
          placeholder="Search..."
          className={styles.searchInput}
        />
      </div>
    </div>
  );
}

export default Profile;
