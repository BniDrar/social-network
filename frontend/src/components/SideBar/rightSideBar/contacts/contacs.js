import styles from "./contacs.module.css";
import profile from "@/assets/images/profile.webp";
import Image from "next/image";
import GetLastChats from "@/services/lastChats";

 function Contacs() {
  const contacts =  GetLastChats();

  console.log("Contacts container:", contacts);

  return (
    <div className={styles.contacts}>
      {Array.from({ length: 10 }).map((user, i) => (
        <Contact user={user} key={i} />
      ))}
    </div>
  );
}
function Contact({user}) {
  console.log(user.name)
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
