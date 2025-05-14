"use client"
import styles from "./userFollow.module.css"
import { useEffect, useState } from "react"

const UserFollow = ({ is_following, id, status }) => {
    const [follow, setFollow] = useState("follow")
    
    const useFollow = async () => {
        try {
            const response = await fetch(`${process.env.BACKEND_URL}/api/user/follow/?followed=${id}`, {
                headers: {"Content-Type": "application/json"},
                credentials: "include"
            })
            if (!response.ok) throw new Error(`response error status ${response.status}`)
            if (status === 1 && follow === "follow") {
                setFollow("pending")
            } else if (status === 1 && follow === "pending") {
                setFollow("pending")
            } else if (status === 0 && follow === "follow") {
                setFollow("unfollow")
            } else if (status === 0 && follow === "unfollow") {
                setFollow("follow")
            }
            console.log(`response status ${response.status}`)
        } catch (error) {
            console.error(error)
        }
    }
    useEffect(()=>{
        if(is_following === 0) {
            setFollow("follow")
        } else if (is_following === 1) {
            setFollow("unfollow")
        } else if (is_following === 2){
            setFollow("pending")
        }
    },[])
    return (
        <>
            <button className={styles.followBtn} onClick={() => useFollow()}>{follow}</button>
        </>
    );
}

export default UserFollow;