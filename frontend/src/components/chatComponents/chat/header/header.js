import styles from "./header.module.css";
import Image from "next/image";

function Header({ id, name, image, status }) {
  return (
    <div className={styles.header}>
      <div className={styles.headerContainer}>
        <div className={styles.profilePic}>
          <Image
            src={`${process.env.BACKEND_URL}/api/pictures/${image}`}
            alt={'avatar'}
            width={50}
            height={50}
            className={styles.profileImage}
          />
          {/* Status indicator */}
        </div>
        <div className={styles.info}>
          <h4>{name}</h4>
          {/* <span className={styles.nickname}>@nickname</span> */}
          <div
            className={status ? styles.online : styles.offline}
          ></div>
        </div>
      </div>
    </div>
  );
}

export default Header;
