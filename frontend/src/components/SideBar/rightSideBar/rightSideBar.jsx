import styles from "./style.module.css";
import Contacts from "./contacts/contacs"

export const RightSideBar = () => {
  return (
    <div className={styles.rightSideBar}>
      <div className={styles.search}>Search</div>
      <Contacts />
      <div className={styles.events}>Events</div>
    </div>
  );
};

