"use client"
import styles from "./userInfo.module.css"
import { FaWindowClose } from "react-icons/fa";
import { useState } from "react";
import Image from "next/image";
import man from "@/assets/man.png"
import UserFollow from "../userFollow/userFollow";



const UserInfo = ({ info }) => {
    const data =  info
    const [show, setShow] = useState(false)
    const [status, setStatus] = useState(data.status === 1 ? "private" : "public")
    const changeStatus = async () => {
        const response = await fetch(`${process.env.BACKEND_URL}/api/user/changestatus`, {
            method: "UPDATE",
            credentials: "include"
        })
        if (response.ok) {
            setStatus(status === "private" ? "public" : "private");
        }
    }

    return (
        <>
            <div className={styles.box}>
                <button className={styles.showBtn} onClick={() => setShow(true)}>Show Info</button>
                {info.profile_owner ? "" : <UserFollow is_following={data.following_state} id={data.id} status={data.status} />}
            </div>
            <div onClick={() => setShow(false)} className={styles.overlay} style={{
                display: show ? 'block' : 'none',
            }}></div>
            <article className={styles.card} style={{
                display: show ? 'block' : 'none',
            }}>
                <div className={styles.card_header}>
                    {/* avatar && first_name && last_name */}
                    <FaWindowClose className={styles.closeBtn} onClick={() => setShow(false)} />
                    <Image className={styles.avatar} src={data.avatar ? `${process.env.MEDIA_URL}/${data.avatar}` : man} width={100} height={100} alt="avatar image" />
                    <div className="">
                        <h5 className={styles.fullName}>{`${data.first} ${data.last}`}</h5>
                        {data.profile_owner ?
                            <select name="privacy" onChange={() => changeStatus()} id="privacy" defaultValue={status} className={styles.privacy}>
                                <option value="public">Public</option>
                                <option value="private">Private</option>
                            </select>
                            : ""}
                    </div>

                </div>
                <div className={styles.cardBody}>
                    {/* following, followers */}
                    <div className={styles.friends}>
                        <div className={styles.friends_link}>{data.following_count ? data.following_count : 0} Following</div>
                        <div className={styles.friends_link}>{data.followers_count ? data.followers_count : 0} Followers</div>
                    </div>
                    <br />
                    {/* nickname, date of birth, email, about me */}
                    {data.nickname ? <p className={styles.text}><b>Nickname:</b> {data.nickname}</p> : ''}
                    {data.email ? <p className={styles.text}><b>Email:</b> {data.email}</p> : ""}
                    {data.birthday ? <p className={styles.text}><b>Date of birth:</b> {data.birthday.split("T").slice(0, 1)}</p> : ""}
                    {data.about_me ? <p className={styles.text}><b>About me:</b> {data.about_me}</p> : ''}

                </div>
            </article>
        </>
    );
}

export default UserInfo;