import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axiosInstance from "../axiosInstance";

const HomePage = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axiosInstance.get("/api/posts");
        setPosts(response.data);
      } catch (error) {
        console.error("Failed to fetch posts", error);
      } finally {
        setLoading(false);
      }
    };
    fetchPosts();
  }, []);

  if (loading) {
    return <p>Loading posts...</p>;
  }

  return (
    <div className="min-h-screen bg-gray-100 dark:bg-gray-900 flex items-center justify-center">
      <div className="p-6 bg-white dark:bg-gray-800 shadow-md rounded-md w-3/4">
        <h1 className="text-3xl font-bold mb-6 dark:text-white">Posts</h1>
        {posts.length > 0 ? (
          <ul>
            {posts.map((post) => (
              <li key={post.id} className="mb-4">
                <Link
                  to={`/posts/${post.id}`}
                  className="text-xl text-blue-600 dark:text-blue-400"
                >
                  {post.title}
                </Link>
                <p className="dark:text-gray-300">{post.content.slice(0, 100)}...</p>
              </li>
            ))}
          </ul>
        ) : (
          <p className="dark:text-gray-300">No posts available</p>
        )}
      </div>
    </div>
  );
};

export default HomePage;
