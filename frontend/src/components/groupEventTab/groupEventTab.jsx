'use client';

import React, { useEffect, useState } from 'react';
import styles from './groupEventTab.module.css';
import { GetEventsByGroupId } from '@/services/events';
import Event from '@/components/event/event';
import CreateEvent from '../createEvent/createEvent';


export default function GroupEventsTab({ id }) {
  const [events, setEvents] = useState([]);

  const fetchEvents = async () => {
    try {
      const fetchedEvents = await GetEventsByGroupId(id);
      setEvents(fetchedEvents.data);
    } catch (error) {
      console.error("Error fetching group events:", error);
    }
  };

  useEffect(() => {
    fetchEvents();
  }, [id]);

  return (
    <div className={styles.groupEventsfetchedEventsTab}>
      <CreateEvent id={id} />
      <div className={styles.events}>
        {events?.length > 0 ? (
          events.map((event) => <Event key={event.id} event={event} />)
        ) : (
          <div className={styles.noEventsfetchedEvents}>No events available</div>
        )}
      </div>
    </div>
  );
}
