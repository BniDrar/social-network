'use client';

import { useEffect, useState, use } from "react";
import Image from "next/image";
import styles from "./page.module.css";
import Navbar from "@/components/naveBar/nave";
import Menu from "@/components/menu/menu";
import Contacts from "@/components/contacts/contacs";
import Profile from "@/components/profile/profile";
import CreateGroup from "@/components/createGroup/createGroup";
import { GetGroup } from "@/services/group";
import GroupPostsTab from "@/components/groupPostsTab/groupPostsTab";
import Groups from "@/components/groups/groups";
import GroupEventsTab from "@/components/groupEventTab/groupEventTab";
import GroupListMumbers from "@/components/groupListMumbers/groupListMumbers";
import SuggestedUsersList from "@/components/SuggestedUsersList/SuggestedUsersList";

export default function GroupPage({ params }) {
  const { id: groupId } = use(params);
  const [group, setGroup] = useState(null);
  const [activeTab, setActiveTab] = useState("posts");

  useEffect(() => {
    async function fetchGroup() {
      try {
        const fetchedGroup = await GetGroup(groupId);
        setGroup(fetchedGroup);
      } catch (error) {
        console.error("Failed to fetch group:", error);
      }
    }
    fetchGroup();
  }, [groupId]);

  return (
    <div className={styles.page}>
      <Navbar />
      <main className={styles.main}>
        <div className={styles.leftSidebar}>
          <Profile />
          <Menu />
        </div>
        <div className={styles.container}>
          <div className={styles.GroupHeader}>
            <Image
              src="/default-cover.jpg"
              className={styles.coverImg}
              width={1200}
              height={500}
              alt="cover"
            />
            {group ? (
              <div className={styles.groupInfo}>
                <h2>{group.name}</h2>
                <p>{group.description}</p>
                <span>Members {group.member_count}</span>
                <span>Posts {group.post_count}</span>
              </div>
              
            ) : (
              <p>Loading...</p>
            )}
           
            <SuggestedUsersList groupId={groupId} className= {styles.SuggestedUsersList} />
            
          </div>

          {/* Tab Buttons */}
          <div className={styles.groupContent}>
            <menu className={styles.menu}>
              <button
                className={`${styles.tabButton} ${activeTab === "posts" ? styles.active : ""}`}
                onClick={() => setActiveTab("posts")}
              >
                Posts
              </button>
              <button
                className={`${styles.tabButton} ${activeTab === "events" ? styles.active : ""}`}
                onClick={() => setActiveTab("events")}
              >
                Events
              </button>
              <button
                className={`${styles.tabButton} ${activeTab === "members" ? styles.active : ""}`}
                onClick={() => setActiveTab("members")}
              >
                Members
              </button>
            </menu>
            <div className={styles.tabContent}>
              <div style={{ display: activeTab === "posts" ? "block" : "none" }}>
                <GroupPostsTab id={groupId} />
              </div>
              <div style={{ display: activeTab === "events" ? "block" : "none" }}>
                <GroupEventsTab id={groupId} />
              </div>
              <div style={{ display: activeTab === "members" ? "block" : "none" }}>
                <GroupListMumbers id={groupId} />
              </div>
            </div>
          </div>
        </div>

        <div className={styles.rightSidebar}>
          <Contacts />
          <Groups />
          <CreateGroup />
        </div>
      </main>
    </div>
  );
}