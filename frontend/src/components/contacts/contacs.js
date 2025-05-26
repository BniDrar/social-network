"use client";

import styles from "./contacs.module.css";
import Image from "next/image";
import ChatPopup from "./ChatPopup/ChatPopup.jsx"; // Import the ChatPopup component
import { GetContacts } from "@/services/contacts";
import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "@/context/wsContext";

function Contacts() {
  const [contacts, setContacts] = useState([]);

  useEffect(() => {
    const fetchContacts = async () => {
      const contacts = await GetContacts();
      setContacts(contacts)
    };

    fetchContacts();
  }, []);
  
  return (
    <div className={styles.contacts}>
      <h3>Contacts</h3>
      <div className={styles.contactsList}>
        {contacts?.map((contact, i) => (
          <Contact
            key={i}
            id={contact.id}
            first={contact.first_name}
            last={contact.last_name}
            image={contact.avatar}
            status={contact.online}
            groupName={contact.group_name}
          />
        ))}
      </div>
    </div>
  );
}

function Contact({ id, first, last, image, status, groupName }) {
  const isGroup = !!groupName
  if (isGroup) {
    status = true
  }
  const ws = useWebSocket();
  const [isChatOpen, setIsChatOpen] = useState(false);
  const [chatPosition, setChatPosition] = useState(null);
  const [count, setCount] = useState(null)

  const contactRef = useRef(null);



  const handleContactClick = (e) => {
    if (contactRef.current) {
      const rect = contactRef.current.getBoundingClientRect();
      const popupHeight = 285; // Example height (if needed)

      let top = rect.top;
      let right = window.innerWidth - rect.left + 10;

      // Optional: adjust top if popup goes off bottom
      if (top + popupHeight > window.innerHeight) {
        top = window.innerHeight - popupHeight;
        if (top < 0) top = 0; // Don't go above the screen
      }

      setChatPosition({ top, right });
    }
    setCount(null)
    setIsChatOpen(true);
  };

  const handleCloseChat = () => {
    setIsChatOpen(false);
  };
  let defaultPic = ""
  if (image) {
    defaultPic = `${process.env.BACKEND_URL}/api/pictures/${image}`
  } else if (first) {
    defaultPic = "/default-avatar.jpg"
  } else if (groupName) {
    defaultPic = "/default-group.jpg"
  }
  let name = ""
  if (first) {
    name = `${first} ${last}`
  } else if (groupName) {
    name = groupName
  }

  useEffect(() => {
    if (!ws) return;

    const handleMessage = (event) => {
      // PrivateMessage MessageType = iota
      // GroupMessage
      // BroadcastMessage
      // NotificationMessage 
      try {
        const notif = JSON.parse(event.data);
        const jsonString = atob(notif.Data);
        const message = JSON.parse(jsonString);

        const msg64 = notif.Data;

        // Safely decode base64 to UTF-8
        const decodedStr = decodeURIComponent(escape(atob(msg64)));

        const decoded = JSON.parse(decodedStr);
        const from = decoded.from

        console.log('decoded notif 404', decoded)

        switch (notif.Type) {
          case 0: // PrivateMessage
            if (from === id && !groupName) {
              setCount(count => (count ?? 0) + 1);
            }
            break;
          case 1: // GroupMessage
            if (from === id && groupName) {
              setCount(count => (count ?? 0) + 1);
            }
            break;
          case 2: // BroadcastMessage
            break;
          case 3: // NotificationMessage
            break;
          default:
            console.warn("Unknown message type:", message.type);
        }

      } catch (err) {
        console.error("Failed to parse message:", err);
      }
    };

    ws.addEventListener("message", handleMessage);

    return () => {
      ws.removeEventListener("message", handleMessage);
    };
  }, [ws]);

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
            src={defaultPic}
            width={50}
            height={50}
            alt={"avatar"}
            className={styles.profileImage}
          />
          {/* Status indicator */}
          <div
            className={status === true ? styles.online : styles.offline}
          ></div>
        </div>
        <div className={styles.info}>
          <p>{name}</p>
        </div>
        <span className={count && styles.notificationCount}>{count}</span>
      </div >

      {/* Chat Popup */}
      {
        isChatOpen && chatPosition && (
          <ChatPopup
            id={id}
            first={first}
            last={last}
            status={status}
            position={chatPosition} // Pass the dynamic position
            onClose={handleCloseChat} // Close function
            isGroup={isGroup}
          />
        )
      }
    </>
  );
}

export default Contacts;
