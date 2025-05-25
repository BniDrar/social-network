"use client"
import styles from "./page.module.css"
import { useState, useEffect } from "react"
import { useParams } from "next/navigation";
import Image from "next/image";

const FollowersPage = () => {
    const { userid } = useParams();
    const [followers, setFollowers] = useState([])
    const getFollowers = async (userid) => {
        try {
            const url = userid != 0 ? `?userid=${userid}` : ``
            const response = await fetch(`${process.env.BACKEND_URL}/api/user/follower_and_followed${url}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                },
                credentials: "include"
            })
            if (!response.ok) throw new Error(`response error ${response.status}`);
            const data = await response.json()
            console.log(await data.Followers);
            
            setFollowers(await data.Followers)
        } catch (err) {
            console.error(err)
        }
    }

    useEffect(() => {
        getFollowers(userid)
    }, [])


    return (
        <main>
            followers {followers.map(user => {
                return (
                    <div key={user.id}>
                        <Image src={`${process.env.MEDIA_URL}/${user.avatar}`} width={100} height={100} alt="profile image" />
                        <h4 key={user.id}>{user.first} {user.last}</h4>
                    </div>
                )
            })}
        </main>
    );
}

export default FollowersPage;