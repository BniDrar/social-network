"use client"
import { useState, useEffect } from "react"
import { useParams } from "next/navigation"
import NavBar from "@/components/naveBar/nave"
import styles from "./page.module.css"
import Image from "next/image"
import Link from "next/link"

const FollowersPage = () => {
    const { userid } = useParams()
    const [followers, setFollowers] = useState([])
    const getFollowers = async (userId) => {
        try {
            const response = await fetch(`${process.env.BACKEND_URL}/api/user/follower_and_followed?userid=${userId}`, {
                method: "GET",
                headers: { "Content-Type": "application/json" },
                credentials: "include"
            })
            if (!response.ok) {
                console.error("Failed to fetch followers")
                return;
            }
            const data = await response.json()
            setFollowers(await data.Followers)
        } catch (err) {
            console.error(err)
        }
    }
    useEffect(() => {
        getFollowers(userid)
    }, [])
    return (
        <>
            <NavBar />
            <main className={styles.container}>
                <h1 className={styles.title}>Following</h1>
                <div className={styles.grid}>
                    {followers.map(el => <div key={el.id} className={styles.card}>
                        <Link href={`/profile/${el.id}`}>
                            <Image src={`${process.env.MEDIA_URL}${el.avatar}`} width={150} height={150} alt="user avatar" />
                            <h4 className={styles.userName}>{el.first} {el.last}</h4>
                        </Link>
                    </div>
                    )}
                </div>

            </main>
        </>
    );
}

export default FollowersPage;