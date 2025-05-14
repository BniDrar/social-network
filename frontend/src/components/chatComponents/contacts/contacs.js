"use client";

import styles from "./contacs.module.css";
import Image from "next/image";
import { GetContacts } from "@/services/contacts";
import { useState, useEffect, useRef } from "react";

function Contacts({ setId }) {
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
          />
        ))}
      </div>
    </div>
  );
}

function Contact({ id, first, last, image, status, setId }) {
  const contactRef = useRef(null);

  const handleContactClick = (e) => {
    setId(id);
  };

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
            src={`${process.env.BACKEND_URL}/api/pictures/${image}` || 'default-avatar.jpeg'}
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
          <p>{`${first} ${last}`}</p>
        </div>
      </div >
    </>
  );
}

export default Contacts;
