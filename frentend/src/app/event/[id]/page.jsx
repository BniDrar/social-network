'use client';

import { use } from 'react';
import styles from "./page.module.css";
import Navbar from "@/components/naveBar/nave";
import Event from "@/components/event/event";
import Menu from "@/components/menu/menu";
import Contacts from "@/components/contacts/contacs";
import Profile from "@/components/profile/profile";
import { getEventById } from "@/services/events";
import { useEffect, useState } from "react";
import CreateGroup from "@/components/createGroup/createGroup";
import Groups from "@/components/groups/groups";

export default function EventPage({ params }) {
  const { id: groupId } = use(params);
  const [event, setEvent] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchEvent() {
      try {
        const fetchedEvent = await getEventById(groupId);
        setEvent(fetchedEvent);
      } catch (error) {
        console.error("Failed to fetch event:", error);
      } finally {
        setLoading(false);
      }
    }
    fetchEvent();
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
          <div className={styles.events}>
            {loading && <p>Loading event...</p>}
            {!loading && event && <Event event={event.data} />}
            {!loading && !event && <p>Event not found.</p>}
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