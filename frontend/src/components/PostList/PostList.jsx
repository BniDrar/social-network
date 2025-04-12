import Post from "../Post/Post";
import { backendUrl, validbackendUrl } from "@/utils/ustil";

const PostList = async () => {
  //const res = await fetch(`${backendUrl}/posts`, { // fake database
  const res = await fetch(`${validbackendUrl}/api/posts?limit=3&offset=0`, {
    cache: "no-store",
    credentials: "include",
  });

  console.log("status:", res.status);
  console.log(res)
  if (!res.ok) {
    const text = await res.json();
    console.log("json:", text)
    console.error("Non-OK response:", res.status, text);
    throw new Error(`Request failed with status ${res.status}`);
  }

  const posts = await res.json();
  console.log("posts:", posts);

  return (
    <section>
      {posts.map((post) => {
        return <Post key={post.id} data={post} />;
      })}
    </section>
  );
};

export default PostList;
