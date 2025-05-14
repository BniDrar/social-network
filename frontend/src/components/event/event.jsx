import React, { useState } from "react";
import styles from "./event.module.css";
import Image from "next/image";
import { voteEvent } from "@/services/events";
import { useRouter } from "next/navigation";

function Event({ event }) {
  const [isGoing, setIsGoing] = useState(event.going === 1);
  const router = useRouter();

  const eventDate = new Date(event.date);
  const isValid = !isNaN(eventDate);

  const dayName = isValid ? eventDate.toLocaleDateString("en-US", { weekday: "short" }).toUpperCase() : "N/A";
  const month = isValid ? eventDate.toLocaleDateString("en-US", { month: "short" }).toUpperCase() : "N/A";
  const day = isValid ? eventDate.getDate() : "--";
  const hour = isValid ? eventDate.toLocaleTimeString("en-US", { hour: "2-digit", minute: "2-digit" }) : "--:--";
  const [eventHour, ampm] = isValid ? hour.split(" ") : ["--", ""];

  const now = new Date();
  let status = "Upcoming";
  if (now > eventDate) {
    status = "Completed";
  } else if (
    now.toDateString() === eventDate.toDateString() &&
    now.getHours() === eventDate.getHours()
  ) {
    status = "Ongoing";
  }

  const handleVote = async () => {
    try {
      await voteEvent({ event_id: event.id, status: 1 });
      setIsGoing(true); // trigger re-render
    } catch (error) {
      console.error("Error voting for event:", error);
    }
  };

  return (
    <div className={styles.event} onClick={() => router.push(`/event/${event.id}`)}>
      <div className={styles.eventHeader}>
        <Image
          src="/default-event.jpeg"
          width={400}  
          height={200}
          alt="Event image"
        />
        <div className={styles.eventDates}>
          <div className={styles.eventTime}>
            <span className={styles.eventDayName}>{dayName}</span>
            <span className={styles.eventHour}>{eventHour}</span>
            <span className={styles.eventAmPm}>{ampm}</span>
          </div>
          <div className={styles.eventDate}>
            <span className={styles.eventMonth}>{month}</span>
            <span className={styles.eventDay}>{day}</span>
          </div>
        </div>
        {!isGoing ? (
          <button className={styles.eventGoing} onClick={handleVote}>Going</button>
        ) : (
          <button className={styles.eventGoing}>You are going</button>
        )}
      </div>
      <div className={styles.eventContent}>
        <div className={styles.eventTitle}>
          <h3>{event.title}</h3>
          <span className={styles.eventStatus}>{status}</span>
          <div className={styles.eventLocation}>{event.location}</div>
        </div>
        <div className={styles.eventDescription}>
          <p>{event.description}</p>
        </div>
      </div>
    </div>
  );
}

export default Event;