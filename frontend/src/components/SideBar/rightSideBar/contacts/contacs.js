import styles from "./contacs.module.css";
import profile from "@/assets/images/profile.webp";
import Image from 'next/image';


function Contacs() {
  return (
    <div className={styles.contacts}>
      {Array.from({ length: 10 }).map((_, i) => (
        <Contact key={i} />
      ))}
    </div>
  );
}
function Contact() {
  return (
    <div className={styles.contact}>
      <div className={styles.picContainer}>
        <Image
          width={600}
          height={600}
          alt="post"
          src={profile}
          className={styles.img}
        />
        :
        <img src="profile.webp" alt="profile" className={styles.img} />
      </div>
      <div className={styles.info}>
        <p>user name</p>
        <button className={styles.online} />
      </div>
    </div>
  );
}
async function GetContacs(groupId) {
  try {
    const response = await fetch(
      `${validbackendUrl}/api/contacts?id=${groupId}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching group:", error);
    throw error;
  }
}

export default Contacs;
