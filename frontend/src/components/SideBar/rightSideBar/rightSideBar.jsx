import styles from "./style.module.css";

export const RightSideBar = () => {
  return (
    <div className={styles.rightSideBar}>
      <div className={styles.search}>Search</div>
      <div className={styles.contacts}>Contacts</div>
      <div className={styles.events}>Events</div>
    </div>
  );
};

