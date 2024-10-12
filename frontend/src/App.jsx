import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import HomePage from "./pages/HomePage";
import PostDetailPage from "./pages/PostDetailPage";
import CreatePostPage from "./pages/CreatePostPage";
import ThemeSwitcher from "./components/ThemeSwitcher";
import TopMenu from "./components/TopMenu";
import ContactsPage from "./pages/ContactsPage";

function App() {
  return (
    <div className="min-h-screen bg-white dark:bg-gray-900 text-black dark:text-white">
      <Router>
        <TopMenu />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/posts/:id" element={<PostDetailPage />} />
          <Route path="/create-post" element={<CreatePostPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/life-circle" element={<div>Life Cirlcle</div>} />
          <Route path="/profile" element={<div>User profile</div>} />
          <Route path="/contacts" element={<ContactsPage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
