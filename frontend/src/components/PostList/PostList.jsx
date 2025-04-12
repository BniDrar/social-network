import Post from "../Post/Post";
import { backendUrl, validbackendUrl } from "@/utils/ustil";

const PostList = async () => {
  const res = await fetch(`${validbackendUrl}/api/posts`, {
    cache: "no-store",
  });
  console.log("res", res);
  const posts = await res.json();
  return (
    <section>
      {posts.map((post) => {
        return <Post key={post.id} data={post} />;
      })}
    </section>
  );
};

export default PostList;
