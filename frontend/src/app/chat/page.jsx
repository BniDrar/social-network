"use client";

import styles from "./chat.module.css";
// import Profile1 from "@/components/chatComponents/profile/profile";
import Contacts from "@/components/chatComponents/contacts/contacs";
import Chat from "@/components/chatComponents/chat/chat";
import { useState } from "react";
import Navbar from "@/components/naveBar/nave";
import Menu from "@/components/menu/menu";
import Profile from "@/components/profile/profile";
import CreateGroup from "@/components/createGroup/createGroup";
import Groups from "@/components/groups/groups";

export default function ChatPage() {
  const [id, setId] = useState(null);
  const [isGroup, setIsGroup] = useState(false);
  
  return (
    <div className={styles.page}>
      <Navbar />
      <main className={styles.main}>
        <div className={styles.leftSidebar}>
          <Profile />
          <Menu />
        </div>
        <div className={styles.container}>
          <Side setId={setId} setIsGroup={setIsGroup} />
          <Main id={id} isGroup={isGroup} />
        </div>
        <div className={styles.rightSidebar}>
          <Groups />
          <CreateGroup />
        </div>
      </main>
    </div>
  );
}

function Side({ setId, setIsGroup }) {
  return (
    <div className={styles.side}>
      <Contacts setId={setId} setIsGroup={setIsGroup} />
    </div>
  );
}

function Main({ id , isGroup }) {
  return (
    <div className={styles.chatForm}>
      {!id && <div className={styles.search}>Search</div>}
      {id && <Chat id={id} isGroup={isGroup} />}
    </div>
  );
}
