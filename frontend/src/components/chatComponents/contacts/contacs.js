"use client";

import styles from "./contacs.module.css";
import Image from "next/image";
import { GetContacts,GetFriends } from "@/services/contacts";
import { useState, useEffect, useRef } from "react";
import { useWebSocket } from "@/context/wsContext";
function Contacts({ setId, setIsGroup ,isFriends}) {
  const [contacts, setContacts] = useState([]);
  const [filteredContacts, setFilteredContacts] = useState([]);

  useEffect(() => {
    const fetchContacts = async () => {
      if (isFriends) {
        const data = await GetFriends();
        const allFriends = [...(data.Following || []), ...(data.Followers || [])];
        const uniqueFriends = Array.from(new Map(allFriends.map(item => [item.id, item])).values());
        setContacts(uniqueFriends);
        setFilteredContacts(uniqueFriends);
      } else {
        const data = await GetContacts();
        setContacts(data);
        setFilteredContacts(data);
      }
    };
    fetchContacts();
  }, []);

  const handleSearch = (e) => {
    e.preventDefault();
    const searchInput = e.target.elements.searchInput.value;
    if (searchInput.trim() === "") {
      setFilteredContacts(contacts);
      return;
    }
    const filtered = contacts.filter(contact => {
      const fullName = `${contact.first} ${contact.last}`.toLowerCase();
      const nickname = (contact.nickname || "").toLowerCase();
      const search = searchInput.toLowerCase();
      return fullName.includes(search) || nickname.includes(search);
    });
    setFilteredContacts(filtered);
    e.target.elements.searchInput.value = "";
  };

  return (
    <div className={styles.contacts}>
      <h2 className={styles.title}>{isFriends ? "Friends" : "Contacts"}</h2>
      {isFriends &&
        <form className={styles.searchForm} onSubmit={handleSearch}>
          <input
            type="text"
            id="searchInput"
            placeholder="Search friends..."
            className={styles.searchInput}
            onChange={e => {
              if (e.target.value.trim() === "") {
                setFilteredContacts(contacts);
              }
            }}
          />
        </form>
      }
      <div className={styles.contactsList}>
        {filteredContacts.length > 0 &&
        filteredContacts.map((contact, i) => (
          <Contact
            key={i}
            id={contact.id}
            first={contact.first_name || contact.first || ""}
            last={contact.last_name || contact.last || ""}
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
  const ws = useWebSocket();
  const [count, setCount] = useState(null)

  const contactRef = useRef(null);

  const handleContactClick = (e) => {
    setCount(null)
    setId(id);
    if (first === "") {
      setIsGroup(true)
    } else {
      setIsGroup(false)
    }
  };

  let name = ""
  if (first === "") {
    name = groupName
    status  = true
  } else {
    name = `${first} ${last}`
  }
  let source = ""
  if (image !== "") {
    source = `${process.env.BACKEND_URL}/api/pictures/${image}`
  } else {
    if (!setIsGroup) {
      source = 'default-avatar.jpeg'
    } else {
      source = 'default-group.jpeg'
    }
  }



  /*_____________________ notifications _____________________________*/
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
  /*_____________________ jsx______________________________________*/
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
        <span className={count && styles.notificationCount}>{count}</span>
      </div >
    </>
  );
}

export default Contacts;
