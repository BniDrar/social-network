import styles from "./style.module.css";
import Contacts from "./contacts/contacs";
import Search from "./search/search";

export const RightSideBar = () => {
  return (
    <div className={styles.rightSideBar}>
      <Search />
      <Contacts />
      <div className={styles.events}>Events</div>
    </div>
  );
};
