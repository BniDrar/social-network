import styles from "./header.module.css";
import Image from "next/image";

function Header({ id, name, image, status, isGroup }) {
  let source = ""
  if (image) {
    source = `${process.env.BACKEND_URL}/api/pictures/${image}`
  } else {
    if (!isGroup) {
      source = 'default-avatar.jpeg'
    } else {
      source = 'default-group.jpeg'
    }
  }
  console.log('------------>source------------>', source)
  return (
    <div className={styles.header}>
      <div className={styles.headerContainer}>
        <div className={styles.profilePic}>
          <Image
            src={source}
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
