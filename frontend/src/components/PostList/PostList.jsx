import Post from "../Post/Post";
import { backendUrl } from "@/utils/ustil";

const PostList = async () => {
    const res = await fetch(`${backendUrl}/posts`, {
        cache: 'no-store'
    })
    const posts = await res.json()
    return (
        <section>
            {posts.map(post => {
                return <Post key={post.id} data={post} />
            })}
        </section>
    );
}

export default PostList;