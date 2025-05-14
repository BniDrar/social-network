"use client";

import styles from "./contacs.module.css";
import Image from "next/image";
import { GetContacts } from "@/services/contacts";
import { useState, useEffect, useRef } from "react";

function Contacts({ setId, setIsGroup }) {
  const [contacts, setContacts] = useState([]);

  useEffect(() => {
    const fetchContacts = async () => {
      const data = await GetContacts();
      setContacts(data);
    };

    fetchContacts();
  }, []);

  return (
    <div className={styles.contacts}>
      <h3>Following</h3>
      <div className={styles.contactsList}>
        {contacts?.map((contact, i) => (
          <Contact
            key={i}
            id={contact.id}
            first={contact.first_name}
            last={contact.last_name}
            image={contact.avatar}
            status={contact.online}
            setId={setId}
            setIsGroup={setIsGroup}
            groupName={contact.group_name}
          />
        ))}
      </div>
    </div>
  );
}

function Contact({ id, first, last, image, status, setId, setIsGroup, groupName }) {
  const contactRef = useRef(null);

  const handleContactClick = (e) => {
    setId(id);
    if (first === "") {
      setIsGroup(true)
    }
  };

  let name = ""
  if (first === "") {
    name = groupName
  } else {
    name = `${first} ${last}`
  }
  let source = ""
  if (image !== "") {
    source = `${process.env.BACKEND_URL}/api/pictures/${image}`
  } else {
    if (!setIsGroup) {
      source = 'default-avatar.jpeg'
    }else{
      source = 'default-group.jpeg'
    }
  }
  return (
    <>
      {/* Contact Item */}
      <div
        className={styles.contact}
        onClick={handleContactClick}
        ref={contactRef}
      >
        <div className={styles.profilePic}>
          <Image
            src={source}
            width={50}
            height={50}
            alt={'avatar'}
            className={styles.profileImage}
          />
          {/* Status indicator */}
          <div
            className={status ? styles.online : styles.offline}
          ></div>
        </div>
        <div className={styles.info}>
          <p>{name}</p>
        </div>
      </div >
    </>
  );
}

export default Contacts;
