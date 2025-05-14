'use client';

import React, { useState, useRef, useEffect } from "react";
import styles from "./createEvent.module.css";
import Image from "next/image";
import { createEvent } from "@/services/events";
import { useUser } from "@/context/userContext";

import {
      BiImageAdd,
      BiCalendarEvent,
      BiLocationPlus,
      BiTime,
      BiX,
} from "react-icons/bi";
import { useRouter } from "next/navigation";


export default function CreateEvent({ id }) {
      const popupRef = useRef(null);
      const [successMessage, setSuccessMessage] = useState("");
      const [errorMessage, setErrorMessage] = useState("");
      const router = useRouter();
      const { user } = useUser();

      const togglePopup = (show) => {
            if (!popupRef.current) return;
            popupRef.current.classList.toggle(styles.hidden, !show);
            popupRef.current.classList.toggle(styles.active, show);
      };

      const handleSubmit = async (e) => {
            e.preventDefault();
            const form = e.target;
            const formData = new FormData(form);
            const date = formData.get("date");
            const time = formData.get("time");

            let fullDate = "";
            if (date && time) {
                  fullDate = `${date}T${time}`;
            } else if (date) {
                  fullDate = `${date}T00:00:00`;
            } else {
                  throw new Error("Date is required");
            }
            const eventData = {
                  title: formData.get("title"),
                  location: formData.get("location"),
                  description: formData.get("description"),
                  date: fullDate,
                  group_id: parseInt(id),
            };


            try {
                  const response = await createEvent(eventData);
                  if (response.status === 201) {
                        form.reset();
                        setSuccessMessage("Event created successfully!");
                        setTimeout(() => setSuccessMessage(""), 2000);
                        setTimeout(() => togglePopup(false), 1000);

                        // Redirect to the event page if available
                        if (response.data && response.data.event_id) {
                              router.push(`/event/${response.data.event_id}`);
                        } else {
                              // Redirect to group page or events page
                              router.push(groupId ? `/group/${groupId}` : '/events');
                        }
                  } else {
                        setErrorMessage(response.error || "Something went wrong.");
                        setTimeout(() => setErrorMessage(""), 5000);
                  }
            } catch (err) {
                  setErrorMessage("Failed to create event.");
                  setTimeout(() => setErrorMessage(""), 5000);
            }
      };

      return (
            <>
                  <div className={styles.createEventContainer}>
                        <div className={styles.createEventHeader}>
                              {user.avatar && (
                                    <Image
                                          className={styles.avatar}
                                          src={user.avatar || "/default-avatar.jpeg"}
                                          alt="Avatar"
                                          width={40}
                                          height={40}
                                    />
                              )}
                              <button
                                    type="button"
                                    className={styles.createEventForm}
                                    onClick={() => togglePopup(true)}
                              >
                                    <input
                                          type="text"
                                          className={styles.createEventInput}
                                          placeholder="Create an event..."
                                          readOnly
                                    />
                              </button>
                        </div>
                        <div className={styles.createEventFooter}>
                              <div className={styles.createEventActions}>
                                    <button
                                          className={styles.createEventActionButton}
                                          type="button"
                                          onClick={() => togglePopup(true)}
                                    >
                                          <BiCalendarEvent className={styles.createEventIcon} />
                                    </button>
                                    <button className={styles.createEventActionButton}>
                                          <BiLocationPlus className={styles.createEventIcon} />
                                    </button>
                              </div>
                        </div>
                  </div>

                  {/* Popup */}
                  <div ref={popupRef} className={`${styles.createEventPopup} ${styles.hidden}`}>
                        <form onSubmit={handleSubmit}>
                              <div className={styles.createEventPopupHeader}>
                                    <div className={styles.createEventPopupUser}>
                                          {user.avatar && (
                                                <Image
                                                      className={styles.avatar}
                                                      src={user.avatar}
                                                      alt="Avatar"
                                                      width={40}
                                                      height={40}
                                                />
                                          )}
                                    </div>
                                    <div className={styles.createEventPopupUserInfo}>
                                          <h3>{user.nickname}</h3>
                                    </div>
                              </div>
                              <div className={styles.createEventPopupContent}>
                                    <div className={styles.formGroup}>
                                          <label htmlFor="title">Event Title</label>
                                          <input
                                                type="text"
                                                id="title"
                                                name="title"
                                                required
                                                placeholder="Enter event title"
                                                className={styles.formControl}
                                          />
                                    </div>
                                    <div className={styles.formGroup}>
                                          <label htmlFor="location">Location</label>
                                          <input
                                                type="text"
                                                id="location"
                                                name="location"
                                                required
                                                placeholder="Enter event location"
                                                className={styles.formControl}
                                          />
                                    </div>
                                    <div className={styles.formGroup}>
                                          <label htmlFor="description">Description</label>
                                          <textarea
                                                id="description"
                                                name="description"
                                                rows="4"
                                                required
                                                placeholder="Describe your event"
                                                className={styles.formControl}
                                          ></textarea>
                                    </div>

                                    <div className={styles.formRow}>
                                          <div className={styles.formGroup}>
                                                <label htmlFor="date">Date</label>
                                                <input type="date" id="date" name="date" required className={styles.formControl} />
                                          </div>
                                          <div className={styles.formGroup}>
                                                <label htmlFor="time">Time</label>
                                                <input type="time" id="time" name="time" required className={styles.formControl} />
                                          </div>
                                    </div>
                                    {errorMessage && (
                                          <div className={styles.errorMessage}>{errorMessage}</div>
                                    )}
                                    {successMessage && (
                                          <div className={styles.successMessage}>{successMessage}</div>
                                    )}

                                    <div className={styles.createEventPopupSubmitButton}>
                                          <button type="submit" className={styles.createEventPopupSubmitButton}>
                                                Create Event
                                          </button>
                                          <button
                                                type="button"
                                                className={styles.createEventPopupCancelButton}
                                                onClick={() => togglePopup(false)}
                                          >
                                                <BiX />
                                          </button>
                                    </div>
                              </div>
                        </form>
                  </div>
            </>
      );
}
