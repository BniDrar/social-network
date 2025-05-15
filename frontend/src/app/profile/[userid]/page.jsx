import NavBar from "@/components/naveBar/nave";
import styles from "./page.module.css"
import Image from "next/image";
import cover from "@/assets/cover.jpg"
import man from "@/assets/man.png"
import CreatePost from "@/components/createPost/createPost";
import UserInfo from "@/components/userInfo/userInfo";
import { getUserInfo } from "@/services/profile";
import { Suspense } from "react";
import Posts from "@/components/posts/posts";
import SowOnmoble from "@/components/showOnMobile/showOnmoble";


async function ProfilePage({ params }) {
    const { userid } = await params
    const info = await getUserInfo(userid)

    return (
        <>
            <NavBar />
            {await info != null ? <main>
                <SowOnmoble />
                <section className={styles.profileHeaderBox}>
                    <div className={styles.profileHeader}>
                        <div className={styles.coverBox}>
                            {/* cover */}
                            <Image src={cover} className={styles.coverImg} width={1200} height={400} alt="cover" />
                        </div>
                        <div className={styles.userProfileBox}>
                            {/* prfile img && username */}
                            <Image className={styles.userProfileImg} src={info.avatar != null ? `${process.env.MEDIA_URL}/${info.avatar}` : man} width={100} height={100} alt="profile" />
                            <h3 className={styles.username}>{info.first}</h3>
                            <div>
                                <Suspense fallback={<div>Loading...</div>}>
                                    <UserInfo className={styles.userInfo} info={info} />
                                </Suspense>
                            </div>
                        </div>
                    </div>
                </section>
                <section className={styles.container}>
                    {userid == 0 ? <CreatePost /> : ""}
                    <br />
                    <Posts body={{ id: info.id, last_id: 0 }} lien={`${process.env.BACKEND_URL}/api/user/posts`} />
                </section>
            </main> : <div className={styles.noPosts} style={{ textAlign: "center" }}>User not found!</div>}
        </>
    );
}

export default ProfilePage;